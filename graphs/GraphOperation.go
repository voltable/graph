package graphs

import (
	"bitbucket.org/rossmerr/caudex/graphs/internal"
	uuid "github.com/satori/go.uuid"
)

type GraphOperation struct {
	vertices chan VertexEnvelop
	result   chan *Vertex
	find     chan string
}

func BuildGraphOperation(c chan VertexEnvelop) *GraphOperation {
	o := GraphOperation{vertices: c}
	return &o
}

// CreateVertex creates a vetex and returns the new vertex.
func (g *GraphOperation) CreateVertex(i *interface{}) *VertexOperation {
	u1 := uuid.NewV4()
	v := Vertex{ID: u1.String(), Value: i}
	array := []Vertex{v}
	ve := VertexEnvelop{State: internal.Add, Vertices: &array}
	g.vertices <- ve
	op := VertexOperation{v: &v}
	return &op
}

func (g *GraphOperation) findVertex(ID string) *Vertex {
	g.find <- ID

	for v := range g.result {
		if v.ID == ID {
			return v
		}
	}

	return nil
}

func (g *GraphOperation) UpdateVertex(ID string, fn func(*VertexOperation) error) {
	//g.findVertex()

}

// RemoveVertex removes the vertex from the graph with any edges linking it
func (g *GraphOperation) RemoveVertex(v *Vertex) {
	if v == nil {
		return
	}

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

}
