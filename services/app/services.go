package app

import (
	"context"
	"fmt"
	"github.com/BzingaApp/user-svc/models"
)

type Services interface {
	HomePage()
}

func (s *service) HomePage() {

	//implement active auctions types

	var auctionType models.AuctionTypes
	ctx := context.Background()

	auctionTypes, _ := auctionType.FetchAll(s.DB, ctx)

	fmt.Println(auctionTypes)
}
