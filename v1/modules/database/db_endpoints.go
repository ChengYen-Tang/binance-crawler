package database

import (
	"context"

	database "github.com/ChengYen-Tang/binance-crawler/modules/database/interface"
)

type DBEndpoints struct {
	dbInstance database.IDatabase
}

func NewEndpoints(dbInstance database.IDatabase) *DBEndpoints {
	return &DBEndpoints{
		dbInstance: dbInstance,
	}
}

func (endpoint *DBEndpoints) IsTableExist(apiName *string, symbol *string, ctx context.Context) (bool, error) {
	return endpoint.dbInstance.IsTableExist(apiName, symbol, ctx)
}

func (endpoint *DBEndpoints) CreateKlineTable(apiName *string, symbol *string, ctx context.Context) error {
	return endpoint.dbInstance.CreateKlineTable(apiName, symbol, ctx)
}

func (endpoint *DBEndpoints) CreateFundingRateTable(apiName *string, symbol *string, ctx context.Context) error {
	return endpoint.dbInstance.CreateFundingRateTable(apiName, symbol, ctx)
}

func (endpoint *DBEndpoints) GetKlineTimeRange(apiName *string, symbol *string, ctx context.Context) (firstTime *int64, lastTime *int64, error error) {
	return endpoint.dbInstance.GetKlineTimeRange(apiName, symbol, ctx)
}

func (endpoint *DBEndpoints) GetFundingRateTimeRange(apiName *string, symbol *string, ctx context.Context) (firstTime *int64, lastTime *int64, error error) {
	return endpoint.dbInstance.GetFundingRateTimeRange(apiName, symbol, ctx)
}

func (endpoint *DBEndpoints) GetKlineLastTime(apiName *string, symbol *string, ctx context.Context) (*int64, error) {
	return endpoint.dbInstance.GetKlineLastTime(apiName, symbol, ctx)
}

func (endpoint *DBEndpoints) GetFundingRateLastTime(apiName *string, symbol *string, ctx context.Context) (*int64, error) {
	return endpoint.dbInstance.GetFundingRateLastTime(apiName, symbol, ctx)
}
