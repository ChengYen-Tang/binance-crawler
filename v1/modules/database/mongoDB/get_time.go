package mongodb

import (
	"context"
	"fmt"

	"github.com/ChengYen-Tang/binance-crawler/models"
	"github.com/ChengYen-Tang/binance-crawler/modules/database/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *database) GetKlineTimeRange(apiName *string, symbol *string, ctx context.Context) (firstTime *int64, lastTime *int64, error error) {
	return d.GetTableTimeRange(apiName, symbol, models.KlineIndex, models.KlineCloseTimeKey, ctx)
}

func (d *database) GetFundingRateTimeRange(apiName *string, symbol *string, ctx context.Context) (firstTime *int64, lastTime *int64, error error) {
	return d.GetTableTimeRange(apiName, symbol, models.FundingRateIndex, "", ctx)
}

func (d *database) GetTableTimeRange(apiName *string, symbol *string, index string, closeKey string, ctx context.Context) (firstTime *int64, lastTime *int64, error error) {
	maxTime := int64(0)
	minTime := int64(0)
	ping_error := d.Ping(ctx)
	if ping_error != nil {
		return &minTime, &maxTime, ping_error
	}

	var maxKey *string
	if closeKey != "" {
		maxKey = &closeKey
	} else {
		maxKey = &index
	}

	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "maxTime", Value: bson.D{{Key: "$max", Value: fmt.Sprintf("$%s", *maxKey)}}},
			{Key: "minTime", Value: bson.D{{Key: "$min", Value: fmt.Sprintf("$%s", index)}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "maxTime", Value: 1},
			{Key: "minTime", Value: 1},
		}}},
	}

	tableName := utils.CombineTableName(apiName, symbol)
	collection := d.db.Collection(tableName)
	// Execute the aggregation
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return &minTime, &maxTime, err
	}
	defer cursor.Close(ctx)

	// Decode the result
	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return &minTime, &maxTime, err
	}

	minTime = result[0]["minTime"].(int64)
	maxTime = result[0]["maxTime"].(int64)
	// Return the result
	return &minTime, &maxTime, nil
}

func (d *database) GetKlineLastTime(apiName *string, symbol *string, ctx context.Context) (*int64, error) {
	return d.GetTableLastTime(apiName, symbol, models.KlineIndex, ctx)
}

func (d *database) GetFundingRateLastTime(apiName *string, symbol *string, ctx context.Context) (*int64, error) {
	return d.GetTableLastTime(apiName, symbol, models.FundingRateIndex, ctx)
}

func (d *database) GetTableLastTime(apiName *string, symbol *string, index string, ctx context.Context) (*int64, error) {
	maxOpenTime := int64(0)
	ping_error := d.Ping(ctx)
	if ping_error != nil {
		return &maxOpenTime, ping_error
	}

	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "maxOpenTime", Value: bson.D{{Key: "$max", Value: fmt.Sprintf("$%s", index)}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "maxOpenTime", Value: 1},
		}}},
	}

	tableName := utils.CombineTableName(apiName, symbol)
	collection := d.db.Collection(tableName)
	// Execute the aggregation
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return &maxOpenTime, err
	}
	defer cursor.Close(ctx)

	// Decode the result
	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return &maxOpenTime, err
	}

	maxOpenTime = result[0]["maxOpenTime"].(int64)
	// Return the result
	return &maxOpenTime, nil
}
