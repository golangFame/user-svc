package app

import (
	"context"
	"encoding/json"
	"github.com/BzingaApp/user-svc/entities"
	"github.com/BzingaApp/user-svc/models"
	"github.com/BzingaApp/user-svc/utils"
)

type Services interface {
	AuctionProductsNow() (auctions []entities.Auctions)
}

func (s *service) getSpotTimer() (seconds int) {
	appPropertySpotTimer := models.AppProperties{
		ID:        20,
		Key:       "spotTimer",
		GroupName: "home",
	}
	appPropertySpotTimer.Fetch(s.db, context.TODO())

	seconds = utils.ConvertStringIntoInt(appPropertySpotTimer.Value)

	return
}

func (s *service) AuctionProductsNow() (auctions []entities.Auctions) {

	//implement active auctions types
	ctx := context.TODO()

	db := s.db

	userID := 16042003

	_userPoint := models.UserPoint{UserID: userID}

	_userPoint.Fetch(db, ctx)

	appPropertyActiveAuctions := models.AppProperties{
		ID:        39,
		Key:       "auctionSpotsHome",
		GroupName: "home",
	}

	appPropertyActiveAuctions.Fetch(db, ctx) // Val = {"1":"WELCOME BID","2":"EASY BID","3":"NOVEMBER SPECIAL","4":"MOBILE BAZAAR","5":"FESTIVAL EXPRESS","6":"WIN FAST","7":"SPOT BID","8":"GADGET GURU","9":"SOUND STUDIO","10":"MENS DEN","11":"VOUCHERS MELA","12":"SUPREME COLLECTION","13":"12 HOUR BID"}

	var orderAuctionTypes map[int]string

	utils.ConvertJSONToGoType(json.RawMessage(appPropertyActiveAuctions.Value), &orderAuctionTypes)

	//var auctionsCur [orderAuctionTypes]models.Auctions //map is not getting stuffs in order

	for _, val := range orderAuctionTypes {

		_auction := models.Auctions{
			Name: val,
		}
		_auction.Fetch(db, ctx)

		_auction.SetUserAccess(_userPoint.Points)

		if _auction.ID == 0 {
			s.Log.Error("invalid _auction -", val)
			continue
		}
		auctionProducts := models.AuctionProducts{
			AuctionID: _auction.ID,
		}

		auctionProds, err := auctionProducts.FetchAll(db, ctx)

		if err != nil {
			s.Log.Error("failed to retrieve _auction products for ", val)
			continue
		}

		var prodRes []entities.AuctionProducts

		for _, _auctionProduct := range auctionProds {
			_product := models.Product{
				ID: auctionProducts.ProductID,
			}
			db.NewSelect().Model(&_product).
				Relation("Images").
				Relation("Currency").
				Relation("Category").
				WherePK() //TODO is joins a good time??

			prodRes = append(prodRes, entities.AuctionProducts{
				_product.Name,
				_product.ID,
				_product.Currency.Symbol,
				_auctionProduct.MinBidPrice,
				_product.Msrp,
				0, // episode_product 's round number
				_product.ExtractLinks(),
				_auction.ExpiresAt, //FIXME Kritika: not sure
				true,               //FIXME
				string(_product.CategoryID),
				_auction.ID,
				0, //FIXME episodeID
				_product.Description,
				_product.CompanyImage,
				false, //FIXME get isLive
				_auctionProduct.GetUserAccess(_userPoint.Points),
				_auction.MinPoints, //FIXME
				[]string{},         //fixme
				"",                 //fixme
				true,               //fixme
				true,               //fixme
				true,               //fixme
				12,                 //total bids
				true,               //fixme
				0,                  //fixme auction_service ....
				"",                 //fixme
				_auction.AuctionTypeID,
				true, //fixme
			})
		}
		if len(prodRes) > 0 {
			auctions = append(auctions, entities.Auctions{
				_auction,
				prodRes,
			})

		}
	}
	return
}
