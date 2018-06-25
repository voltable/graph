package graph

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/RossMerr/Caudex.Graph/storage/keyvalue"
)

// VertexID a UUID
type VertexID [16]byte

var _ keyvalue.MarshalKeyValue = (*Vertex)(nil)

// Vertex .
type Vertex struct {
	id         VertexID
	Edges      map[VertexID]*Edge
	label      string
	properties map[string]interface{}
}

var (
	errCreatVertexID = errors.New("Failed to create Vertex ID")
	errEdgeNotFound  = errors.New("Edge Not found")
	errIdNotSet      = errors.New("Use NewVertex to create a new Vertex")
)

var _ Properties = (*Vertex)(nil)

// NewVertex creates the default vertex
func NewVertex() (*Vertex, error) {
	return NewVertexWithLabel("")
}

// NewVertexFromID creates a vertex using the id
func NewVertexFromID(ID VertexID) (*Vertex, error) {
	v, err := NewVertexWithLabel("")
	v.id = ID
	return v, err
}

// NewVertexWithLabel create a vertex with the set label
func NewVertexWithLabel(label string) (*Vertex, error) {
	var id VertexID
	var err error

	if id, err = generateRandomVertexID(); err != nil {
		return nil, errCreatVertexID
	}

	v := NewEmptyVertex()
	v.id = id
	v.label = label
	return v, nil
}

// NewEmptyVertex create's a empty vertex with no ID
func NewEmptyVertex() *Vertex {
	v := Vertex{Edges: make(map[VertexID]*Edge), properties: make(map[string]interface{})}
	return &v
}

// SetProperty set a property to store against this Vertex
func (v *Vertex) SetProperty(name string, property interface{}) {
	v.properties[name] = property
}

// Property gets a property to store on the Vertex
func (v *Vertex) Property(name string) interface{} {
	if value, ok := v.properties[name]; ok {
		return value
	}

	return nil
}

func (v *Vertex) DeleteProperty(name string) {
	delete(v.properties, name)
}

func (v *Vertex) PropertiesCount() int {
	return len(v.properties)
}

func (v *Vertex) Properties() map[string]interface{} {
	return v.properties
}

// ID returns the generate UUID
func (v *Vertex) ID() string {
	return formatUUID(v.id)
}

// Label vertex label type
func (v *Vertex) Label() string {
	return v.label
}

// Edges a array of all edges against this vertex
// func (v *Vertex) Edges() Edges {
// 	edges := make(Edges, 0, len(v.edges))
// 	for _, value := range v.edges {
// 		edges = append(edges, value)
// 	}

// 	return edges
// }

func (v *Vertex) removeRelationshipOnLabel(label string) Digraph {
	return v.removeRelationshipsF(func(id VertexID, e Edge) bool {
		return e.relationshipType == label
	})
}

func (v *Vertex) removeRelationships() {
	v.removeRelationshipsF(func(id VertexID, e Edge) bool {
		return true
	})
}

func (v *Vertex) removeRelationshipsOnVertex(to *Vertex) Digraph {
	return v.removeRelationshipsF(func(id VertexID, e Edge) bool {
		return id == to.id
	})
}

func (v *Vertex) removeRelationshipsF(f func(id VertexID, e Edge) bool) Digraph {
	for id, edge := range v.Edges {
		if f(id, *edge) {
			delete(v.Edges, edge.to)
			return edge.isDirected
		}
	}
	return Undirected
}

func (v *Vertex) SetLabel(label string) *Vertex {
	v.label = strings.ToLower(label)
	return v
}

// AddDirectedEdge links two vertex's and returns the edge
func (v *Vertex) AddDirectedEdge(to *Vertex) (*Edge, error) {
	edge := &Edge{from: v.id, to: to.id, isDirected: Directed, properties: make(map[string]interface{})}
	v.Edges[edge.to] = edge
	return edge, nil
}

// AddEdge links two vertex's and returns the edge
func (v *Vertex) AddEdge(to *Edge) {

	v.Edges[to.to] = to
}

// RemoveEdgeByLabel remove a edge
func (v *Vertex) RemoveEdgeByLabel(to *Vertex, label string) error {
	if to == nil {
		return errEdgeNotFound
	}

	isDirected := v.removeRelationshipOnLabel(label)

	if isDirected == Undirected {
		to.removeRelationshipOnLabel(label)
	}

	return nil
}

