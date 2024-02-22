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

func (w *WorkItemFactory) CreateSpotKlinesWorkItem(symbol *string, klines *[]models.Kline) *models.KlinesWorkIten {
	apiName := models.SpotKline
	return models.NewKlinesWorkIten(&apiName, symbol, klines, w.db.InsertKlines)
}

func (w *WorkItemFactory) CreateUSDPremiumIndexKlineWorkItem(symbol *string, kline *models.Kline) *models.KlineWorkIten {
	apiName := models.USDFuturesPremiumIndexKline
	return models.NewKlineWorkIten(&apiName, symbol, kline, w.db.InsertKline)
}

func (w *WorkItemFactory) CreateUSDPremiumIndexKlinesWorkItem(symbol *string, klines *[]models.Kline) *models.KlinesWorkIten {
	apiName := models.USDFuturesPremiumIndexKline
	return models.NewKlinesWorkIten(&apiName, symbol, klines, w.db.InsertKlines)
}

func (w *WorkItemFactory) CreateUSDFuturesKlineWorkItem(symbol *string, kline *models.Kline) *models.KlineWorkIten {
	apiName := models.USDFuturesKline
	return models.NewKlineWorkIten(&apiName, symbol, kline, w.db.InsertKline)
}

func (w *WorkItemFactory) CreateUSDFuturesKlinesWorkItem(symbol *string, klines *[]models.Kline) *models.KlinesWorkIten {
	apiName := models.USDFuturesKline
	return models.NewKlinesWorkIten(&apiName, symbol, klines, w.db.InsertKlines)
}

func (w *WorkItemFactory) CreateUSDFuturesFundingRateWorkItem(symbol *string, fundingRate *models.FundingRate) *models.FundingRateWorkIten {
	apiName := models.USDFuturesFundingRate
	return models.NewFundingRateWorkIten(&apiName, symbol, fundingRate, w.db.InsertFundingRate)
}

func (w *WorkItemFactory) CreateUSDFuturesFundingRatesWorkItem(symbol *string, fundingRates *[]models.FundingRate) *models.FundingRatesWorkIten {
	apiName := models.USDFuturesFundingRate
	return models.NewFundingRatesWorkIten(&apiName, symbol, fundingRates, w.db.InsertFundingRates)
}
