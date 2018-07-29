package keyvalue

import (
	"bytes"

	"github.com/RossMerr/Caudex.Graph/arch"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

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
		split := bytes.Split(kv.Key, US)
		if bytes.Equal(split[1], Vertex) {
			id := uuid.SliceToUUID(split[0])
			v.SetID(id)
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				v.SetLabel(value)
			}
			continue
		}
		if bytes.Equal(split[1], Properties) {
			key := split[2]
			v.SetProperty(string(key), kv.Value.Unmarshal())
			continue
		}
		if bytes.Equal(split[1], Relationship) {
			relationshipType := split[2]

			to := split[3]
			edgeID := uuid.SliceToUUID(to)
			edge, ok := v.Edges()[*edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}
			edge.SetRelationshipType(string(relationshipType))

			value, ok := kv.Value.Unmarshal().(float64)
			if ok {
				edge.Weight = value
			}

			continue
		}

		if bytes.Equal(split[1], Relationshipproperties) {
			key := string(split[2])
			edgeID := uuid.SliceToUUID(split[3])
			edge, ok := v.Edges()[*edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}

			edge.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
	}
}

// UnmarshalKeyValueTranspose a KeyValue into Vertex
func UnmarshalKeyValueTranspose(v *graph.Vertex, c []*KeyValue) {
	for _, kv := range c {
		split := bytes.Split(kv.Key, US)

		if bytes.Equal(split[0], Vertex) {
			if s, ok := kv.Value.Unmarshal().([]byte); ok {
				id := uuid.SliceToUUID(s)
				v.SetID(id)
			}
			v.SetLabel(string(split[1]))
			continue
		}
		if bytes.Equal(split[0], Properties) {
			v.SetProperty(string(split[1]), kv.Value.Unmarshal())
			continue
		}
		if bytes.Equal(split[0], Relationship) {
			relationshipType := split[1]

			value, ok := kv.Value.Unmarshal().([]byte)
			if ok {
				edgeID := uuid.SliceToUUID(value)

				edge, ok := v.Edges()[*edgeID]
				if !ok {
					edge = graph.NewEdgeFromID(v.ID(), edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(string(relationshipType))
				if weight, ok := arch.DecodeFloat64Bytes(split[2]).(float64); ok {
					edge.Weight = weight
				}
			}
			continue
		}

		if bytes.Equal(split[0], Relationshipproperties) {
			key := split[1]
			edgeID := uuid.SliceToUUID(split[2])
			edge, ok := v.Edges()[*edgeID]
			if !ok {
				edge = graph.NewEdgeFromID(v.ID(), edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(string(key), kv.Value.Unmarshal())
			continue
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
