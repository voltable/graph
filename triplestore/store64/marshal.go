package store64

import (
	"fmt"

	graph "github.com/RossMerr/Caudex.Graph"
)

var Delimiter = '|'

func Marshal(c ...*graph.Vertex) []*Triple {
	tt := []*Triple{}
	for _, v := range c {
		for k, p := range v.Properties() {
			t := &Triple{
				Row:    k + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(p),
				Value:  float64(1),
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &Triple{
				Row:    e.RelationshipType() + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(e.ID()),
				Value:  e.Weight,
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {

				t := &Triple{
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
