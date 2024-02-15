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
		error := val.Run(ctx)
		fmt.Println("DBService: ", error)
		for error != nil {
			select {
			case <-ctx.Done():
				return
			default:
				error = val.Run(ctx)
				fmt.Println("DBService: ", error)
			}
			time.Sleep(10 * time.Second)
		}
	}
}
