package memorydb_test

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/vertices"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/storage/memorydb"
)

func Test_Query(t *testing.T) {
	cypher.RegisterQueryEngine()
	options := graph.NewOptions()
	g, err := memorydb.NewStorageEngine(options)

	if err != nil {
		t.Errorf("Failed to create the storageEngine %v", err)
	}

	v, _ := vertices.NewVertex()
	g.Create(v)
}
