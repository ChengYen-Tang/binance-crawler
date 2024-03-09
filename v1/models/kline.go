package models

import "context"

type Kline struct {
	OpenTime                 int64   `json:"openTime"`
	Open                     float64 `json:"open"`
	High                     float64 `json:"high"`
	Low                      float64 `json:"low"`
	Close                    float64 `json:"close"`
	Volume                   float64 `json:"volume"`
	CloseTime                int64   `json:"closeTime"`
	QuoteAssetVolume         float64 `json:"quoteAssetVolume"`
	TradeNum                 int64   `json:"tradeNum"`
	TakerBuyBaseAssetVolume  float64 `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume float64 `json:"takerBuyQuoteAssetVolume"`
}

const KlineIndex = "opentime"

type KlineWorkIten struct {
	apiName        *string
	symbol         *string
	kline          *Kline
	insertFunction func(apiName *string, symbol *string, kline *Kline, ctx context.Context) error
}

type KlinesWorkIten struct {
	apiName        *string
	symbol         *string
	klines         *[]Kline
	insertFunction func(apiName *string, symbol *string, klines *[]Kline, ctx context.Context) error
}

func NewKlineWorkIten(apiName *string, symbol *string, kline *Kline, insertFunction func(apiName *string, symbol *string, kline *Kline, ctx context.Context) error) *KlineWorkIten {
	return &KlineWorkIten{
		apiName:        apiName,
		symbol:         symbol,
		kline:          kline,
		insertFunction: insertFunction,
	}
}

func NewKlinesWorkIten(apiName *string, symbol *string, klines *[]Kline, insertFunction func(apiName *string, symbol *string, klines *[]Kline, ctx context.Context) error) *KlinesWorkIten {
	return &KlinesWorkIten{
		apiName:        apiName,
		symbol:         symbol,
		klines:         klines,
		insertFunction: insertFunction,
	}
}

func (k *KlineWorkIten) Run(ctx context.Context) error {
	return k.insertFunction(k.apiName, k.symbol, k.kline, ctx)
}

func (k *KlinesWorkIten) Run(ctx context.Context) error {
	return k.insertFunction(k.apiName, k.symbol, k.klines, ctx)
}
