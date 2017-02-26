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

	id = v.ID
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
