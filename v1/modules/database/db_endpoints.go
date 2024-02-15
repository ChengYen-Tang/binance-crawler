package database

import database "github.com/ChengYen-Tang/binance-crawler/modules/database/interface"

type DBEndpoints struct {
	dbInstance database.IDatabase
}

func NewEndpoints(dbInstance database.IDatabase) *DBEndpoints {
	return &DBEndpoints{
		dbInstance: dbInstance,
	}
}

func (endpoint *DBEndpoints) IsTableExist(apiName *string, symbol *string) (bool, error) {
	return endpoint.dbInstance.IsTableExist(apiName, symbol)
}

func (endpoint *DBEndpoints) CreateTable(apiName *string, symbol *string) error {
	return endpoint.dbInstance.CreateTable(apiName, symbol)
}

func (endpoint *DBEndpoints) GetTableTimeRange(apiName *string, symbol *string) (firstTime *int64, lastTime *int64, error error) {
	return endpoint.dbInstance.GetTableTimeRange(apiName, symbol)
}
