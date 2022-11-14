package app

import (
	"context"
	"encoding/json"
	"github.com/BzingaApp/user-svc/models"
	"github.com/BzingaApp/user-svc/utils"
)

type Services interface {
	getSpotTimer() (seconds int)
	AuctionProductsNow() (auctions []models.Auctions)
}

func (s *service) getSpotTimer() (seconds int) {
	appPropertySpotTimer := models.AppProperties{
		ID:        20,
		Key:       "spotTimer",
		GroupName: "home",
	}
	appPropertySpotTimer.Fetch(s.DB, context.TODO())

	seconds = utils.ConvertStringIntoInt(appPropertySpotTimer.Value)

	return
}

func (s *service) AuctionProductsNow() (auctions []models.Auctions) {

	//implement active auctions types

	ctx := context.TODO()

	db := s.DB

	appPropertyActiveAuctions := models.AppProperties{
		ID:        39,
		Key:       "auctionSpotsHome",
		GroupName: "home",
	}

	appPropertyActiveAuctions.Fetch(db, ctx) // Val = {"1":"WELCOME BID","2":"EASY BID","3":"NOVEMBER SPECIAL","4":"MOBILE BAZAAR","5":"FESTIVAL EXPRESS","6":"WIN FAST","7":"SPOT BID","8":"GADGET GURU","9":"SOUND STUDIO","10":"MENS DEN","11":"VOUCHERS MELA","12":"SUPREME COLLECTION","13":"12 HOUR BID"}

	var orderAuctionTypes map[int]string

	utils.ConvertJSONToGoType(json.RawMessage(appPropertyActiveAuctions.Value), &orderAuctionTypes)

	for _, val := range orderAuctionTypes {
		//get the auction type id
		/*
			auctionType := models.AuctionTypes{
				Name:val,
			}

			auctionType.Fetch(db, ctx)
			if auctionType.ID==0{
				s.Log.Error("invalid auction type name ",val)
				continue
			}

		*/
		auction := models.Auctions{
			Name: val,
		}
		auction.Fetch(db, ctx)

		if auction.ID == 0 {
			s.Log.Error("invalid auction -", val)
			continue
		}
		auctionProducts := models.AuctionProducts{
			AuctionID: auction.ID,
		}

		var err error

		auction.Products, err = auctionProducts.FetchAll(db, ctx)

		if err != nil {
			s.Log.Error("failed to retrieve auction products for ", val)
		}

		auctions = append(auctions, auction)

	}

	return
}
