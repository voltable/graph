package triplestore

import (
	"encoding/binary"
	"fmt"
	"math"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/container/triples"
)

var Delimiter = '|'

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func Marshal(c ...*graph.Vertex) []*triples.Triple {
	tt := []*triples.Triple{}
	for _, v := range c {
		for k, p := range v.Properties() {
			a := &triples.Any{
				Value: float64ToByte(float64(1)),
				Type:  "float64",
			}
			t := &triples.Triple{
				Row:    k + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(p),
				Value:  a,
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			a := &triples.Any{
				Value: float64ToByte(float64(e.Weight)),
				Type:  "float64",
			}

			t := &triples.Triple{
				Row:    e.RelationshipType() + string(Delimiter) + v.ID(),
				Column: fmt.Sprint(e.ID()),
				Value:  a,
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {

				a := &triples.Any{
					Value: float64ToByte(float64(1)),
					Type:  "float64",
				}

				t := &triples.Triple{
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
