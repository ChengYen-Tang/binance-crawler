package mongodb

import (
	"context"

	"github.com/ChengYen-Tang/binance-crawler/modules/database/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *database) IsTableExist(apiName *string, symbol *string, ctx context.Context) (bool, error) {
	ping_error := d.Ping(ctx)
	if ping_error != nil {
		return false, ping_error
	}
	tableName := utils.CombineTableName(apiName, symbol)
	collections, err := d.db.ListCollectionNames(ctx, bson.D{{}})
	exists := false
	if err != nil {
		return false, err
	}
	for _, collection := range collections {
		if collection == tableName {
			exists = true
			break
		}
	}
	return exists, nil
}
