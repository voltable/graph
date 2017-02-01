package graphs

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type GraphOperation struct {
	DB Persistence
}

var errVertexNotFound = errors.New("Vertex Not found")
var errCreatVertex = errors.New("Failed to create Vertex")

// CreateVertex creates a vetex and returns the VertexOperation.
func (g *GraphOperation) CreateVertex(i *interface{}) (*Vertex, error) {
	u1 := uuid.NewV4()
	v := Vertex{ID: u1.String(), Value: i}
	arr := &[]Vertex{v}
	if err := g.DB.Create(arr); err != nil {
		return &v, nil
	}
	return nil, errCreatVertex
}

func (g *GraphOperation) ReadVertex(ID string) (*Vertex, error) {

	if v, err := g.DB.Find(ID); err != nil {
		return v, nil
	}
	return nil, errVertexNotFound

}

func (g *GraphOperation) UpdateVertex(ID string, fn func(*Vertex) error) error {

	if v, err := g.DB.Find(ID); err != nil {
		return fn(v)
	} else {
		return err
	}
}

// DeleteVertex removes the vertex from the graph with any edges linking it
func (g *GraphOperation) DeleteVertex(ID string) error {

	if v, err := g.DB.Find(ID); err != nil {
		if len(v.edges) > 0 {
			for _, edge := range v.edges {
				for i, otherEdge := range edge.to.edges {
					if otherEdge.edge == edge.edge {
						c := make([]Edge, len(edge.to.edges)-1)
						edge.to.edges = append(append(c, edge.to.edges[:i]...), edge.to.edges[i+1:]...)
						break
					}
				}
			}
		}
		arr := &[]Vertex{*v}
		return g.DB.Delete(arr)
	}

	return errVertexNotFound

}
