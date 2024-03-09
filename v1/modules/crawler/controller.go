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
	controller.check(models.SpotKline, controller.dbEndPoint.GetKlineTimeRange, controller.dbEndPoint.CreateKlineTable, nil, ctx)
	controller.check(models.USDFuturesKline, controller.dbEndPoint.GetKlineTimeRange, controller.dbEndPoint.CreateKlineTable, nil, ctx)
	controller.check(models.USDFuturesPremiumIndexKline, controller.dbEndPoint.GetKlineTimeRange, controller.dbEndPoint.CreateKlineTable, nil, ctx)
	controller.check(models.USDFuturesFundingRate, controller.dbEndPoint.GetFundingRateTimeRange, controller.dbEndPoint.CreateFundingRateTable, nil, ctx)
}

func (controller *Controller) check(
	apiName string,
	getTimeRange func(apiName *string, symbol *string, ctx context.Context) (*int64, *int64, error),
	createTable func(apiName *string, symbol *string, ctx context.Context) error,
	getter crawler.IGet,
	ctx context.Context) error {
	for _, symbol := range controller.params.Symbols {
		isTableExist, err := controller.dbEndPoint.IsTableExist(&apiName, &symbol, ctx)
		if err != nil {
			return err
		}

		startTime := controller.params.StartTime.UnixMilli()
		var endtime *int64

		if isTableExist {
			endtime, _, err = getTimeRange(&apiName, &symbol, ctx)
		} else {
			err = createTable(&apiName, &symbol, ctx)
		}
		if err != nil {
			return err
		}

		if endtime == nil {
			err = getter.GetToNow(&startTime)
		} else if *endtime > startTime {
			err = getter.Get(&startTime, endtime)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
