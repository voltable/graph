package keyvalue

import (
	"bytes"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

// MarshalKeyValue marshal a Vertex into KeyValue
func MarshalKeyValue(v *graph.Vertex) []*KeyValue {
	tt := []*KeyValue{}

	id := v.ID()
	tt = append(tt, NewKeyValue(v.Label(), id[:], US, Vertex))

	for k, p := range v.Properties() {
		tt = append(tt, NewKeyValue(p, id[:], US, Properties, US, []byte(k)))
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
	tt = append(tt, NewKeyValue(id[:], Label, US, []byte(v.Label()), US, id[:]))

	for k, p := range v.Properties() {
		tt = append(tt, NewKeyValue(p, Properties, US, []byte(k), US, id[:]))
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
			value, ok := kv.Value.Unmarshal().([]byte)
			if ok {
				edgeID := uuid.SliceToUUID(value)

				edge, ok := v.Edges()[edgeID]
				if !ok {
					edge = graph.NewEdgeFromID(v.ID(), edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(string(relationshipType))
			}
			continue
		}

		if bytes.Equal(split[1], Relationshipproperties) {
			key := string(split[2])
			edgeID := uuid.SliceToUUID(split[3])
			edge, ok := v.Edges()[edgeID]
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

		if bytes.Equal(split[0], Label) {
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

				edge, ok := v.Edges()[edgeID]
				if !ok {
					edge = graph.NewEdgeFromID(v.ID(), edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(string(relationshipType))
			}
			continue
		}

		if bytes.Equal(split[0], Relationshipproperties) {
			key := split[1]
			edgeID := uuid.SliceToUUID(split[2])
			edge, ok := v.Edges()[edgeID]
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
	tt = append(tt, NewKeyValue(to[:], from[:], US, Relationship, US, []byte(e.RelationshipType())))

	for k, p := range e.Properties() {
		tt = append(tt, NewKeyValue(p, from[:], US, Relationshipproperties, US, []byte(k), US, to[:]))
	}

	return tt
}

// MarshalEdgeKeyValueTranspose mashal a Edge into a transposed KeyValue
func MarshalEdgeKeyValueTranspose(e *graph.Edge) []*KeyValue {
	tt := []*KeyValue{}

	from := e.From()
	to := e.To()
	tt = append(tt, NewKeyValue(to[:], Relationship, US, []byte(e.RelationshipType()), US, from[:]))

	for k, p := range e.Properties() {
		tt = append(tt, NewKeyValue(p, Relationshipproperties, US, []byte(k), US, to[:], US, from[:]))
	}

	return tt
}
