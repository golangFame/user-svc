package entities

import (
	"github.com/BzingaApp/user-svc/models"
	"time"
)

type AuctionProducts struct { //FIXME inherited from auction svc
	Name               string    `json:"name"`
	ID                 int       `json:"id"`
	Currency           string    `json:"currency"`
	MinBidPrice        float64   `json:"minBidPrice"`
	Msrp               float64   `json:"mrp"`
	RoundNumber        int       `json:"roundNumber"`
	ProductImages      []string  `json:"productImages,omitempty"`
	ExpiresAt          time.Time `json:"expiresAt,omitempty"`
	IsInWishList       bool      `json:"isInWishlist"`
	Category           string    `json:"category"`
	AuctionID          int       `json:"auctionId,omitempty"`
	EpisodeID          int       `json:"episodeId,omitempty"`
	Description        string    `json:"description"`
	CompanyImage       string    `json:"companyImage,omitempty"`
	IsLive             bool      `json:"isLiveNow,omitempty"`
	HasAccess          bool      `json:"hasUserAccess"`
	MinPoints          int       `json:"minPoints,omitempty"`
	RandomAvatars      []string  `json:"avatars,omitempty"`
	BidType            string    `json:"bidType,omitempty"`
	IsListed           bool      `json:"isListed,omitempty"`
	IsExpired          bool      `json:"isExpired,omitempty"`
	NoWinner           bool      `json:"noWinner,omitempty"`
	TotalBids          int       `json:"totalBids"`
	ShowMrp            bool      `json:"showMRP"`
	CoachCard          int64     `json:"coachCard"`
	CoachCardMsg       string    `json:"coachCardMsg"`
	AuctionTypeID      int       `json:"auctionTypeID"`
	IsMultibidsEnabled bool      `json:"isMultibidsEnabled"`
}

type Auctions struct {
	models.Auctions
	Products []AuctionProducts `json:"products"`
}
