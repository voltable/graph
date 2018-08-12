package keyvalue

import (
	"bytes"

	"github.com/RossMerr/Caudex.Graph/arch"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

// NewKeyValueVertex creates a vertex KeyValue
func NewKeyValueVertex(id *uuid.UUID, label string) *KeyValue {
	return &KeyValue{
		Value: NewAny(label),
		Key:   NewKey(id[:], &Column{Label, nil}).Marshal(),
	}
}

// NewKeyValueProperty creates a property KeyValue
func NewKeyValueProperty(id *uuid.UUID, key string, value interface{}) *KeyValue {
	return &KeyValue{
		Value: NewAny(value),
		Key:   NewKey(id[:], &Column{Properties, []byte(key)}).Marshal(),
	}
}

// NewKeyValueRelationship creates a relationship KeyValue
func NewKeyValueRelationship(from, to *uuid.UUID, relationshipType string, weight float64) *KeyValue {
	return &KeyValue{
		Value: NewAny(weight),
		Key:   NewKey(from[:], &Column{append(Relationship, append(US, []byte(relationshipType)...)...), to[:]}).Marshal(),
	}
}

// NewKeyValueRelationshipProperty creates a relationship property KeyValue
func NewKeyValueRelationshipProperty(from, to *uuid.UUID, key string, value interface{}) *KeyValue {
	return &KeyValue{
		Value: NewAny(value),
		Key:   NewKey(from[:], &Column{append(Relationshipproperties, append(US, []byte(key)...)...), to[:]}).Marshal(),
	}
}

// Transpose

// NewKeyValueVertexTranspose creates a vertex KeyValue
func NewKeyValueVertexTranspose(id *uuid.UUID, label string) *KeyValue {
	return &KeyValue{
		Value: NewAny(id[:]),
		Key:   NewKey(TLabel, &Column{[]byte(label), id[:]}).Marshal(),
	}
}

// NewKeyValuePropertyTranspose creates a property KeyValue
func NewKeyValuePropertyTranspose(id *uuid.UUID, key string, value interface{}) *KeyValue {
	return &KeyValue{
		Value: NewAny(value),
		Key:   NewKey(TProperties, &Column{[]byte(key), id[:]}).Marshal(),
	}
}

// NewKeyValueRelationshipTranspose creates a relationship KeyValue
func NewKeyValueRelationshipTranspose(from, to *uuid.UUID, relationshipType string, weight float64) *KeyValue {
	return &KeyValue{
		Value: NewAny(to[:]),
		Key:   NewKey(append(TRelationship, append(US, []byte(relationshipType)...)...), &Column{arch.EncodeFloat64Bytes(weight), from[:]}).Marshal(),
	}
}

// NewKeyValueRelationshipPropertyTranspose creates a relationship property KeyValue
func NewKeyValueRelationshipPropertyTranspose(from, to *uuid.UUID, key string, value interface{}) *KeyValue {
	return &KeyValue{
		Value: NewAny(value),
		Key:   NewKey(append(TRelationshipproperties, append(US, []byte(key)...)...), &Column{to[:], from[:]}).Marshal(),
	}
}

// MarshalKeyValue marshal a Vertex into KeyValue
func MarshalKeyValue(v *graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}

	id := v.ID()
	tt = append(tt, NewKeyValueVertex(id, v.Label()))

	for k, p := range v.Properties() {
		tt = append(tt, NewKeyValueProperty(id, k, p))
	}

	for _, e := range v.Edges() {
		tt = append(tt, MarshalEdgeKeyValue(e)...)
	}

	return tt
}

// MarshalKeyValueTranspose mashal a Vertex into a transposed KeyValue
func MarshalKeyValueTranspose(v *graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}

	id := v.ID()
	tt = append(tt, NewKeyValueVertexTranspose(id, v.Label()))

	for k, p := range v.Properties() {
		tt = append(tt, NewKeyValuePropertyTranspose(id, k, p))
	}

	for _, e := range v.Edges() {
		tt = append(tt, MarshalEdgeKeyValueTranspose(e)...)
	}
	return tt
}

