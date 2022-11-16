package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Product struct {
	bun.BaseModel `bun:"table:products"`

	ID           int    `bun:"id,pk,autoincrement"`
	Name         string `bun:"name"`
	CurrencyID   int    `bun:"currency_id"`
	Code         string `bun:"code" json:"code,omitempty"`
	CompanyImage string `bun:"company_image"`
	Msrp         float64
	FloorPrice   float64
	StockUnits   int
	Description  string

	BidThreshold struct {
		MaxBidCount int `json:"max_bid_count"`
		MinBidCount int `json:"min_bid_count"`
	} `bun:"bid_threshold"` //FIXME Kritika : is this generalization correct?

	IsMultiBidsEnabled bool `bun:"is_multibids_enabled"`

	CategoryID int
	CoachCard  int
	ImageID    string `bun:"image_id"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`

	Currency *Currency `bun:"rel:belongs-to,join:currency_id=id"`
	Category *Category

	Images *[]Image `bun:"rel:has-many,join:image_id=id"`
	Bids   *[]Bid
}

type Currency struct {
	bun.BaseModel `bun:"table:currencies"`

	ID     int    `bun:"id,pk,autoincrement"`
	Name   string `bun:"name"`
	Symbol string `bun:"symbol"`
}

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	ID   int    `bun:"id,pk,autoincrement"`
	Name string `json:"name"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`

	Products []Product
}

type Image struct {
	bun.BaseModel `bun:"table:images"`

	ID        int    `bun:"id,pk,autoincrement"`
	Name      string `bun:"name"`
	ImageData struct {
		Id       string `json:"id"`
		Storage  string `json:"storage"`
		Metadata struct {
			Size     int         `json:"size"`
			Filename string      `json:"filename"`
			MimeType interface{} `json:"mime_type"` //FIXME @Kritika: essentially mimetypes are strings!!!!?
		} `json:"metadata"`
	} `bun:"image_data"`
	Order int `bun:"order"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`
}

func (p *Product) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(p)

	if p.ID != 0 {
		query.WherePK()
	}

	err = query.Scan(ctx)
	return
}

func (p *Product) FetchAll(db *bun.DB, ctx context.Context) (products []Product, err error) {
	query := db.NewSelect().Model(&products)

	query.Where("deleted_at is null")

	err = query.Scan(ctx)
	return
}

func (p *Product) ExtractLinks() (links []string) {

	for _, i := range *p.Images {
		links = append(links, i.ImageData.Id)
	}

	return
}
