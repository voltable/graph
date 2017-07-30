package graph

// Options for the graph
type Options struct {
	Name        string
	QueryEngine string
}

// NewOptions creates the default graph options
func NewOptions() *Options {
	return &Options{QueryEngine: "Cypher"}
}
