package graphs

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

// GraphOperation a CRUD operation to perform over a graph
type GraphOperation struct {
	db StorageEngine
}

var (
	errVertexNotFound = errors.New("Vertex Not found")
	errCreatVertex    = errors.New("Failed to create Vertex")
)

// CreateGraphOperation builds a GraphOperation from a StorageEngine
func CreateGraphOperation(p StorageEngine) *GraphOperation {
	return &GraphOperation{db: p}
}

// CreateVertex creates a vetex and returns the VertexOperation.
func (g *GraphOperation) CreateVertex(i *interface{}) (*Vertex, error) {
	u1 := uuid.NewV4()
	v := Vertex{ID: u1.String(), Value: i}
	arr := []Vertex{v}
	if err := g.db.Create(arr); err != nil {
		return &v, nil
	}
	return nil, errCreatVertex
}

// ReadVertex retrieves a give vertex
func (g *GraphOperation) ReadVertex(ID string) (*Vertex, error) {

	if v, err := g.db.Find(ID); err != nil {
		return v, nil
	}
	return nil, errVertexNotFound

}

// UpdateVertex retrieves a give vertex then lets you update it
func (g *GraphOperation) UpdateVertex(ID string, fn func(*Vertex) error) error {

	var v *Vertex
	var err error
	if v, err = g.db.Find(ID); err != nil {
		return fn(v)
	}
	return err
}

// DeleteVertex removes the vertex from the graph with any edges linking it
func (g *GraphOperation) DeleteVertex(ID string) error {

	if v, err := g.db.Find(ID); err != nil {
		v.removeRelationships()
		arr := []Vertex{*v}
		return g.db.Delete(arr)
	}

	return errVertexNotFound

}
