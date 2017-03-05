package vertices

import "testing"

func Test_VertexLabels(t *testing.T) {
	v := Vertex{}
	v.SetLabel("foo")

	if v.Label() != "foo" {
		t.Fatalf("Expected label to be %s but was %s", "foo", v.Label())
	}
}

func Test_NewVertex(t *testing.T) {
	v, err := newVertex()

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
	v, _ := newVertex()
	e := Edge{}
	v.edges["test"] = e

	if len(v.Edges()) != 1 {
		t.Fatalf("Expected 1 edge but was %s", len(v.Edges()))
	}
}
