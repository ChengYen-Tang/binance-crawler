package factory

import (
	"github.com/ChengYen-Tang/binance-crawler/models"
	database "github.com/ChengYen-Tang/binance-crawler/modules/database/interface"
)

type WorkItemFactory struct {
	db database.IDatabase
}

func NewWorkItemFactory(db database.IDatabase) *WorkItemFactory {
	return &WorkItemFactory{
		db: db,
	}
}

func (w *WorkItemFactory) CreateSpotKlineWorkItem(symbol *string, kline *models.Kline) *models.KlineWorkIten {
	apiName := models.SpotKline
	return models.NewKlineWorkIten(&apiName, symbol, kline, w.db.InsertKline)
}

func (w *WorkItemFactory) CreateUSDPremiumIndexKlineWorkItem(symbol *string, kline *models.Kline) *models.KlineWorkIten {
	apiName := models.USDFuturesPremiumIndexKline
	return models.NewKlineWorkIten(&apiName, symbol, kline, w.db.InsertKline)
}

func (w *WorkItemFactory) CreateUSDFuturesKlineWorkItem(symbol *string, kline *models.Kline) *models.KlineWorkIten {
	apiName := models.USDFuturesKline
	return models.NewKlineWorkIten(&apiName, symbol, kline, w.db.InsertKline)
}

func (w *WorkItemFactory) CreateUSDFuturesFundingRateWorkItem(symbol *string, fundingRate *models.FundingRate) *models.FundingRateWorkIten {
	apiName := models.USDFuturesFundingRate
	return models.NewFundingRateWorkIten(&apiName, symbol, fundingRate, w.db.InsertFundingRate)
}
