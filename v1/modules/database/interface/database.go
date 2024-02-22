package database

import (
	"context"

	"github.com/ChengYen-Tang/binance-crawler/models"
)

type IDatabase interface {
	IsTableExist(apiName *string, symbol *string, ctx context.Context) (bool, error)
	CreateKlineTable(apiName *string, symbol *string, ctx context.Context) error
	CreateFundingRateTable(apiName *string, symbol *string, ctx context.Context) error
	GetKlineTimeRange(apiName *string, symbol *string, ctx context.Context) (firstTime *int64, lastTime *int64, error error)
	GetFundingRateTimeRange(apiName *string, symbol *string, ctx context.Context) (firstTime *int64, lastTime *int64, error error)
	InsertKline(apiName *string, symbol *string, kline *models.Kline, ctx context.Context) error
	InsertKlines(apiName *string, symbol *string, kline *[]models.Kline, ctx context.Context) error
	InsertFundingRate(apiName *string, symbol *string, fundingRate *models.FundingRate, ctx context.Context) error
	InsertFundingRates(apiName *string, symbol *string, fundingRate *[]models.FundingRate, ctx context.Context) error
	Close(ctx context.Context) error
}
