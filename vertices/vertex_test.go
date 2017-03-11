package vertices

import (
	"testing"
)

func Test_VertexLabels(t *testing.T) {
	v := Vertex{}
	v.SetLabel("foo")

	if v.Label() != "foo" {
		t.Fatalf("Expected label to be %s but was %s", "foo", v.Label())
	}
}

func Test_NewVertex(t *testing.T) {
	v, err := NewVertex()

	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	if v.ID() == "" {
		t.Fatalf("Expected ID to be set but was %s", v.ID)
	}

	if v.edges == nil {
		t.Fatalf("Expected edges to be set but was %s", v.edges)
	}

}

func Test_Edges(t *testing.T) {
	v, _ := NewVertex()
	e := Edge{}
	v.edges["test"] = e

	if len(v.Edges()) != 1 {
		t.Fatalf("Expected 1 edge but was %s", len(v.Edges()))
	}
}

func Test_AddDirectedEdge(t *testing.T) {
	vertex, _ := NewVertex()
	vertexDirection, _ := NewVertex()
	_, err := vertex.AddDirectedEdge(vertexDirection)

	if err != nil {
		t.Fatalf("Unexpected AddDirectedEdge error %s", err)
	}

	results := vertex.Edges()

	if len(results) != 1 {
		t.Fatalf("Expected 1 edge but was %s", len(results))
	}

	results2 := vertexDirection.Edges()

	if len(results2) != 0 {
		t.Fatalf("Expected 0 edge but was %s", len(results2))
	}
}
