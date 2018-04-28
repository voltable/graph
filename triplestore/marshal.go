package triplestore

import (
	"fmt"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/container/triples"
)

var Delimiter = '|'

func Marshal(c ...*graph.Vertex) []*triples.Triple {
	tt := []*triples.Triple{}
	for _, v := range c {
		for k, p := range v.Properties() {
			t := &triples.Triple{
				Row:    k + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(p),
				Value:  float64(1),
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &triples.Triple{
				Row:    e.RelationshipType() + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(e.ID()),
				Value:  float64(e.Weight),
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {
				t := &triples.Triple{
					Row:    k + string(Delimiter) + e.RelationshipType() + string(Delimiter) + v.ID(),
					Column: fmt.Sprint(p),
					Value:  float64(1),
				}
				tt = append(tt, t)
			}
		}
	}
	return tt
}
