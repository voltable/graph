package keyValue

import (
	"fmt"

	graph "github.com/RossMerr/Caudex.Graph"
)

const (
	// Delimiter used for KeyValue column
	delimiter    = "|"
	label        = "l"
	properties   = "p"
	relationship = "r"
)

// Marshal a Vertex into triples
func Marshal(c ...*graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}
	for _, v := range c {
		t := &KeyValue{
			Key: []byte(label + delimiter + v.Label() + delimiter + v.ID()),
			Value: &Any{
				TypeUrl: "Vertex",
				Value:   []byte(v.ID()),
			},
		}
		tt = append(tt, t)

		for k, p := range v.Properties() {
			t := &KeyValue{
				Key: []byte(properties + delimiter + k + delimiter + v.ID()),
				Value: &Any{
					TypeUrl: fmt.Sprintf("%T", p),
					Value:   []byte(fmt.Sprint(p)),
				},
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &KeyValue{
				Key: []byte(relationship + delimiter + e.RelationshipType() + delimiter + v.ID()),
				Value: &Any{
					TypeUrl: "Vertex",
					Value:   []byte(e.ID()),
				},
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {
				t := &KeyValue{
					Key: []byte(relationship + delimiter + properties + delimiter + k + delimiter + v.ID()),
					Value: &Any{
						TypeUrl: fmt.Sprintf("%T", p),
						Value:   []byte(fmt.Sprint(p)),
					},
				}
				tt = append(tt, t)
			}
		}
	}
	return tt
}

// MarshalTranspose mashal a Vertex into a transposed triples
func MarshalTranspose(c ...*graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}
	for _, v := range c {
		t := &KeyValue{
			Key: []byte(v.ID()),
			Value: &Any{
				TypeUrl: label,
				Value:   []byte(v.Label()),
			},
		}
		tt = append(tt, t)

		for k, p := range v.Properties() {
			t := &KeyValue{
				Key: []byte(v.ID() + properties + delimiter + k),
				Value: &Any{
					TypeUrl: fmt.Sprintf("%T", p),
					Value:   []byte(fmt.Sprint(p)),
				},
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &KeyValue{
				Key: []byte(v.ID() + delimiter + "Relationship" + delimiter + e.RelationshipType()),
				Value: &Any{
					TypeUrl: "Vertex",
					Value:   []byte(e.ID()),
				},
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {
				t := &KeyValue{
					Key: []byte(v.ID() + delimiter + "Relationship" + delimiter + properties + delimiter + k + delimiter + e.ID()),
					Value: &Any{
						TypeUrl: fmt.Sprintf("%T", p),
						Value:   []byte(fmt.Sprint(p)),
					},
				}
				tt = append(tt, t)
			}
		}
	}
	return tt
}

// // MarshalTranspose mashal a Vertex into a transposed triples
// func MarshalTranspose(c ...*graph.Vertex) []*Triple {
// 	delimiter := string(Delimiter)
// 	tt := []*Triple{}
// 	for _, v := range c {
// 		t := &Triple{
// 			Row: &Row{Vertex + delimiter + v.Label() + delimiter + v.ID()},
// 			Body: &Body{
// 				Column: v.ID(),
// 				Value:  float64(1),
// 			},
// 		}
// 		tt = append(tt, t)

// 		for k, p := range v.Properties() {
// 			t := &Triple{
// 				Row: &Row{VertexProperties + delimiter + k + delimiter + v.ID()},
// 				Body: &Body{
// 					Column: fmt.Sprint(p),
// 					Value:  float64(1),
// 				},
// 			}
// 			tt = append(tt, t)
// 		}

// 		for _, e := range v.Edges() {
// 			t := &Triple{
// 				Row: &Row{Edge + delimiter + e.RelationshipType() + delimiter + e.ID()},
// 				Body: &Body{
// 					Column: v.ID(),
// 					Value:  e.Weight,
// 				},
// 			}
// 			tt = append(tt, t)

// 			for k, p := range e.Properties() {
// 				t := &Triple{
// 					Row: &Row{EdgeProperties + delimiter + k + delimiter + e.ID() + delimiter + v.ID()},
// 					Body: &Body{
// 						Column: fmt.Sprint(p),
// 						Value:  float64(1),
// 					},
// 				}
// 				tt = append(tt, t)
// 			}
// 		}
// 	}
// 	return tt
// }

// func Unmarshal(c ...*Triple) *graph.Vertex {
// 	return nil
// }
