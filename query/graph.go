package query

import (
	"errors"

	"github.com/voltable/graph"
	"github.com/voltable/graph/widecolumnstore"
)

func init() {
	graph.RegisterGraph(GraphType, graph.GraphRegistration{
		NewFunc: NewGraphEngine,
	})
}

const GraphType = "graph"

var (
	errRecordNotFound = errors.New("Record Not found")
)

type Graph struct {
	Storage widecolumnstore.Storage
	Options *graph.Options
	query   QueryEngine
}

var _ graph.Graph = (*Graph)(nil)

func (s Graph) Close() {

}

// NewGraphEngine creates a new in graph engine
func NewGraphEngine(o *graph.Options) (graph.Graph, error) {
	s, err := widecolumnstore.NewStorage(o.StorageEngine)
	if err != nil {
		return nil, err
	}

	return NewGraphEngineFromStorageEngine(s, o)
}

// NewGraphEngineFromStorageEngine creates a new in graph engine with the supplied storage engine
func NewGraphEngineFromStorageEngine(s widecolumnstore.Storage, o *graph.Options) (graph.Graph, error) {
	q, err := NewQueryEngine(o.QueryEngine, s)
	if err != nil {
		return nil, err
	}

	se := Graph{
		Options: o,
		Storage: s,
		query:   q,
	}

	// queryEngine, err := NewQueryEngine(o.QueryEngine, se.storage)
	// if err != nil {
	// 	return nil, err
	// }
	// se.engine = queryEngine

	return &se, nil
}

func (s Graph) Query(str string) (*graph.Query, error) {
	return s.query.Parse(str)
}

func (s *Graph) HasPrefix(prefix []byte) widecolumnstore.Iterator {
	return s.Storage.HasPrefix(prefix)
}
