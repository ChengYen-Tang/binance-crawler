package crawler

// IGet is an interface for getting data from the binance api
type IGet interface {
	// GetToNow gets data from the binance api from the start time to the current time
	GetToNow(startTime *int64) error
	// Get gets data from the binance api from the start time to the end time
	Get(startTime *int64, endTime *int64) error
}
