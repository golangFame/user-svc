package models

import (
	"context"
	"github.com/uptrace/bun"
)

type Dummy struct {
	bun.BaseModel `bun:"table:dummy"`
	ID            int `bun:"id"`
}

func (d *Dummy) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(d)
	query.Where("id=?", d.ID)
	err = query.Scan(ctx)
	return
}

func (d *Dummy) FetchAll(db *bun.DB, ctx context.Context) (providers []Dummy, err error) {
	query := db.NewSelect().Model(&providers)
	err = query.Scan(ctx)
	return
}
