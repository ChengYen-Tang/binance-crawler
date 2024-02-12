package go_binance

import (
	"github.com/ChengYen-Tang/binance-crawler/modules/crawler"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

type BinanceClient struct {
	SpotClient    *binance.Client
	FuturesClient *futures.Client
}

// NewClient creates a new instance of IClient
func NewClient(apiKey, secretKey string) crawler.IClient {
	return &BinanceClient{
		SpotClient:    binance.NewClient("", ""),
		FuturesClient: futures.NewClient("", ""),
	}
}
