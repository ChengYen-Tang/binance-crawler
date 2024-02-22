package mongodb

import (
	"context"
	"strings"

	"github.com/ChengYen-Tang/binance-crawler/models"
	"github.com/ChengYen-Tang/binance-crawler/modules/database/utils"
)

func (d *database) InsertKline(apiName *string, symbol *string, kline *models.Kline, ctx context.Context) error {
	return d.InsertData(apiName, symbol, kline, ctx)
}

func (d *database) InsertFundingRate(apiName *string, symbol *string, fundingRate *models.FundingRate, ctx context.Context) error {
	return d.InsertData(apiName, symbol, fundingRate, ctx)
}

func (d *database) InsertData(apiName *string, symbol *string, data interface{}, ctx context.Context) error {
	ping_error := d.Ping(ctx)
	if ping_error != nil {
		return ping_error
	}
	tableName := utils.CombineTableName(apiName, symbol)
	collection := d.db.Collection(tableName)
	_, err := collection.InsertOne(ctx, data)
	if err != nil && strings.Contains(err.Error(), "E11000") {
		return nil
	}
	return err
}

func (d *database) InsertKlines(apiName *string, symbol *string, klines *[]models.Kline, ctx context.Context) error {
	klines = utils.RemoveDuplicateKlinesByIndex(klines)
	interfaceSlice := make([]interface{}, len(*klines))
	for i, v := range *klines {
		interfaceSlice[i] = v
	}
	return d.InsertManyData(apiName, symbol, interfaceSlice, ctx)
}

func (d *database) InsertFundingRates(apiName *string, symbol *string, fundingRate *[]models.FundingRate, ctx context.Context) error {
	fundingRate = utils.RemoveDuplicateFundingRatesByIndex(fundingRate)
	interfaceSlice := make([]interface{}, len(*fundingRate))
	for i, v := range *fundingRate {
		interfaceSlice[i] = v
	}
	return d.InsertManyData(apiName, symbol, interfaceSlice, ctx)
}

func (d *database) InsertManyData(apiName *string, symbol *string, data []interface{}, ctx context.Context) error {
	ping_error := d.Ping(ctx)
	if ping_error != nil {
		return ping_error
	}
	tableName := utils.CombineTableName(apiName, symbol)
	collection := d.db.Collection(tableName)
	_, err := collection.InsertMany(ctx, data)
	if err != nil && strings.Contains(err.Error(), "E11000") {
		return d.insertManyDataUseInsertOne(apiName, symbol, data, ctx)
	}
	return err
}

func (d *database) insertManyDataUseInsertOne(apiName *string, symbol *string, data []interface{}, ctx context.Context) error {
	for _, v := range data {
		err := d.InsertData(apiName, symbol, v, ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
