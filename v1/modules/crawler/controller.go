package crawler

import (
	"github.com/ChengYen-Tang/binance-crawler/models"
)

// Controller is a struct for controlling the crawler
type Controller struct {
	client  *IClient
	channel *chan *models.IWorkItem
	getters *[]IGet
}

// New creates a new instance of Controller
func (c *Controller) New(client *IClient) *Controller {
	channel := make(chan *models.IWorkItem, 1000*10)
	getters := &[]IGet{}

	return &Controller{
		client:  client,
		channel: &channel,
		getters: getters,
	}
}
