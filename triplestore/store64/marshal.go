package store64

import (
	"fmt"

	graph "github.com/RossMerr/Caudex.Graph"
)

// Delimiter used for triplestore column
var Delimiter = '|'

const (
	// Vertex prefix used triplestore column
	Vertex = "v"
	// VertexProperties prefix used triplestore column
	VertexProperties = "p"
	// Edge prefix used triplestore column
	Edge = "e"
	// EdgeProperties prefix used triplestore column
	EdgeProperties = "k"
)

// Marshal a Vertex into triples
func Marshal(c ...*graph.Vertex) []*Triple {
	delimiter := string(Delimiter)
	tt := []*Triple{}
	for _, v := range c {
		t := &Triple{
			Row:    Vertex + delimiter + v.ID(),
			Column: v.Label(),
			Value:  float64(1),
		}
		tt = append(tt, t)

		for k, p := range v.Properties() {
			t := &Triple{
				Row:    VertexProperties + delimiter + k + delimiter + v.ID(),
				Column: fmt.Sprint(p),
				Value:  float64(1),
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &Triple{
				Row:    Edge + delimiter + e.RelationshipType() + delimiter + v.ID(),
				Column: fmt.Sprint(e.ID()),
				Value:  e.Weight,
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {

				t := &Triple{
					Row:    EdgeProperties + delimiter + k + delimiter + e.RelationshipType() + delimiter + v.ID(),
					Column: fmt.Sprint(p),
					Value:  float64(1),
				}
				tt = append(tt, t)
			}
		}
	}
	return tt
}
