package graphs

import (
	"errors"

	"github.com/hashicorp/go-uuid"
)

// GraphOperation a CRUD operation to perform over a graph
type GraphOperation struct {
	StorageEngine
}

var (
	errVertexNotFound = errors.New("Vertex Not found")
	errCreatVertex    = errors.New("Failed to create Vertex")
	errCreatVertexID  = errors.New("Failed to create Vertex ID")
)

// CreateGraphOperation builds a GraphOperation from a StorageEngine
func NewGraphOperation(p StorageEngine) *GraphOperation {
	return &GraphOperation{p}
}

// CreateVertex creates a vetex and returns the VertexOperation.
func (g *GraphOperation) CreateVertex(i interface{}) (*Vertex, error) {
	var id string
	var err error
	if id, err = uuid.GenerateUUID(); err != nil {
		return nil, errCreatVertexID
	}

	v := Vertex{ID: id, Value: i}
	if err := g.Create(v); err == nil {
		return &v, nil
	}

	return nil, errCreatVertex
}

// ReadVertex retrieves a give vertex
func (g *GraphOperation) ReadVertex(ID string) (*Vertex, error) {
	if v, err := g.Find(ID); err == nil {
		return v, nil
	}

	return nil, errVertexNotFound
}

// UpdateVertex retrieves a give vertex then lets you update it
func (g *GraphOperation) UpdateVertex(ID string, fn func(*Vertex) error) error {
	var v *Vertex
	var err error
	if v, err = g.Find(ID); err == nil {
		return fn(v)
	}

	return err
}

// DeleteVertex removes the vertex from the graph with any edges linking it
func (g *GraphOperation) DeleteVertex(ID string) error {
	if v, err := g.Find(ID); err == nil {
		v.removeRelationships()
		return g.Delete(*v)
	}

	return errVertexNotFound
}
