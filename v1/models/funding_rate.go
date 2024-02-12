package models

import "context"

type FundingRateModel struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime"`
	Time        int64  `json:"time"`
}

type FundingRateWorkIten struct {
	FundingRate *FundingRateModel
}

func (f *FundingRateModel) CreateWorkItem() IWorkItem {
	return &FundingRateWorkIten{FundingRate: f}
}

func (f *FundingRateWorkIten) Run(ctx context.Context) {
	// do something
}
