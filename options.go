package graph

import (
	"time"
)

const (
	// QueryTimeout default duration for how long to run a query
	QueryTimeout time.Duration = time.Second * 30

	// QueryEngine default query engine to use
	QueryEngine string = "cypher"

	// StorageEngine default storage engine to use
	StorageEngine string = "memorydb"
)

// Options for the graph
type Options struct {
	Name          string
	QueryEngine   string
	StorageEngine string
	Timeout       time.Duration
}

// NewOptions creates the default graph options
func NewOptions() *Options {
	return &Options{
		QueryEngine:   QueryEngine,
		Timeout:       QueryTimeout,
		StorageEngine: StorageEngine,
	}
}
