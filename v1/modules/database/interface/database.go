package database

import (
	"context"

	"github.com/ChengYen-Tang/binance-crawler/models"
)

type IDatabase interface {
	IsTableExist(apiName *string, symbol *string) (bool, error)
	CreateTable(apiName *string, symbol *string) error
	GetTableTimeRange(apiName *string, symbol *string) (firstTime *int64, lastTime *int64, error error)
	InsertKline(apiName *string, symbol *string, kline *models.Kline, ctx context.Context) error
	InsertFundingRate(apiName *string, symbol *string, fundingRate *models.FundingRate, ctx context.Context) error
}