// RemoveEdge remove a edge
func (v *Vertex) RemoveEdge(to *Vertex) error {
	if to == nil {
		return errEdgeNotFound
	}

	isDirected := v.removeRelationshipsOnVertex(to)

	if isDirected == Undirected {
		to.removeRelationshipsOnVertex(v)
	}

	return nil
}

func (v Vertex) String() string {

	var buffer bytes.Buffer
	buffer.WriteString("{")
	for k, b := range v.properties {
		buffer.WriteString(fmt.Sprintf("%v => %#v", k, b))
		buffer.WriteString(", ")
	}
	w := bytes.NewBuffer(buffer.Bytes()[:buffer.Len()-2])
	w.WriteString("}")
	return fmt.Sprintf(w.String())
}

// MarshalKeyValue marshal a Vertex into KeyValue
func (v *Vertex) MarshalKeyValue() []*keyvalue.KeyValue {
	tt := []*keyvalue.KeyValue{}

	tt = append(tt, keyvalue.NewKeyValue(v.Label(), v.id[:], US, vertex))

	for k, p := range v.Properties() {
		tt = append(tt, keyvalue.NewKeyValue(p, v.id[:], US, properties, US, []byte(k)))
	}

	for _, e := range v.Edges {
		tt = append(tt, e.MarshalKeyValue()...)
	}

	return tt
}

// MarshalKeyValueTranspose mashal a Vertex into a transposed KeyValue
func (v *Vertex) MarshalKeyValueTranspose() []*keyvalue.KeyValue {
	tt := []*keyvalue.KeyValue{}

	tt = append(tt, keyvalue.NewKeyValue(v.ID(), label, US, []byte(v.Label()), US, v.id[:]))

	for k, p := range v.Properties() {
		tt = append(tt, keyvalue.NewKeyValue(p, properties, US, []byte(k), US, v.id[:]))
	}

	for _, e := range v.Edges {
		tt = append(tt, e.MarshalKeyValueTranspose()...)
	}
	return tt
}

// UnmarshalKeyValue a KeyValue into Vertex
func (v *Vertex) UnmarshalKeyValue(c []*keyvalue.KeyValue) {
	parts := bytes.Split(c[0].Key, US)
	uuid := sliceToVertexID(parts[0])
	v.id = uuid

	for _, kv := range c {
		split := bytes.Split(kv.Key, US)
		if bytes.Equal(split[1], vertex) {
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				v.SetLabel(value)
			}
			continue
		}
		if bytes.Equal(split[1], properties) {
			key := split[2]
			v.SetProperty(string(key), kv.Value.Unmarshal())
			continue
		}
		if bytes.Equal(split[1], relationship) {
			relationshipType := split[2]
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				edgeID, _ := parseUUID(value)

				edge, ok := v.Edges[edgeID]
				if !ok {
					edge = NewEdgeFromID(v.id, edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(string(relationshipType))
			}
			continue
		}

		if bytes.Equal(split[1], relationshipproperties) {
			key := split[2]
			edgeID := sliceToVertexID(split[3])
			edge, ok := v.Edges[edgeID]
			if !ok {
				edge := NewEdgeFromID(v.id, edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(string(key), kv.Value.Unmarshal())
			continue
		}
	}
}

// UnmarshalKeyValueTranspose a KeyValue into Vertex
func (v *Vertex) UnmarshalKeyValueTranspose(c []*keyvalue.KeyValue) {

	if s, ok := c[0].Value.Unmarshal().(string); ok {
		uuid, _ := parseUUID(s)
		v.id = uuid
	}

	for _, kv := range c {
		split := bytes.Split(kv.Key, US)

		if bytes.Equal(split[0], label) {
			v.SetLabel(string(split[1]))
			continue
		}
		if bytes.Equal(split[0], properties) {
			v.SetProperty(string(split[1]), kv.Value.Unmarshal())
			continue
		}
		if bytes.Equal(split[0], relationship) {
			relationshipType := split[1]
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				edgeID, _ := parseUUID(value)

				edge, ok := v.Edges[edgeID]
				if !ok {
					edge = NewEdgeFromID(v.id, edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(string(relationshipType))
			}
			continue
		}

		if bytes.Equal(split[0], relationshipproperties) {
			key := split[1]
			edgeID := sliceToVertexID(split[2])
			edge, ok := v.Edges[edgeID]
			if !ok {
				edge := NewEdgeFromID(v.id, edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(string(key), kv.Value.Unmarshal())
			continue
		}
	}
}
