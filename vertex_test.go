package graph_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/keyvalue"
)

func Test_VertexLabels(t *testing.T) {
	v := graph.Vertex{}
	v.SetLabel("foo")

	if v.Label() != "foo" {
		t.Fatalf("Expected label to be %s but was %s", "foo", v.Label())
	}
}

func Test_NewVertex(t *testing.T) {
	v, err := graph.NewVertex()

	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	if v.ID() == "" {
		t.Fatalf("Expected ID to be set but was %+v", v.ID())
	}

	if v.Edges() == nil {
		t.Fatalf("Expected edges to be set but was %+v", v.Edges())
	}

}

func Test_AddDirectedEdge(t *testing.T) {
	vertex, _ := graph.NewVertex()
	vertexDirection, _ := graph.NewVertex()
	_, err := vertex.AddDirectedEdge(vertexDirection)

	if err != nil {
		t.Fatalf("Unexpected AddDirectedEdge error %s", err)
	}

	results := vertex.Edges()

	if len(results) != 1 {
		t.Fatalf("Expected 1 edge but was %+v", len(results))
	}

	results2 := vertexDirection.Edges()

	if len(results2) != 0 {
		t.Fatalf("Expected 0 edge but was %+v", len(results2))
	}
}

func Test_Value(t *testing.T) {
	x, _ := graph.NewVertex()
	x.SetProperty("Age", 10)
	if x.Property("Age") != 10 {
		t.Fatalf("Expected %d edge but was %v", 10, x.Property("Age"))
	}

}

func TestMarshalKeyValue(t *testing.T) {
	tests := []struct {
		name   string
		vertex *graph.Vertex
		want   []*keyvalue.KeyValue
	}{
		{
			name: "vertex",
			vertex: func() *graph.Vertex {
				v, _ := graph.NewVertex()
				v.SetLabel("person")
				v.SetProperty("name", "john smith")
				t, _ := graph.NewVertex()
				e := graph.NewEdge(v, t)
				e.SetRelationshipType("friend")
				e.SetProperty("years", 10)
				v.AddEdge(e)
				return v
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.vertex.MarshalKeyValue()
			v := graph.NewEmptyVertex()
			v.UnmarshalKeyValue(got)
			if !reflect.DeepEqual(v, tt.vertex) {
				t.Errorf("Marshal() = %v, want %v", v, tt.vertex)
			}
		})
	}
}

func TestMarshalKeyValueTranspose(t *testing.T) {
	tests := []struct {
		name   string
		vertex *graph.Vertex
		want   []*keyvalue.KeyValue
	}{
		{
			name: "vertex",
			vertex: func() *graph.Vertex {
				v, _ := graph.NewVertex()
				v.SetLabel("person")
				v.SetProperty("name", "john smith")
				t, _ := graph.NewVertex()
				e := graph.NewEdge(v, t)
				e.SetRelationshipType("friend")
				e.SetProperty("years", 10)
				v.AddEdge(e)
				return v
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.vertex.MarshalKeyValueTranspose()
			v := graph.NewEmptyVertex()
			v.UnmarshalKeyValueTranspose(got)
			if !reflect.DeepEqual(v, tt.vertex) {
				t.Errorf("Marshal() = %v, want %v", v, tt.vertex)
			}
		})
	}
}
