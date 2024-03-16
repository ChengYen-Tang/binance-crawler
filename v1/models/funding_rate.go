package models

import "context"

type FundingRate struct {
	FundingRate float64 `json:"fundingRate"`
	FundingTime int64   `json:"fundingTime"`
	MarkPrice   float64 `json:"markPrice"`
}

const FundingRateIndex = "fundingtime"

type FundingRateWorkIten struct {
	apiName        *string
	symbol         *string
	FundingRate    *FundingRate
	insertFunction func(apiName *string, symbol *string, fundingRate *FundingRate, ctx context.Context) error
}

type FundingRatesWorkIten struct {
	apiName        *string
	symbol         *string
	FundingRates   *[]FundingRate
	insertFunction func(apiName *string, symbol *string, fundingRates *[]FundingRate, ctx context.Context) error
}

func NewFundingRateWorkIten(apiName *string, symbol *string, fundingRate *FundingRate, insertFunction func(apiName *string, symbol *string, fundingRate *FundingRate, ctx context.Context) error) *FundingRateWorkIten {
	return &FundingRateWorkIten{
		apiName:        apiName,
		symbol:         symbol,
		FundingRate:    fundingRate,
		insertFunction: insertFunction,
	}
}

func NewFundingRatesWorkIten(apiName *string, symbol *string, fundingRates *[]FundingRate, insertFunction func(apiName *string, symbol *string, fundingRates *[]FundingRate, ctx context.Context) error) *FundingRatesWorkIten {
	return &FundingRatesWorkIten{
		apiName:        apiName,
		symbol:         symbol,
		FundingRates:   fundingRates,
		insertFunction: insertFunction,
	}
}

func (f *FundingRateWorkIten) Run(ctx context.Context) error {
	return f.insertFunction(f.apiName, f.symbol, f.FundingRate, ctx)
}

func (f *FundingRatesWorkIten) Run(ctx context.Context) error {
	return f.insertFunction(f.apiName, f.symbol, f.FundingRates, ctx)
}
