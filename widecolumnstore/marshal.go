package widecolumnstore

import (
	"bytes"

	"github.com/RossMerr/Caudex.Graph/arch"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

// NewKeyValueVertex creates a vertex KeyValue
func NewKeyValueVertex(id *uuid.UUID, label string) (*KeyValue, *KeyValue) {
	return &KeyValue{
			Value: NewAny(label),
			Key:   NewKey(id[:], &Column{Label, nil, nil}).Marshal(),
		}, &KeyValue{
			Value: NewAny(id[:]),
			Key:   NewKey(TLabel, &Column{[]byte(label), nil, id[:]}).Marshal(),
		}
}

// NewKeyValueProperty creates a property KeyValue
func NewKeyValueProperty(id *uuid.UUID, key string, value interface{}) (*KeyValue, *KeyValue) {
	return &KeyValue{
			Value: NewAny(value),
			Key:   NewKey(id[:], &Column{Properties, nil, []byte(key)}).Marshal(),
		}, &KeyValue{
			Value: NewAny(value),
			Key:   NewKey(TProperties, &Column{[]byte(key), nil, id[:]}).Marshal(),
		}
}

// NewKeyValueRelationship creates a relationship KeyValue
func NewKeyValueRelationship(from, to *uuid.UUID, relationshipType string, weight float64) (*KeyValue, *KeyValue) {
	return &KeyValue{
			Value: NewAny(weight),
			Key:   NewKey(from[:], &Column{Relationship, []byte(relationshipType), to[:]}).Marshal(),
		}, &KeyValue{
			Value: NewAny(to[:]),
			Key:   NewKey(TRelationship, &Column{[]byte(relationshipType), arch.EncodeFloat64Bytes(weight), from[:]}).Marshal(),
		}
}

// NewKeyValueRelationshipProperty creates a relationship property KeyValue
func NewKeyValueRelationshipProperty(from, to *uuid.UUID, key string, value interface{}) (*KeyValue, *KeyValue) {
	return &KeyValue{
			Value: NewAny(value),
			Key:   NewKey(from[:], &Column{Relationshipproperties, []byte(key), to[:]}).Marshal(),
		}, &KeyValue{
			Value: NewAny(value),
			Key:   NewKey(TRelationshipproperties, &Column{[]byte(key), to[:], from[:]}).Marshal(),
		}
}

// MarshalKeyValue marshal a Vertex into KeyValue
func MarshalKeyValue(v *graph.Vertex) ([]*KeyValue, []*KeyValue) {
	keyvalues := []*KeyValue{}
	transposed := []*KeyValue{}
	id := v.ID()

	k, t := NewKeyValueVertex(id, v.Label())
	keyvalues = append(keyvalues, k)
	transposed = append(transposed, t)

	for k, p := range v.Properties() {
		k, t := NewKeyValueProperty(id, k, p)
		keyvalues = append(keyvalues, k)
		transposed = append(transposed, t)
	}

	for _, e := range v.Edges() {
		k, t := MarshalEdgeKeyValue(e)
		keyvalues = append(keyvalues, k...)
		transposed = append(transposed, t...)
	}

	return keyvalues, transposed
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

		if bytes.Equal(key.Column.Family, Relationship) {
			edgeID := uuid.SliceToUUID(key.Column.Qualifier)
			edge, ok := v.Edges()[*edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}
			edge.SetRelationshipType(string(key.Column.Extended))

			value, ok := Unmarshal(kv.Value).(float64)
			if ok {
				edge.Weight = value
			}

			continue
		}

		if bytes.Equal(key.Column.Family, Relationshipproperties) {
			edgeID := uuid.SliceToUUID(key.Column.Qualifier)
			edge, ok := v.Edges()[*edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}

			edge.SetProperty(string(key.Column.Extended), Unmarshal(kv.Value))
			continue
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

		if bytes.Equal(key.ID, TRelationship) {
			relationshipType := string(key.Column.Family)

			value, ok := Unmarshal(kv.Value).([]byte)
			if ok {
				edgeID := uuid.SliceToUUID(value)

				edge, ok := v.Edges()[*edgeID]
				if !ok {
					edge = graph.NewEdgeFromID(v.ID(), edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(relationshipType)
				if weight, ok := arch.DecodeFloat64Bytes(key.Column.Extended).(float64); ok {
					edge.Weight = weight
				}
			}
			continue
		}

		if bytes.Equal(key.ID, TRelationshipproperties) {
			edgeID := uuid.SliceToUUID(key.Column.Extended)
			edge, ok := v.Edges()[*edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(string(key.Column.Family), Unmarshal(kv.Value))
			continue
		}

	}
}

// MarshalEdgeKeyValue marshal a edge into KeyValue
func MarshalEdgeKeyValue(e *graph.Edge) ([]*KeyValue, []*KeyValue) {
	keyvalues := []*KeyValue{}
	transposed := []*KeyValue{}
	from := e.From()
	to := e.To()

	k, t := NewKeyValueRelationship(from, to, e.RelationshipType(), e.Weight)
	keyvalues = append(keyvalues, k)
	transposed = append(transposed, t)

	for key, value := range e.Properties() {
		k, t := NewKeyValueRelationshipProperty(from, to, key, value)
		keyvalues = append(keyvalues, k)
		transposed = append(transposed, t)
	}

	return keyvalues, transposed
}
