// Package graph provides instance of the caudex.graph database.
package graph

//Graph structure for saving graph data
type Graph interface {
	Close()
	Query(string) (*Query, error)
}
