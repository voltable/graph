package memorydb

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/traversal"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func init() {
	graph.RegisterGraph(GraphType, graph.GraphRegistration{
		NewFunc: newStorageEngine,
	})
}

const GraphType = "memory"

var (
	errRecordNotFound = errors.New("Record Not found")
)

type StorageEngine struct {
	vertices map[string]vertices.Vertex
	Options  *graph.Options
}

func (se *StorageEngine) Close() {

}

func newStorageEngine(o *graph.Options) (graph.Graph, error) {
	se := StorageEngine{Options: o, vertices: make(map[string]vertices.Vertex)}
	return &se, nil
}

// Create adds a array of vertices to the persistence
func (se *StorageEngine) Create(c ...*vertices.Vertex) error {
	for _, v := range c {
		se.vertices[v.ID()] = *v
	}

	return nil
}

// Delete the array of vertices from the persistence
func (se *StorageEngine) Delete(c ...*vertices.Vertex) error {
	for _, v := range c {
		delete(se.vertices, v.ID())
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

func (se *StorageEngine) Query() *query.VertexPath {
	return nil
}

func (se *StorageEngine) QueryRoot(predicate func(*vertices.Vertex) bool) *query.Query {
	return traversal.QueryRoot(predicate)
}
