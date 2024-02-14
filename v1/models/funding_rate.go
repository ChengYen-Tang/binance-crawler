package models

import "context"

type FundingRate struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime"`
	Time        int64  `json:"time"`
}

type FundingRateWorkIten struct {
	apiName        *string
	symbol         *string
	FundingRate    *FundingRate
	insertFunction func(apiName *string, symbol *string, fundingRate *FundingRate, ctx context.Context) error
}

func NewFundingRateWorkIten(apiName *string, symbol *string, fundingRate *FundingRate, insertFunction func(apiName *string, symbol *string, fundingRate *FundingRate, ctx context.Context) error) *FundingRateWorkIten {
	return &FundingRateWorkIten{
		apiName:        apiName,
		symbol:         symbol,
		FundingRate:    fundingRate,
		insertFunction: insertFunction,
	}
}

func (f *FundingRateWorkIten) Run(ctx context.Context) error {
	return f.insertFunction(f.apiName, f.symbol, f.FundingRate, ctx)
}
