package models

import "context"

type Kline struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	TradeNum                 int64  `json:"tradeNum"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

type KlineWorkIten struct {
	apiName        *string
	symbol         *string
	kline          *Kline
	insertFunction func(apiName *string, symbol *string, kline *Kline, ctx context.Context) error
}

func NewKlineWorkIten(apiName *string, symbol *string, kline *Kline, insertFunction func(apiName *string, symbol *string, kline *Kline, ctx context.Context) error) *KlineWorkIten {
	return &KlineWorkIten{
		apiName:        apiName,
		symbol:         symbol,
		kline:          kline,
		insertFunction: insertFunction,
	}
}

func (k *KlineWorkIten) Run(ctx context.Context) error {
	return k.insertFunction(k.apiName, k.symbol, k.kline, ctx)
}
