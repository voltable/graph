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
func NewVertexFromID(ID [16]byte) (*Vertex, error) {
	v, err := NewVertexWithLabel("")
	v.id = ID
	return v, err
}

// NewVertexWithLabel create a vertex with the set label
func NewVertexWithLabel(label string) (*Vertex, error) {
	var id VertexID
	var err error

	if id, err = generateRandomBytes(); err != nil {
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

	t := &keyvalue.KeyValue{
		Key:   []byte(v.ID() + US + vertex),
		Value: keyvalue.NewAny(v.Label()),
	}
	tt = append(tt, t)

	for k, p := range v.Properties() {
		t := &keyvalue.KeyValue{
			Key:   []byte(v.ID() + US + properties + US + k),
			Value: keyvalue.NewAny(p),
		}
		tt = append(tt, t)
	}

	for _, e := range v.Edges {
		tt = append(tt, e.MarshalKeyValue()...)
	}

	return tt
}

// MarshalKeyValueTranspose mashal a Vertex into a transposed KeyValue
func (v *Vertex) MarshalKeyValueTranspose() []*keyvalue.KeyValue {
	tt := []*keyvalue.KeyValue{}

	t := &keyvalue.KeyValue{
		Key:   []byte(label + US + v.Label() + US + v.ID()),
		Value: keyvalue.NewAny(v.ID()),
	}
	tt = append(tt, t)

	for k, p := range v.Properties() {
		t := &keyvalue.KeyValue{
			Key:   []byte(properties + US + k + US + v.ID()),
			Value: keyvalue.NewAny(p),
		}
		tt = append(tt, t)
	}

	for _, e := range v.Edges {
		tt = append(tt, e.MarshalTranspose()...)
	}
	return tt
}

// UnmarshalKeyValue a KeyValue into Vertex
func (v *Vertex) UnmarshalKeyValue(c ...*keyvalue.KeyValue) {
	parts := strings.Split(string(c[0].Key), US)
	id := parts[0]
	uuid, _ := ParseUUID(id)
	v.id = uuid

	for _, kv := range c {
		split := strings.Split(string(kv.Key), US)

		if split[1] == vertex {
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				v.SetLabel(value)
			}
			continue
		}
		if split[1] == properties {
			key := split[2]
			v.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
		if split[1] == relationship {
			relationshipType := split[2]
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				edgeID, _ := ParseUUID(value)

				edge, ok := v.Edges[edgeID]
				if !ok {
					edge = NewEdgeFromID(v.id, edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(relationshipType)
			}
			continue
		}

		if split[1] == relationshipproperties {
			key := split[2]
			edgeID, _ := ParseUUID(split[3])
			edge, ok := v.Edges[edgeID]
			if !ok {
				edge := NewEdgeFromID(v.id, edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
	}
}

// UnmarshalKeyValueTranspose a KeyValue into Vertex
func (v *Vertex) UnmarshalKeyValueTranspose(c ...*keyvalue.KeyValue) {

	if s, ok := c[0].Value.Unmarshal().(string); ok {
		uuid, _ := ParseUUID(s)
		v.id = uuid
	}

	for _, kv := range c {
		split := strings.Split(string(kv.Key), US)

		if split[0] == label {
			label := split[1]
			v.SetLabel(label)
			continue
		}
		if split[0] == properties {
			key := split[1]
			v.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
		if split[0] == relationship {
			relationshipType := split[1]
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				edgeID, _ := ParseUUID(value)

				edge, ok := v.Edges[edgeID]
				if !ok {
					edge = NewEdgeFromID(v.id, edgeID)
					v.AddEdge(edge)
				}

				edge.SetRelationshipType(relationshipType)
			}
			continue
		}

		if split[0] == relationshipproperties {
			key := split[1]
			edgeID, _ := ParseUUID(split[2])
			edge, ok := v.Edges[edgeID]
			if !ok {
				edge := NewEdgeFromID(v.id, edgeID)
				v.AddEdge(edge)
			}
			edge.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
	}
}
