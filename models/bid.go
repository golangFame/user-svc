package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Bid struct {
	bun.BaseModel `bun:"table:auctions"`

	ID int `bun:"id,pk,autoincrement"`

	UserID           int       `bun:"user_id"`
	Value            float64   `bun:"value"'`
	IsMultiBid       bool      `bun:"is_multi_bid"`
	BidForID         int       `bun:"bid_for_id"`
	BidAt            time.Time `bun:"bid_at"`
	ProductID        int       `bun:"product_id"`
	AuctionID        int       `bun:"auction_id"`
	AuctionProductID int       `bun:"auction_product_id"`
	EpisodeID        int       `bun:"episode_id"`
	EpisodeProductID int       `bun:"episode_product_id"`

	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
	DeletedAt time.Time `bun:"deleted_at,nullzero,soft_delete" json:"deletedAt,omitempty"`

	User *User `bun:"rel:belongs-to,join:app_user_id=id"`

	/*	Auction          *Auction
		Product          *Product
		WinningBid       *WinningBid

		AuctionProduct *AuctionProduct
		EpisodeProduct *EpisodeProduct
		WinningBot     *WinningBot*/
}
