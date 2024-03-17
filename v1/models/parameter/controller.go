package parameter

import "time"

type ControllerParams struct {
	StartTime *time.Time
	Symbols   *[]string
}
