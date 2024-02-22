package utils

import "github.com/ChengYen-Tang/binance-crawler/models"

func CombineTableName(apiName *string, symbol *string) string {
	return *apiName + "_" + *symbol
}

func RemoveDuplicateKlinesByIndex(klines *[]models.Kline) *[]models.Kline {
	keys := make(map[int64]bool)
	list := []models.Kline{}
	for _, entry := range *klines {
		if _, value := keys[entry.OpenTime]; !value {
			keys[entry.OpenTime] = true
			list = append(list, entry)
		}
	}
	return &list
}

func RemoveDuplicateFundingRatesByIndex(fundingRates *[]models.FundingRate) *[]models.FundingRate {
	keys := make(map[int64]bool)
	list := []models.FundingRate{}
	for _, entry := range *fundingRates {
		if _, value := keys[entry.FundingTime]; !value {
			keys[entry.FundingTime] = true
			list = append(list, entry)
		}
	}
	return &list
}
