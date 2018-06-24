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

		if isVertex(split) {
			//id, _ := Vertex(split)
			value, ok := kv.Value.Unmarshal().(string)
			if ok {
				v.SetLabel(value)
			}
			continue
		}
		if isProperty(split) {
			_, key := property(split)
			v.SetProperty(key, kv.Value.Unmarshal())
			continue
		}
		if isRelationship(split) {
			_, relationshipType, _ := relationshipKey(split)
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

		if isRelationshipProperties(split) {
			_, value, key, _ := relationshipProperties(split)
			edgeID, _ := ParseUUID(value)
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

func isVertex(split []string) bool {
	if split[1] == vertex {
		return true
	}

	return false
}

func isProperty(split []string) bool {
	if split[1] == properties {
		return true
	}

	return false
}

// Property generate the properties key
func property(split []string) (string, string) {
	id := split[0]

	property := split[2]

	return id, property
}

func isRelationship(split []string) bool {
	if split[1] == relationship {
		return true
	}

	return false
}

//Relationship generate the relationship key
func relationshipKey(split []string) (string, string, error) {
	id := split[0]

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
func relationshipProperties(split []string) (string, string, string, error) {
	id := split[0]
	key := split[2]
	edgeID := split[3]

	return id, edgeID, key, nil
}

// UnmarshalKeyValueTranspose a KeyValue into Vertex
func (v *Vertex) UnmarshalKeyValueTranspose(c ...*keyvalue.KeyValue) {

	if s, ok := c[0].Value.Unmarshal().(string); ok {
		uuid, _ := ParseUUID(s)
		v.id = uuid
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

		if isRelationshipPropertiesTranspose(split) {
			_, value, key := relationshipPropertiesTranspose(split)
			edgeID, _ := ParseUUID(value)
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
