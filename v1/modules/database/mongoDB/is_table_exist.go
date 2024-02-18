package mongodb

import (
	"context"

	"github.com/ChengYen-Tang/binance-crawler/modules/database/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d *database) IsTableExist(apiName *string, symbol *string, ctx context.Context) (bool, error) {
	ping_error := d.Ping(ctx)
	if ping_error != nil {
		return false, ping_error
	}
	tableName := utils.CombineTableName(apiName, symbol)
	collection := d.db.Collection(tableName)
	err := collection.FindOne(ctx, nil).Err()
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
