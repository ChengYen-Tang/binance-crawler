package parameter

import "time"

type Configuration struct {
	DbConnectionString string
	StartTime          time.Time
	Symbols            []string
}
