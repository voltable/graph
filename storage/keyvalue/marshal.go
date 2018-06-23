package keyvalue

import (
	"fmt"
	"strings"

	graph "github.com/RossMerr/Caudex.Graph"
)

const (
	vertex                 = "v"
	label                  = "l"
	properties             = "p"
	relationship           = "r"
	relationshipproperties = "k"
	// US unit separator can be used as delimiters to mark fields of data structures. If used for hierarchical levels, US is the lowest level (dividing plain-text data items)
	US = string('\u241F')

	stringEmpty = ""
)

// VertexID generate the vertex key
func VertexID(ID string) []byte {
	return []byte(ID + US + vertex)
}

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
	return []byte(ID + US + relationshipproperties + US + key + US + edgeID)
}

// Marshal a Vertex into KeyValue
func Marshal(v *graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}

	t := &KeyValue{
		Key:   VertexID(v.ID()),
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

	for _, e := range v.Edges {
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

	return tt
}

func LabelID(ID, l string) []byte {
	return []byte(label + US + l + US + ID)
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
	return []byte(relationshipproperties + US + key + US + edgeID + US + ID)
}

// MarshalTranspose mashal a Vertex into a transposed KeyValue
func MarshalTranspose(v *graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}
	t := &KeyValue{
		Key:   LabelID(v.ID(), v.Label()),
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

	for _, e := range v.Edges {
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
	return tt
}

func isVertex(split []string) bool {
	if split[1] == vertex {
		return true
	}

	return false
}

func Vertex(split []string) (string, error) {
	id := split[0]

	if split[1] != vertex {
		return stringEmpty, fmt.Errorf("key is not a vertex")
	}
	return id, nil
}

func isProperty(split []string) bool {
	if split[1] == properties {
		return true
	}

	return false
}

// Property generate the properties key
func Property(split []string) (string, string, error) {
	id := split[0]

	if split[1] != properties {
		return stringEmpty, stringEmpty, fmt.Errorf("key is not a property")
	}

	property := split[2]

	return id, property, nil
}

func isRelationship(split []string) bool {
	if split[1] == relationship {
		return true
	}

	return false
}

//Relationship generate the relationship key
func Relationship(split []string) (string, string, error) {
	id := split[0]

	if split[1] != relationship {
		return stringEmpty, stringEmpty, fmt.Errorf("key is not a relationship")
	}

	relationshipType := split[2]

	return id, relationshipType, nil
}

func isRelationshipProperties(split []string) bool {
	if split[1] == relationshipproperties {
		return true
	}

	return false
}

// RelationshipProperties generate the properties key for a relationship
func RelationshipProperties(split []string) (string, string, string, error) {
	id := split[0]

	if split[1] != relationshipproperties {
		return stringEmpty, stringEmpty, stringEmpty, fmt.Errorf("key is not a relationshipproperties")
	}

	key := split[2]
	edgeID := split[3]

	return id, edgeID, key, nil
}

// Unmarshal a KeyValue into Vertex
func Unmarshal(c ...*KeyValue) *graph.Vertex {
	parts := strings.Split(string(c[0].Key), US)
	id := parts[0]
	uuid, _ := graph.ParseUUID(id)
	v, _ := graph.NewVertexFromID(uuid)

	for _, kv := range c {
		split := strings.Split(string(kv.Key), US)

		if isVertex(split) {
			//id, _ := Vertex(split)
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				v.SetLabel(value)
			}
			continue
		}
		if isProperty(split) {
			_, key, _ := Property(split)
			v.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
		if isRelationship(split) {
			_, relationshipType, _ := Relationship(split)
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				edgeID, _ := graph.ParseUUID(value)

				edge, ok := v.Edges[edgeID]
				if !ok {
					edge, _ = graph.NewEdgeFromID(edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(relationshipType)
			}
			continue
		}

		if isRelationshipProperties(split) {
			_, value, key, _ := RelationshipProperties(split)
			edgeID, _ := graph.ParseUUID(value)
			edge, ok := v.Edges[edgeID]
			if !ok {
				edge, _ := graph.NewEdgeFromID(edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
	}
	return v
}

func isVertexTranspose(split []string) bool {
	if split[0] == label {
		return true
	}

	return false
}

func vertexTranspose(split []string) (string, string) {
	label := split[1]
	id := split[2]

	return id, label
}

func isPropertiesTranspose(split []string) bool {
	if split[0] == properties {
		return true
	}

	return false
}

func propertiesTranspose(split []string) (string, string) {
	key := split[1]
	id := split[2]

	return id, key
}

func isRelationshipTranspose(split []string) bool {
	if split[0] == relationship {
		return true
	}

	return false
}

//relationshipTranspose generate the relationship key
func relationshipTranspose(split []string) (string, string) {
	relationshipType := split[1]
	id := split[2]

	return id, relationshipType
}

func isRelationshipPropertiesTranspose(split []string) bool {
	if split[0] == relationshipproperties {
		return true
	}

	return false
}

// relationshipProperties generate the properties key for a relationship
func relationshipPropertiesTranspose(split []string) (string, string, string) {
	key := split[1]
	edgeID := split[2]
	id := split[3]
	return id, edgeID, key
}

// UnmarshalTranspose a KeyValue into Vertex
func UnmarshalTranspose(c ...*KeyValue) *graph.Vertex {
	var v *graph.Vertex
	if s, ok := c[0].Value.Unmarshal().(string); ok {
		uuid, _ := graph.ParseUUID(s)
		v, _ = graph.NewVertexFromID(uuid)
	}

	for _, kv := range c {
		split := strings.Split(string(kv.Key), US)

		if isVertexTranspose(split) {
			_, label := vertexTranspose(split)
			v.SetLabel(label)
			continue
		}
		if isPropertiesTranspose(split) {
			_, key := propertiesTranspose(split)
			v.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
		if isRelationshipTranspose(split) {
			_, relationshipType := relationshipTranspose(split)
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				edgeID, _ := graph.ParseUUID(value)

				edge, ok := v.Edges[edgeID]
				if !ok {
					edge, _ = graph.NewEdgeFromID(edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(relationshipType)
			}
			continue
		}

		if isRelationshipPropertiesTranspose(split) {
			_, value, key := relationshipPropertiesTranspose(split)
			edgeID, _ := graph.ParseUUID(value)
			edge, ok := v.Edges[edgeID]
			if !ok {
				edge, _ := graph.NewEdgeFromID(edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
	}
	return v
}
