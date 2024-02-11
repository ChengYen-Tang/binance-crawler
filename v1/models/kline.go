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
	Kline *Kline
}

func (k *Kline) CreateSpotKlineWorkItem() IWorkItem {
	return &KlineWorkIten{Kline: k}
}

func (k *Kline) CreateFuturesKlineWorkItem() IWorkItem {
	return &KlineWorkIten{Kline: k}
}

func (k *Kline) CreatePremiumIndexKlineWorkItem() IWorkItem {
	return &KlineWorkIten{Kline: k}
}

func (k *KlineWorkIten) Run(ctx context.Context) {
	// do something
}
