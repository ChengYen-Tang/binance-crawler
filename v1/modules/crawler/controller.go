package crawler

import (
	"github.com/ChengYen-Tang/binance-crawler/factory"
	models "github.com/ChengYen-Tang/binance-crawler/models/interface"
	crawler "github.com/ChengYen-Tang/binance-crawler/modules/crawler/interface"
)

// Controller is a struct for controlling the crawler
type Controller struct {
	client  *crawler.IClient
	channel *chan models.IWorkItem
	factory *factory.WorkItemFactory
	getters *[]crawler.IGet
}

// New creates a new instance of Controller
func NewController(client *crawler.IClient, channel *chan models.IWorkItem, factory *factory.WorkItemFactory) *Controller {
	getters := &[]crawler.IGet{}

	return &Controller{
		client:  client,
		channel: channel,
		factory: factory,
		getters: getters,
	}
}
