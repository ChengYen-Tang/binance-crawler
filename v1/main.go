package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ChengYen-Tang/binance-crawler/factory"
	iModel "github.com/ChengYen-Tang/binance-crawler/models/interface"
	"github.com/ChengYen-Tang/binance-crawler/models/parameter"
	"github.com/ChengYen-Tang/binance-crawler/modules/crawler"
	go_binance "github.com/ChengYen-Tang/binance-crawler/modules/crawler/go-binance"
	"github.com/ChengYen-Tang/binance-crawler/modules/database"
	mongodb "github.com/ChengYen-Tang/binance-crawler/modules/database/mongoDB"
)

func main() {
	configuration, configErr := LoadConfig()
	if configErr != nil {
		fmt.Println("error:", *configErr)
		return
	}

	ctx := context.Background()
	db, err := mongodb.New(configuration.DbConnectionString, ctx)
	if err != nil {
		panic(err)
	}
	endpoints := database.NewEndpoints(db)
	client := go_binance.NewClient("", "")
	channel := make(chan iModel.IWorkItem, 10000)
	factory := factory.NewWorkItemFactory(db)
	controllerParams := parameter.ControllerParams{
		StartTime: &configuration.StartTime,
		Symbols:   &configuration.Symbols,
	}

	controller := crawler.NewController(endpoints, client, &channel, factory, &controllerParams)
	dbService := database.NewService(&channel)

	go dbService.Run(ctx)
	controller.Check(ctx)
}

func LoadConfig() (*parameter.Configuration, *error) {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := parameter.Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return nil, &err
	}
	return &configuration, nil
}
