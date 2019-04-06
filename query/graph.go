package query

import (
	"errors"
	"fmt"
	"strings"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/uuid"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
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
	storage widecolumnstore.Storage
	Options *graph.Options
	query   QueryEngine
}

var _ graph.Graph = (*Graph)(nil)

func (s Graph) Close() {

}

// NewGraphEngine creates a new in memory storage engine
func NewGraphEngine(o *graph.Options) (graph.Graph, error) {
	s, err := widecolumnstore.NewStorage(o.StorageEngine)
	if err != nil {
		return nil, err
	}

	q, err := NewQueryEngine(o.QueryEngine, s)
	if err != nil {
		return nil, err
	}

	se := Graph{
		Options: o,
		storage: s,
		query:   q,
	}

	// queryEngine, err := NewQueryEngine(o.QueryEngine, se.storage)
	// if err != nil {
	// 	return nil, err
	// }
	// se.engine = queryEngine

	return &se, nil
}

// Create adds a array of vertices to the persistence
func (s Graph) Create(c ...*graph.Vertex) error {
	for _, v := range c {
		triples, transposes := MarshalKeyValue(v)
		var errstrings []string

		s.storage.Create(triples...)
		s.storage.Create(transposes...)

		if len(errstrings) > 0 {
			return fmt.Errorf(strings.Join(errstrings, "\n"))
		}
	}

	return nil
}

// Delete the array of vertices from the persistence
func (s Graph) Delete(c ...*graph.Vertex) error {
	// for _, v := range c {
	// 	triples, transposes := MarshalKeyValue(v)
	// 	var errstrings []string

	// 	for i := 0; i < len(triples); i++ {
	// 		triple := triples[i]
	// 		delete(se.tKey, string(triple.Key))

	// 		transpose := transposes[i]
	// 		delete(se.tKeyT, string(transpose.Key))
	// 	}

	// 	if len(errstrings) > 0 {
	// 		return fmt.Errorf(strings.Join(errstrings, "\n"))
	// 	}
	// }

	return nil
}

// Update the array of vertices from the persistence
func (s Graph) Update(c ...*graph.Vertex) error {
	s.Create(c...)
	return nil
}

func (s Graph) Query(str string) (*graph.Query, error) {
	return s.query.Parse(str)
}

func (s *Graph) HasPrefix(prefix []byte) widecolumnstore.Iterator {
	return s.storage.HasPrefix(prefix)
}

func (s *Graph) Edges(id uuid.UUID) widecolumnstore.Iterator {
	key := widecolumnstore.NewKey(id[:], &widecolumnstore.Column{Family: Relationship})
	prefix := key.Marshal()
	return s.storage.HasPrefix(prefix)
}
