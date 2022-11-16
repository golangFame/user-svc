package models

import (
	"context"
	log "github.com/sirupsen/logrus"
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
	IsMultiBidsEnabled bool `bun:"is_multibids_enabled"`

	StartTime time.Time `bun:"start_time"`
	EndTime   time.Time

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`

	//Auction *Auctions `rel:"belongs-to"`
}

type AuctionTypes struct {
	bun.BaseModel `bun:"table:auction_types"`

	ID int `bun:"id,pk,autoincrement"`

	Name string `bun:"name" json:"name"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`
}

type Auctions struct {
	bun.BaseModel `bun:"table:auctions"`

	ID int `bun:"id,pk,autoincrement"`

	Name               string `bun:"name" json:"name"`
	MinPoints          int
	AdditionalText     string   `json:"additionalText,omitempty"`
	AdditionalIcon     string   `json:"additionalIcon,omitempty"`
	Subtitle           string   `json:"subtitle,omitempty"`
	LayoutID           int      `json:"layoutId"`
	TitleIcon          string   `json:"titleIcon,omitempty"`
	Avatars            []string `json:"avatars,omitempty"`
	StartsAt           time.Time
	ExpiresAt          time.Time
	ShouldShowAvatars  bool
	Image              string
	LayoutNumber       int   `bun:"layout_id"` //FIXME check why this is working without the tag
	AppUserGroupID     []int `bun:"app_user_group_id,array"`
	Color              string
	ShowMrp            bool
	IsMultibidsEnabled bool

	AuctionTypeID int `bun:"auction_type_id" json:"auctionTypeID"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`

	Products *[]AuctionProducts `bun:"rel:has-many,join:id=auction_id" json:"products,omitempty""`

	HasUserAccess bool `json:"hasUserAccess"`
}

func (a *Auctions) SetUserAccess(userPoints int) {
	var flag bool
	if a.MinPoints > 0 && a.MinPoints > userPoints {
		flag = false
	} else {
		flag = true
	}
	a.HasUserAccess = flag
	return
}

func (a *AuctionTypes) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(a)

	if a.ID != 0 {
		query.WherePK()
	}
	if a.Name != "" {
		query.Where("name=?", a.Name)
	}

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

	if a.ID != 0 {
		query.WherePK()
	}

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
	if a.AuctionID != 0 {
		query.Where("auction_id=?", a.AuctionID)
	}

	query.
		Where("deleted_at is null").
		Where("start_time < ? AND end_time > ?", a.StartTime, a.EndTime)

	err = query.Scan(ctx)

	if err != nil {
		log.Error(err)
	}
	return
}

func (a *Auctions) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(a)

	if a.ID != 0 {
		query.WherePK()
	}

	if a.Name != "" {
		query.Where("name=?", a.Name)
	}

	err = query.Scan(ctx)
	return
}

func (a *Auctions) FetchAll(db *bun.DB, ctx context.Context) (auctions []Auctions, err error) {
	query := db.NewSelect().Model(&auctions)

	query.Where("deleted_at is null")

	if a.AuctionTypeID != 0 {
		query.Where("auction_type_id=?", a.AuctionTypeID)
	}

	err = query.Scan(ctx)
	return
}
