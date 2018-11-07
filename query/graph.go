package query

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph"
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

func (s *Graph) Close() {

}

// NewGraphEngine creates anew in memory storage engine
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
func (s *Graph) Create(c ...*graph.Vertex) error {
	// for _, v := range c {
	// 	triples, transposes := MarshalKeyValue(v)
	// 	var errstrings []string

	// 	for i := 0; i < len(triples); i++ {
	// 		triple := triples[i]
	// 		se.tKeyIndex[len(se.tKey)] = string(triple.Key)
	// 		se.tKey[string(triple.Key)] = triple.Value

	// 		transpose := transposes[i]
	// 		se.tKeyTIndex[len(se.tKeyT)] = string(transpose.Key)
	// 		se.tKeyT[string(transpose.Key)] = transpose.Value
	// 	}

	// 	if len(errstrings) > 0 {
	// 		return fmt.Errorf(strings.Join(errstrings, "\n"))
	// 	}
	// }

	return nil
}

// Delete the array of vertices from the persistence
func (s *Graph) Delete(c ...*graph.Vertex) error {
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
func (s *Graph) Update(c ...*graph.Vertex) error {
	s.Create(c...)
	return nil
}

func (s *Graph) Query(str string) (*graph.Query, error) {
	return s.query.Parse(str)
}

func (s *Graph) HasPrefix(prefix []byte) widecolumnstore.Iterator {
	return s.storage.HasPrefix(prefix)
}
