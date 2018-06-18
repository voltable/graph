package graph_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph"
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

	if v.Edges == nil {
		t.Fatalf("Expected edges to be set but was %+v", v.Edges)
	}

}

func Test_AddDirectedEdge(t *testing.T) {
	vertex, _ := graph.NewVertex()
	vertexDirection, _ := graph.NewVertex()
	_, err := vertex.AddDirectedEdge(vertexDirection)

	if err != nil {
		t.Fatalf("Unexpected AddDirectedEdge error %s", err)
	}

	results := vertex.Edges

	if len(results) != 1 {
		t.Fatalf("Expected 1 edge but was %+v", len(results))
	}

	results2 := vertexDirection.Edges

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
