package graph

import (
	"time"
)

const (
	// QueryTimeout default duration for how long to run a query
	QueryTimeout time.Duration = time.Second * 30

	// QueryEngine default query engine to use
	QueryEngine string = "Cypher"
)

// Options for the graph
type Options struct {
	Name        string
	QueryEngine string
	Timeout     time.Duration
}

// NewOptions creates the default graph options
func NewOptions() *Options {
	return &Options{QueryEngine: QueryEngine, Timeout: QueryTimeout}
}
