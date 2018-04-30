package triplestore

import (
	"fmt"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/triplestore/store64"
)

var Delimiter = '|'

func Marshal(c ...*graph.Vertex) []*store64.Triple {
	tt := []*store64.Triple{}
	for _, v := range c {
		for k, p := range v.Properties() {
			t := &store64.Triple{
				Row:    k + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(p),
				Value:  float64(1),
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &store64.Triple{
				Row:    e.RelationshipType() + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(e.ID()),
				Value:  e.Weight,
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {

				t := &store64.Triple{
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
