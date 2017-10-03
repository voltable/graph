package memorydb_test

import (
	"testing"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/storage/memorydb"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_Query(t *testing.T) {
	cypher.RegisterEngine()
	options := graph.NewOptions()
	var g graph.Graph
	var err error

	var tests = []struct {
		setup   func()
		query   string
		results int
	}{
		{
			setup: func() {
				v1, _ := vertices.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				g.Create(v1)

				v2, _ := vertices.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "foo bar")
				g.Create(v2)
			},
			query:   "MATCH (n:person) WHERE n.name = 'john smith'",
			results: 1,
		},
	}

	for i, tt := range tests {
		g, err = memorydb.NewStorageEngine(options)
		tt.setup()
		if err != nil {
			t.Errorf("Failed to create the storageEngine %v", err)
		}

		q, err := g.Query(tt.query)

		if err != nil {
			t.Errorf("%d. Bad Query \n%v", i, tt.query)
		}

		if len(q.Results) != 1 {
			t.Errorf("Failed to match expected 1 got %v", len(q.Results))
		}
		reults := len(q.Results)
		if reults != tt.results {
			t.Errorf("%d. expected %d got %d", i, tt.results, reults)
		}
	}

}
