package models

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type AuctionProducts struct {
	bun.BaseModel `bun:"table:auction_products"`

	ID int `bun:"id,pk,autoincrement"`

	ProductID        int     `bun:"product_id"`
	AuctionID        int     `bun:"auction_id"`
	MinBidPrice      float32 `bun:"min_bid_price"`
	MinBidTickets    int     `bun:"min_bid_tickets"`
	MinStatusTickets int     `bun:"min_status_tickets"`
	MinPoints        int     `bun:"min_points"`
	StockUnits       int     `bun:"stock_units"`

	ConsolationStock   int
	IsMultiBidsEnabled bool

	StartTime time.Time `bun:"start_time"`
	EndTime   time.Time

	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`

	DeletedAt time.Time `bun:"deleted_at"`
}

type AuctionTypes struct {
	bun.BaseModel `bun:"table:auction_types"`

	ID int `bun:"id,pk,autoincrement"`

	Name string `bun:"name"`

	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`

	DeletedAt time.Time `bun:"deleted_at"`
}

type Auctions struct {
	bun.BaseModel `bun:"table:auctions"`

	ID int `bun:"id,pk,autoincrement"`

	Name               string `bun:"name"`
	MinPoints          uint
	Subtitle           string
	AdditionalText     sql.NullString
	AdditionalIcon     sql.NullString
	TitleIcon          sql.NullString
	StartsAt           time.Time
	ExpiresAt          sql.NullTime
	ShouldShowAvatars  bool
	Image              string
	LayoutNumber       uint
	AppUserGroupID     int
	Color              string
	ShowMrp            bool
	IsMultibidsEnabled bool

	AuctionTypeID int `bun:"auction_type_id"`

	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`

	DeletedAt time.Time `bun:"deleted_at"`
}

func (a *AuctionTypes) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(a)

	query.WherePK()

	err = query.Scan(ctx)
	return
}

func (a *AuctionTypes) FetchAll(db *bun.DB, ctx context.Context) (auctionTypes []AuctionTypes, err error) {
	query := db.NewSelect().Model(&auctionTypes)

	query.Where("deleted_at is null")

	err = query.Scan(ctx)
	return
}

func (a *AuctionProducts) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(a)

	query.WherePK()

	err = query.Scan(ctx)
	return
}

func (a *AuctionProducts) FetchAll(db *bun.DB, ctx context.Context) (auctionProducts []AuctionProducts, err error) {
	query := db.NewSelect().Model(&auctionProducts)

	if a.StartTime.IsZero() {
		a.StartTime = time.Now()
	}
	if a.EndTime.IsZero() {
		a.EndTime = time.Now()
	}

	query.
		Where("deleted_at is null").
		Where("start_time < ? AND end_time > ?", a.StartTime, a.EndTime)

	err = query.Scan(ctx)
	return
}

func (a *Auctions) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(a)

	query.WherePK()

	err = query.Scan(ctx)
	return
}

func (a *Auctions) FetchAll(db *bun.DB, ctx context.Context) (auctions []Auctions, err error) {
	query := db.NewSelect().Model(&auctions)

	query.Where("deleted_at is null")

	err = query.Scan(ctx)
	return
}
