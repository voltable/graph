package graphs

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

// GraphOperation a CRUD operation to perform over a graph
type GraphOperation struct {
	DB Persistence
}

var (
	errVertexNotFound = errors.New("Vertex Not found")
	errCreatVertex    = errors.New("Failed to create Vertex")
)

// CreateVertex creates a vetex and returns the VertexOperation.
func (g *GraphOperation) CreateVertex(i *interface{}) (*Vertex, error) {
	u1 := uuid.NewV4()
	v := Vertex{ID: u1.String(), Value: i}
	arr := []Vertex{v}
	if err := g.DB.Create(arr); err != nil {
		return &v, nil
	}
	return nil, errCreatVertex
}

// ReadVertex retrieves a give vertex
func (g *GraphOperation) ReadVertex(ID string) (*Vertex, error) {

	if v, err := g.DB.Find(ID); err != nil {
		return v, nil
	}
	return nil, errVertexNotFound

}

// UpdateVertex retrieves a give vertex then lets you update it
func (g *GraphOperation) UpdateVertex(ID string, fn func(*Vertex) error) error {

	var v *Vertex
	var err error
	if v, err = g.DB.Find(ID); err != nil {
		return fn(v)
	}
	return err
}

// DeleteVertex removes the vertex from the graph with any edges linking it
func (g *GraphOperation) DeleteVertex(ID string) error {

	if v, err := g.DB.Find(ID); err != nil {
		for _, edge := range v.edges {
			edge.removeTo()
		}

		arr := []Vertex{*v}
		return g.DB.Delete(arr)
	}

	return errVertexNotFound

}
