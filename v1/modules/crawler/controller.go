package crawler

import (
	"context"

	"github.com/ChengYen-Tang/binance-crawler/factory"
	models "github.com/ChengYen-Tang/binance-crawler/models"
	iModel "github.com/ChengYen-Tang/binance-crawler/models/interface"
	"github.com/ChengYen-Tang/binance-crawler/models/parameter"
	crawler "github.com/ChengYen-Tang/binance-crawler/modules/crawler/interface"
	"github.com/ChengYen-Tang/binance-crawler/modules/database"
)

// Controller is a struct for controlling the crawler
type Controller struct {
	dbEndPoint *database.DBEndpoints
	client     *crawler.IClient
	channel    *chan iModel.IWorkItem
	factory    *factory.WorkItemFactory
	getters    *[]crawler.IGet
	params     *parameter.ControllerParams
}

// New creates a new instance of Controller
func NewController(dbEndPoint *database.DBEndpoints, client *crawler.IClient, channel *chan iModel.IWorkItem, factory *factory.WorkItemFactory, params *parameter.ControllerParams) *Controller {
	getters := &[]crawler.IGet{}

	return &Controller{
		dbEndPoint: dbEndPoint,
		client:     client,
		channel:    channel,
		factory:    factory,
		getters:    getters,
		params:     params,
	}
}

func (controller *Controller) Run(ctx context.Context) {
	controller.check(ctx)
}

func (controller *Controller) check(ctx context.Context) error {
	apiName := models.SpotKline
	for _, symbol := range controller.params.Symbols {
		isTableExist, err := controller.dbEndPoint.IsTableExist(&apiName, &symbol, ctx)
		if err != nil {
			return err
		}

		var (
			firstTime *int64
			lastTime  *int64
		)

		if isTableExist {
			firstTime, lastTime, err = controller.dbEndPoint.GetKlineTimeRange(&apiName, &symbol, ctx)
		} else {
			err = controller.dbEndPoint.CreateKlineTable(&apiName, &symbol, ctx)
			startTime := controller.params.StartTime.UnixMilli()
			firstTime = &startTime
		}
		if err != nil {
			return err
		}
	}
	apiName = models.USDFuturesKline
	for _, symbol := range controller.params.Symbols {
		isTableExist, err := controller.dbEndPoint.IsTableExist(&apiName, &symbol, ctx)
		if err != nil {
			return err
		}

		var (
			firstTime *int64
			lastTime  *int64
		)

		if isTableExist {
			firstTime, lastTime, err = controller.dbEndPoint.GetKlineTimeRange(&apiName, &symbol, ctx)
		} else {
			err = controller.dbEndPoint.CreateKlineTable(&apiName, &symbol, ctx)
			startTime := controller.params.StartTime.UnixMilli()
			firstTime = &startTime
		}
		if err != nil {
			return err
		}
	}
	apiName = models.USDFuturesPremiumIndexKline
	for _, symbol := range controller.params.Symbols {
		isTableExist, err := controller.dbEndPoint.IsTableExist(&apiName, &symbol, ctx)
		if err != nil {
			return err
		}

		var (
			firstTime *int64
			lastTime  *int64
		)

		if isTableExist {
			firstTime, lastTime, err = controller.dbEndPoint.GetKlineTimeRange(&apiName, &symbol, ctx)
			if err != nil {
				return err
			}
		} else {
			err = controller.dbEndPoint.CreateKlineTable(&apiName, &symbol, ctx)
			startTime := controller.params.StartTime.UnixMilli()
			firstTime = &startTime
		}
		if err != nil {
			return err
		}
	}
	apiName = models.USDFuturesFundingRate
	for _, symbol := range controller.params.Symbols {
		isTableExist, err := controller.dbEndPoint.IsTableExist(&apiName, &symbol, ctx)
		if err != nil {
			return err
		}

		var (
			firstTime *int64
			lastTime  *int64
		)

		if isTableExist {
			firstTime, lastTime, err = controller.dbEndPoint.GetFundingRateTimeRange(&apiName, &symbol, ctx)
			if err != nil {
				return err
			}
		} else {
			err = controller.dbEndPoint.CreateFundingRateTable(&apiName, &symbol, ctx)
			startTime := controller.params.StartTime.UnixMilli()
			firstTime = &startTime
		}
		if err != nil {
			return err
		}
	}

	return nil
}
