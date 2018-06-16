package keyvalue

import (
	"fmt"
	"strings"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/arch"
)

const (
	vertex       = "v"
	label        = "l"
	properties   = "p"
	relationship = "r"

	// US unit separator can be used as delimiters to mark fields of data structures. If used for hierarchical levels, US is the lowest level (dividing plain-text data items)
	US = string('\u241F')

	stringEmpty = ""
)

// PropertiesID generate the properties key
func PropertiesID(ID, key string) []byte {
	return []byte(ID + US + properties + US + key)
}

// RelationshipID generate the relationship key
func RelationshipID(ID, relationshipType string) []byte {
	return []byte(ID + US + relationship + US + relationshipType)
}

// RelationshipPropertiesID generate the properties key for a relationship
func RelationshipPropertiesID(ID, edgeID, key string) []byte {
	return []byte(ID + US + relationship + US + properties + US + key + US + edgeID)
}

// Marshal a Vertex into KeyValue
func Marshal(c ...*graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}
	for _, v := range c {
		t := &KeyValue{
			Key:   []byte(v.ID()),
			Value: NewAny(v.Label()),
		}
		tt = append(tt, t)

		for k, p := range v.Properties() {
			t := &KeyValue{
				Key:   PropertiesID(v.ID(), k),
				Value: NewAny(p),
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &KeyValue{
				Key:   RelationshipID(v.ID(), e.RelationshipType()),
				Value: NewAny(e.ID()),
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {
				t := &KeyValue{
					Key:   RelationshipPropertiesID(v.ID(), e.ID(), k),
					Value: NewAny(p),
				}
				tt = append(tt, t)
			}
		}
	}
	return tt
}

// PropertiesIDTranspose generate the transpose properties key
func PropertiesIDTranspose(ID, key string) []byte {
	return []byte(properties + US + key + US + ID)
}

// RelationshipIDTranspose generate the transpose relationship key
func RelationshipIDTranspose(ID, relationshipType string) []byte {
	return []byte(relationship + US + relationshipType + US + ID)
}

// RelationshipPropertiesIDTranspose generate the transpose properties key for a relationship
func RelationshipPropertiesIDTranspose(ID, edgeID, key string) []byte {
	return []byte(relationship + US + properties + US + key + US + edgeID + US + ID)
}

// MarshalTranspose mashal a Vertex into a transposed KeyValue
func MarshalTranspose(c ...*graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}
	for _, v := range c {
		t := &KeyValue{
			Key:   []byte(label + US + v.Label() + US + v.ID()),
			Value: NewAny(v.ID()),
		}
		tt = append(tt, t)

		for k, p := range v.Properties() {
			t := &KeyValue{
				Key:   PropertiesIDTranspose(v.ID(), k),
				Value: NewAny(p),
			}
			tt = append(tt, t)
		}

		for _, e := range v.Edges() {
			t := &KeyValue{
				Key:   RelationshipIDTranspose(v.ID(), e.RelationshipType()),
				Value: NewAny(e.ID()),
			}
			tt = append(tt, t)

			for k, p := range e.Properties() {
				t := &KeyValue{
					Key:   RelationshipPropertiesIDTranspose(v.ID(), e.ID(), k),
					Value: NewAny(p),
				}
				tt = append(tt, t)
			}
		}
	}
	return tt
}

// Property generate the properties key
func Property(kv KeyValue) (string, interface{}, error) {
	split := strings.Split(string(kv.Key), US)

	if split[2] != properties {
		return stringEmpty, nil, fmt.Errorf("key is not a property")
	}

	property := split[4]

	return property, arch.DecodeType(kv.Value.TypeUrl, kv.Value.Value), nil
}

// // RelationshipID generate the relationship key
// func RelationshipID(ID, relationshipType string) []byte {
// 	return []byte(ID + US + relationship + US + relationshipType)
// }

// // RelationshipPropertiesID generate the properties key for a relationship
// func RelationshipPropertiesID(ID, edgeID, key string) []byte {
// 	return []byte(ID + US + relationship + US + properties + US + key + US + edgeID)
// }

// Unmarshal a KeyValue into Vertex
// func Unmarshal(c ...*KeyValue) graph.Vertex {
// 	tt := []*KeyValue{}
// 	for _, v := range c {
// 		t := &KeyValue{
// 			Key: []byte(v.ID()),
// 			Value: &Any{
// 				TypeUrl: label,
// 				Value:   []byte(v.Label()),
// 			},
// 		}
// 		tt = append(tt, t)

// 		for k, p := range v.Properties() {
// 			t := &KeyValue{
// 				Key: PropertiesID(v.ID(), k),
// 				Value: &Any{
// 					TypeUrl: fmt.Sprintf("%T", p),
// 					Value:   []byte(fmt.Sprint(p)),
// 				},
// 			}
// 			tt = append(tt, t)
// 		}

// 		for _, e := range v.Edges() {
// 			t := &KeyValue{
// 				Key: RelationshipID(v.ID(), e.RelationshipType()),
// 				Value: &Any{
// 					TypeUrl: vertex,
// 					Value:   []byte(e.ID()),
// 				},
// 			}
// 			tt = append(tt, t)

// 			for k, p := range e.Properties() {
// 				t := &KeyValue{
// 					Key: RelationshipPropertiesID(v.ID(), e.ID(), k),
// 					Value: &Any{
// 						TypeUrl: fmt.Sprintf("%T", p),
// 						Value:   []byte(fmt.Sprint(p)),
// 					},
// 				}
// 				tt = append(tt, t)
// 			}
// 		}
// 	}
// 	return tt
// }
