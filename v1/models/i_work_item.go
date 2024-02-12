package models

import "context"

// IWorkItem is an interface for the WorkItem struct
type IWorkItem interface {
	Run(ctx context.Context)
}
