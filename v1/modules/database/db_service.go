package database

import (
	"context"
	"fmt"
	"time"

	models "github.com/ChengYen-Tang/binance-crawler/models/interface"
)

type DBService struct {
	channel *chan models.IWorkItem
}

func NewService(channel *chan models.IWorkItem) *DBService {
	return &DBService{
		channel: channel,
	}
}

func (service *DBService) Run(ctx context.Context) {
	for val := range *service.channel {
		select {
		case <-ctx.Done():
			return
		default:
			service.retry(ctx, val)
		}
	}
}

func (service *DBService) retry(ctx context.Context, val models.IWorkItem) {
	error := val.Run(ctx)
	for error != nil {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("DBService error:", error)
			fmt.Println("Waiting for 10 seconds, then retry...")
			time.Sleep(10 * time.Second)
			error = val.Run(ctx)
		}
	}
}
