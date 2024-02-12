package crawler

import (
	"time"

	"github.com/ChengYen-Tang/binance-crawler/models"
)

type IClient interface {
}

// IGet is an interface for getting data from the binance api
type IGet interface {
	// New creates a new instance of IGet
	New(*IClient, *chan models.IWorkItem) IGet
	// GetToNow gets data from the binance api from the start time to the current time
	GetToNow(time.Time)
	// Get gets data from the binance api from the start time to the end time
	Get(time.Time, time.Time)
}
