package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:app_users"`

	ID int `bun:"id,pk,autoincrement"`

	FirstName           string
	LastName            string
	Avatar              string
	EncryptedPin        string
	Email               string
	Phone               int64
	CountryCode         int
	AddressLine         string
	State               string
	Country             string
	City                string
	Pincode             int32
	PinAttempts         int
	OTPSubmitAttempts   int
	OTPGenerateAttempts int
	UserGroupID         []int `bun:"user_group_id,array"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`

	UserPoints *UserPoint `bun:"rel:has-one,join:id=app_user_id"`
}

type UserPoint struct {
	bun.BaseModel `bun:"table:user_points"`

	ID int `bun:"id,pk,autoincrement"`

	UserID int `bun:"app_user_id"`
	Points int

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`

	User *User `bun:"rel:belongs-to,join:app_user_id=id"`
}

func (u *User) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(u)

	if u.ID != 0 {
		query.WherePK()
	}

	err = query.Scan(ctx)
	return
}

func (u *User) FetchAll(db *bun.DB, ctx context.Context) (User []User, err error) {
	query := db.NewSelect().Model(&User)

	query.Where("deleted_at is null")

	err = query.Scan(ctx)
	return
}

func (u *UserPoint) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(u)

	if u.ID != 0 {
		query.WherePK()
	}
	if u.UserID != 0 {
		query.Where("app_user_id=?", u.UserID)
	}

	err = query.Scan(ctx)
	return
}
