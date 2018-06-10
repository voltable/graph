package keyvalue

import (
	"fmt"

	graph "github.com/RossMerr/Caudex.Graph"
)

const (
	label        = "l"
	properties   = "p"
	relationship = "r"

	// US unit separator can be used as delimiters to mark fields of data structures. If used for hierarchical levels, US is the lowest level (dividing plain-text data items)
	US = string('\u241F')
)

// Marshal a Vertex into triples
func Marshal(c ...*graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}
	for _, v := range c {
		t := &KeyValue{
			Key: []byte(label + US + v.Label() + US + v.ID()),
			Value: &Any{
				TypeUrl: "Vertex",
				Value:   []byte(v.ID()),
			},
		}
		tt = append(tt, t)

		for k, p := range v.Properties() {
			t := &KeyValue{
				Key: []byte(properties + US + k + US + v.ID()),
				Value: &Any{
					TypeUrl: fmt.Sprintf("%T", p),
					Value:   []byte(fmt.Sprint(p)),
				},
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &KeyValue{
				Key: []byte(relationship + US + e.RelationshipType() + US + v.ID()),
				Value: &Any{
					TypeUrl: "Vertex",
					Value:   []byte(e.ID()),
				},
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {
				t := &KeyValue{
					Key: []byte(relationship + US + properties + US + k + US + v.ID()),
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
				Key: []byte(v.ID() + properties + US + k),
				Value: &Any{
					TypeUrl: fmt.Sprintf("%T", p),
					Value:   []byte(fmt.Sprint(p)),
				},
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &KeyValue{
				Key: []byte(v.ID() + US + "Relationship" + US + e.RelationshipType()),
				Value: &Any{
					TypeUrl: "Vertex",
					Value:   []byte(e.ID()),
				},
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {
				t := &KeyValue{
					Key: []byte(v.ID() + US + "Relationship" + US + properties + US + k + US + e.ID()),
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
