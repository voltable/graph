package triplestore

import (
	"fmt"

	graph "github.com/RossMerr/Caudex.Graph"
)

var Delimiter = '|'

func Marshal(c ...*graph.Vertex) []*Triple {
	tt := []*Triple{}
	for _, v := range c {
		for k, p := range v.Properties() {
			a, _ := NewAny(float64(1))
			t := &Triple{
				Row:    k + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(p),
				Value:  a,
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			a, _ := NewAny(e.Weight)
			t := &Triple{
				Row:    e.RelationshipType() + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(e.ID()),
				Value:  a,
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {

				a, _ := NewAny(float64(1))
				t := &Triple{
					Row:    k + string(Delimiter) + e.RelationshipType() + string(Delimiter) + v.ID(),
					Column: fmt.Sprint(p),
					Value:  a,
				}
				tt = append(tt, t)
			}
		}
	}
	return tt
}