// UnmarshalKeyValue a KeyValue into Vertex
func UnmarshalKeyValue(v *graph.Vertex, c []*KeyValue) {
	for _, kv := range c {

		key := &Key{}
		key.Unmarshal(kv.Key)

		if bytes.Equal(key.Column.Family, Label) {
			id := uuid.SliceToUUID(key.ID)
			v.SetID(id)
			value, ok := Unmarshal(kv.Value).(string)
			if ok {
				v.SetLabel(value)
			}
			continue
		}

		if bytes.Equal(key.Column.Family, Properties) {
			v.SetProperty(string(key.Column.Qualifier), Unmarshal(kv.Value))
			continue
		}

		splitSub := bytes.Split(key.Column.Family, US)
		if len(splitSub) > 0 {
			if bytes.Equal(splitSub[0], Relationship) {
				edgeID := uuid.SliceToUUID(key.Column.Qualifier)
				edge, ok := v.Edges()[*edgeID]
				if !ok {
					edge = graph.NewEdgeFromID(v.ID(), edgeID)
					v.AddEdge(edge)
				}
				edge.SetRelationshipType(string(splitSub[1]))

				value, ok := Unmarshal(kv.Value).(float64)
				if ok {
					edge.Weight = value
				}

				continue
			}

			if bytes.Equal(splitSub[0], Relationshipproperties) {
				edgeID := uuid.SliceToUUID(key.Column.Qualifier)
				edge, ok := v.Edges()[*edgeID]
				if !ok {
					edge = graph.NewEdgeFromID(v.ID(), edgeID)
					v.AddEdge(edge)
				}

				edge.SetProperty(string(splitSub[1]), Unmarshal(kv.Value))
				continue
			}
		}
	}
}

// UnmarshalKeyValueTranspose a KeyValue into Vertex
func UnmarshalKeyValueTranspose(v *graph.Vertex, c []*KeyValue) {
	for _, kv := range c {

		key := &Key{}
		key.Unmarshal(kv.Key)

		if bytes.Equal(key.ID, TLabel) {
			if s, ok := Unmarshal(kv.Value).([]byte); ok {
				id := uuid.SliceToUUID(s)
				v.SetID(id)
			}
			v.SetLabel(string(key.Column.Family))
			continue
		}
		if bytes.Equal(key.ID, TProperties) {
			v.SetProperty(string(key.Column.Family), Unmarshal(kv.Value))
			continue
		}
		splitSub := bytes.Split(key.ID, US)

		if len(splitSub) > 0 {
			if bytes.Equal(splitSub[0], TRelationship) {
				relationshipType := string(splitSub[1])

				value, ok := Unmarshal(kv.Value).([]byte)
				if ok {
					edgeID := uuid.SliceToUUID(value)

					edge, ok := v.Edges()[*edgeID]
					if !ok {
						edge = graph.NewEdgeFromID(v.ID(), edgeID)
						v.AddEdge(edge)
					}

					edge.SetRelationshipType(relationshipType)
					if weight, ok := arch.DecodeFloat64Bytes(key.Column.Family).(float64); ok {
						edge.Weight = weight
					}
				}
				continue
			}

			if bytes.Equal(splitSub[0], TRelationshipproperties) {
				edgeID := uuid.SliceToUUID(key.Column.Family)
				edge, ok := v.Edges()[*edgeID]
				if !ok {
					edge = graph.NewEdgeFromID(v.ID(), edgeID)
					v.AddEdge(edge)
				}
				edge.SetProperty(string(splitSub[1]), Unmarshal(kv.Value))
				continue
			}
		}
	}
}

// MarshalEdgeKeyValue marshal a edge into KeyValue
func MarshalEdgeKeyValue(e *graph.Edge) []*KeyValue {
	tt := []*KeyValue{}

	from := e.From()
	to := e.To()
	tt = append(tt, NewKeyValueRelationship(from, to, e.RelationshipType(), e.Weight))

	for k, p := range e.Properties() {
		tt = append(tt, NewKeyValueRelationshipProperty(from, to, k, p))
	}

	return tt
}

// MarshalEdgeKeyValueTranspose mashal a Edge into a transposed KeyValue
func MarshalEdgeKeyValueTranspose(e *graph.Edge) []*KeyValue {
	tt := []*KeyValue{}

	from := e.From()
	to := e.To()
	tt = append(tt, NewKeyValueRelationshipTranspose(from, to, e.RelationshipType(), e.Weight))

	for k, p := range e.Properties() {
		tt = append(tt, NewKeyValueRelationshipPropertyTranspose(from, to, k, p))
	}

	return tt
}
