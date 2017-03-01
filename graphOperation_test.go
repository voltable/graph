package graphs

import (
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
)

var (
	se  *fakeStorageEngine
	obj *testObject
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.DebugLevel)
	se = &fakeStorageEngine{}
	obj = &testObject{value: "hello"}
	se.vertices = make(map[string]Vertex)
	os.Exit(m.Run())
}

type testObject struct {
	value string
}

func Test_CreateVertex(t *testing.T) {
	g := GraphOperation{se}
	v, err := g.CreateVertex(obj)
	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	if v.Value != obj {
		t.Fatalf("Expected %s but was %s", obj, v.Value)
	}
}

func Test_ReadVertex(t *testing.T) {
	g := GraphOperation{se}
	v, _ := g.CreateVertex(obj)
	v, err := g.ReadVertex(v.ID)
	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	if v.Value != obj {
		t.Fatalf("Expected %s but was %s", obj, v.Value)
	}
}

func Test_UpdateVertex(t *testing.T) {
	g := GraphOperation{se}
	v, _ := g.CreateVertex(obj)
	v.SetLabel("hi")

	g.UpdateVertex(v.ID, func(v *Vertex) error {
		v.SetLabel("label test")
		return nil
	})

	v, err := g.ReadVertex(v.ID)
	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	if v.Label() != "label test" {
		t.Fatalf("Expected 'label test' but was '%s'", v.Label())

	}
}

func Test_DeleteVertex(t *testing.T) {
	g := GraphOperation{se}
	v, err := g.CreateVertex(obj)
	err = g.DeleteVertex(v.ID)

	if err != nil {
		t.Fatalf("Expected err to be nil but was %s", err)
	}

	v, err = g.ReadVertex(v.ID)

	if err != errVertexNotFound {
		t.Fatalf("Expected err to be %s but was %s", errVertexNotFound, err)
	}
}
