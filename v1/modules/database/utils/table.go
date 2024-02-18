package utils

func CombineTableName(apiName *string, symbol *string) string {
	return *apiName + "_" + *symbol
}
