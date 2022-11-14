package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type AppProperties struct {
	bun.BaseModel `bun:"table:auction_products"`

	ID int `bun:"id,pk,autoincrement"`

	Key         string `bun:"key"`
	GroupName   string
	Value       string
	Description string

	FieldType string

	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`

	DeletedAt time.Time `bun:"deleted_at"`
}

func (a *AppProperties) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(a)

	query.WherePK()

	err = query.Scan(ctx)
	return
}

func (a *AppProperties) FetchAll(db *bun.DB, ctx context.Context) (appProperties []AppProperties, err error) {
	query := db.NewSelect().Model(&appProperties)

	query.Where("deleted_at is null")

	err = query.Scan(ctx)
	return
}
