package graph

import (
	"time"
)

type StorageType string
type QueryType string

const (
	// QueryTimeout default duration for how long to run a query
	QueryTimeout time.Duration = time.Second * 30

	// QueryEngine default query engine to use
	QueryEngine QueryType = "cypher"

	// StorageEngine default storage engine to use
	StorageEngine StorageType = "memorydb"
)

// Options for the graph
type Options struct {
	Name          string
	QueryEngine   QueryType
	StorageEngine StorageType
	Timeout       time.Duration
}

// NewOptions creates the default graph options
func NewOptions(queryEngine QueryType, storageEngine StorageType) *Options {
	return &Options{
		QueryEngine:   queryEngine,
		Timeout:       QueryTimeout,
		StorageEngine: storageEngine,
	}
}
