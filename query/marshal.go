package query

import (
	"bytes"

	"github.com/voltable/graph/widecolumnstore"

	graph "github.com/voltable/graph"
	"github.com/voltable/graph/uuid"
)

// NewKeyValueID creates a ID KeyValue
func NewKeyValueID(id uuid.UUID, label string) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
	return &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(TID),
			Key:   widecolumnstore.NewKey(id[:], &widecolumnstore.Column{nil, nil, nil}).Marshal(),
		}, &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(id[:]),
			Key:   widecolumnstore.NewKey(TID, &widecolumnstore.Column{id[:], nil, nil}).Marshal(),
		}
}

// NewKeyValueVertex creates a vertex KeyValue
func NewKeyValueVertex(id uuid.UUID, label string) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
	return &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(label),
			Key:   widecolumnstore.NewKey(id[:], &widecolumnstore.Column{Label, nil, nil}).Marshal(),
		}, &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(id[:]),
			Key:   widecolumnstore.NewKey(TLabel, &widecolumnstore.Column{[]byte(label), nil, id[:]}).Marshal(),
		}
}

// NewKeyValueProperty creates a property KeyValue
func NewKeyValueProperty(id uuid.UUID, key string, value interface{}) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
	return &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(value),
			Key:   widecolumnstore.NewKey(id[:], &widecolumnstore.Column{Properties, nil, []byte(key)}).Marshal(),
		}, &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(value),
			Key:   widecolumnstore.NewKey(TProperties, &widecolumnstore.Column{[]byte(key), nil, id[:]}).Marshal(),
		}
}

// NewKeyValueRelationship creates a relationship KeyValue
func NewKeyValueRelationship(from, to uuid.UUID, relationshipType string, weight float64) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
	return &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(weight),
			Key:   widecolumnstore.NewKey(from[:], &widecolumnstore.Column{Relationship, []byte(relationshipType), to[:]}).Marshal(),
		}, &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(weight),
			Key:   widecolumnstore.NewKey(to[:], &widecolumnstore.Column{TRelationship, []byte(relationshipType), from[:]}).Marshal(),
		}
}

// NewKeyValueRelationshipProperty creates a relationship property KeyValue
func NewKeyValueRelationshipProperty(from, to uuid.UUID, key string, value interface{}) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
	return &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(value),
			Key:   widecolumnstore.NewKey(from[:], &widecolumnstore.Column{Relationshipproperties, []byte(key), to[:]}).Marshal(),
		}, &widecolumnstore.KeyValue{
			Value: widecolumnstore.NewAny(value),
			Key:   widecolumnstore.NewKey(TRelationshipproperties, &widecolumnstore.Column{[]byte(key), to[:], from[:]}).Marshal(),
		}
}

// MarshalKeyValue marshal a Vertex into KeyValue
func MarshalKeyValue(v *graph.Vertex) ([]*widecolumnstore.KeyValue, []*widecolumnstore.KeyValue) {
	keyvalues := []*widecolumnstore.KeyValue{}
	transposed := []*widecolumnstore.KeyValue{}
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
func UnmarshalKeyValue(v *graph.Vertex, c []*widecolumnstore.KeyValue) {
	for _, kv := range c {

		key := &widecolumnstore.Key{}
		key.Unmarshal(kv.Key)

		if bytes.Equal(key.Column.Family, Label) {
			id := uuid.SliceToUUID(key.ID)
			v.SetID(id)
			value, ok := widecolumnstore.Unmarshal(kv.Value).(string)
			if ok {
				v.SetLabel(value)
			}
			continue
		}

		if bytes.Equal(key.Column.Family, Properties) {
			v.SetProperty(string(key.Column.Qualifier), widecolumnstore.Unmarshal(kv.Value))
			continue
		}

		if bytes.Equal(key.Column.Family, Relationship) {
			edgeID := uuid.SliceToUUID(key.Column.Qualifier)
			edge, ok := v.Edges()[edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}
			edge.SetRelationshipType(string(key.Column.Extended))

			value, ok := widecolumnstore.Unmarshal(kv.Value).(float64)
			if ok {
				edge.Weight = value
			}

			continue
		}

		if bytes.Equal(key.Column.Family, Relationshipproperties) {
			edgeID := uuid.SliceToUUID(key.Column.Qualifier)
			edge, ok := v.Edges()[edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}

			edge.SetProperty(string(key.Column.Extended), widecolumnstore.Unmarshal(kv.Value))
			continue
		}

	}
}

// UnmarshalKeyValueTranspose a KeyValue into Vertex
func UnmarshalKeyValueTranspose(v *graph.Vertex, c []*widecolumnstore.KeyValue) {
	for _, kv := range c {

		key := &widecolumnstore.Key{}
		key.Unmarshal(kv.Key)

		if bytes.Equal(key.ID, TLabel) {
			if s, ok := widecolumnstore.Unmarshal(kv.Value).([]byte); ok {
				id := uuid.SliceToUUID(s)
				v.SetID(id)
			}
			v.SetLabel(string(key.Column.Family))
			continue
		}
		if bytes.Equal(key.ID, TProperties) {
			v.SetProperty(string(key.Column.Family), widecolumnstore.Unmarshal(kv.Value))
			continue
		}

		if bytes.Equal(key.Column.Family, TRelationship) {
			relationshipType := string(key.Column.Extended)

			edgeID := uuid.SliceToUUID(key.ID)

			edge, ok := v.Edges()[edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}

			edge.SetRelationshipType(relationshipType)
			weight, ok := widecolumnstore.Unmarshal(kv.Value).(float64)
			if ok {
				edge.Weight = weight
			}

			continue
		}

		if bytes.Equal(key.ID, TRelationshipproperties) {
			edgeID := uuid.SliceToUUID(key.Column.Extended)
			edge, ok := v.Edges()[edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(string(key.Column.Family), widecolumnstore.Unmarshal(kv.Value))
			continue
		}

	}
}


// MarshalEdgeKeyValue marshal a edge into KeyValue
func MarshalEdgeKeyValue(e *graph.Edge) ([]*widecolumnstore.KeyValue, []*widecolumnstore.KeyValue) {
	keyvalues := []*widecolumnstore.KeyValue{}
	transposed := []*widecolumnstore.KeyValue{}
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
