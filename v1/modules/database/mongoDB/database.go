package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "binance"

type database struct {
	client *mongo.Client
	db     *mongo.Database
}

func New(url string, ctx context.Context) (database, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return database{}, err
	}
	return database{client: client, db: client.Database(dbName)}, nil
}

func (d *database) Close(ctx context.Context) error {
	return d.client.Disconnect(ctx)
}

func (d *database) Ping(ctx context.Context) error {
	return d.client.Ping(ctx, nil)
}
