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
	UserGroupID         []int `bun:"user_group_id"`

	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`

	DeletedAt time.Time `bun:"deleted_at"`
}

func (u *User) Fetch(db *bun.DB, ctx context.Context) (err error) {
	query := db.NewSelect().Model(u)

	query.WherePK()

	err = query.Scan(ctx)
	return
}

func (u *User) FetchAll(db *bun.DB, ctx context.Context) (User []User, err error) {
	query := db.NewSelect().Model(&User)

	query.Where("deleted_at is null")

	err = query.Scan(ctx)
	return
}
