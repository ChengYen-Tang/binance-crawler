package mongodb

import (
	"context"

	"github.com/ChengYen-Tang/binance-crawler/models"
	"github.com/ChengYen-Tang/binance-crawler/modules/database/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *database) CreateKlineTable(apiName *string, symbol *string, ctx context.Context) error {
	return d.CreateTable(apiName, symbol, models.KlineIndex, ctx)
}

func (d *database) CreateFundingRateTable(apiName *string, symbol *string, ctx context.Context) error {
	return d.CreateTable(apiName, symbol, models.FundingRateIndex, ctx)
}

func (d *database) CreateTable(apiName *string, symbol *string, index string, ctx context.Context) error {
	ping_error := d.Ping(ctx)
	if ping_error != nil {
		return ping_error
	}
	tableName := utils.CombineTableName(apiName, symbol)
	indexModel := mongo.IndexModel{
		Keys:    map[string]interface{}{index: 1},
		Options: options.Index().SetUnique(true),
	}
	collection := d.db.Collection(tableName)
	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
