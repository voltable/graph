package memorydb_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/vertices"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/storage/memorydb"
)

func Test_Query(t *testing.T) {
	cypher.RegisterEngine()
	options := graph.NewOptions()
	g, err := memorydb.NewStorageEngine(options)

	if err != nil {
		t.Errorf("Failed to create the storageEngine %v", err)
	}

	v1, _ := vertices.NewVertex()
	v1.SetLabel("person")
	v1.SetProperty("name", "john smith")
	g.Create(v1)

	v2, _ := vertices.NewVertex()
	v2.SetLabel("place")
	v2.SetProperty("name", "london")
	g.Create(v2)

	q, err := g.Query("MATCH (n:Person) WHERE n.name = 'john smith'")

	if err != nil {
		t.Errorf("Bad Query")
	}

	if len(q.Results()) != 1 {
		t.Errorf("Failed to match")
	}
}
