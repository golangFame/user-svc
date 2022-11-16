package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type AppProperties struct {
	bun.BaseModel `bun:"table:app_properties"`

	ID int `bun:"id,pk,autoincrement"`

	Key         string `bun:"key"`
	GroupName   string
	Value       string
	Description string

	FieldType string

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`
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
