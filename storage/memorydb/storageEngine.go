package memorydb

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func init() {
	graph.RegisterGraph(GraphType, graph.GraphRegistration{
		NewFunc: NewStorageEngine,
	})
}

const GraphType = "memory"

var (
	errRecordNotFound = errors.New("Record Not found")
)

type StorageEngine struct {
	vertices    map[string]vertices.Vertex
	keys        []string
	Options     *graph.Options
	queryEngine query.Engine
	traversal   query.Traversal
}

func (se *StorageEngine) Close() {

}

// NewStorageEngine creates anew in memory storage engine
func NewStorageEngine(o *graph.Options) (graph.Graph, error) {
	queryEngine, err := query.NewQueryEngine(o.QueryEngine)
	if err != nil {
		return nil, err
	}
	se := StorageEngine{
		Options:     o,
		vertices:    make(map[string]vertices.Vertex),
		queryEngine: queryEngine}

	se.traversal = *query.NewTraversal(se.Find)
	return &se, nil
}

// Create adds a array of vertices to the persistence
func (se *StorageEngine) Create(c ...*vertices.Vertex) error {
	for _, v := range c {
		se.vertices[v.ID()] = *v
		se.keys = append(se.keys, v.ID())
	}

	return nil
}

// Delete the array of vertices from the persistence
func (se *StorageEngine) Delete(c ...*vertices.Vertex) error {
	for _, v := range c {
		delete(se.vertices, v.ID())
		for i, k := range se.keys {
			if k == v.ID() {
				se.keys = append(se.keys[:i], se.keys[i+1:]...)
				break

			}
		}
		//delete(se.keys, v.ID())
	}

	return nil
}

// Find a vertex from the persistence
func (se *StorageEngine) Find(ID string) (*vertices.Vertex, error) {
	if v, ok := se.vertices[ID]; ok {
		return &v, nil
	} else {
		return nil, errRecordNotFound
	}
}

// Update the array of vertices from the persistence
func (se *StorageEngine) Update(c ...*vertices.Vertex) error {
	se.Create(c...)
	return nil
}

func (se *StorageEngine) Query(str string) (*query.Query, error) {
	path, err := se.queryEngine.Parser(str)

	if err != nil {
		return nil, err
	}

	q := query.NewQuery(path, str)

	// should do something clever to pick the right index not just iterate
	se.traversal.Travers(se.forEach(), q)
	return q, err
}

func (se *StorageEngine) forEach() func() query.Iterator {
	return func() query.Iterator {
		position := 0
		length := len(se.keys)
		return func() (item interface{}, ok bool) {
			if position < length {
				key := se.keys[position]
				v := se.vertices[key]
				position = position + 1
				return query.NewFrontier(&v), true
			}
			return nil, false
		}
	}
}
