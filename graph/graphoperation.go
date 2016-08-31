package caudex

import "github.com/satori/go.uuid"

type GraphOperation struct {
	db     *Graph
	add    map[string]Vertex
	change map[string]Vertex
	delete map[string]Vertex
}

// CreateVertex creates a vetex and returns the new vertex.
func (o *GraphOperation) CreateVertex() (*VertexOperation, *Vertex) {
	u1 := uuid.NewV4()
	vertex := Vertex{ID: u1.String(), Value: new(interface{})}
	o.add[u1.String()] = vertex
	op := VertexOperation{v: &vertex}
	return &op, &vertex
}

// RemoveVertex removes the vertex from the graph with any edges linking it
func (o *GraphOperation) RemoveVertex(v *Vertex) {
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
	o.delete[v.ID] = *v
}
