package memorydb_test

import (
	"reflect"
	"testing"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/storage/memorydb"
	"github.com/RossMerr/Caudex.Graph/vertices"
)

func Test_Query(t *testing.T) {
	cypher.RegisterEngine()
	options := graph.NewOptions()

	var tests = []struct {
		matching    []*vertices.Vertex
		nonMatching []*vertices.Vertex
		query       string
	}{
		{
			matching: func() []*vertices.Vertex {
				arr := make([]*vertices.Vertex, 0, 0)
				v1, _ := vertices.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)
				return arr
			}(),
			nonMatching: func() []*vertices.Vertex {
				arr := make([]*vertices.Vertex, 0, 0)
				v2, _ := vertices.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "foo bar")
				arr = append(arr, v2)
				return arr
			}(),
			query: "MATCH (n:person) WHERE n.name = 'john smith'",
		},
	}

	for i, tt := range tests {
		g, err := memorydb.NewStorageEngine(options)
		g.Create(tt.matching...)
		g.Create(tt.nonMatching...)

		if err != nil {
			t.Errorf("Failed to create the storageEngine %v", err)
		}

		q, err := g.Query(tt.query)

		if err != nil {
			t.Errorf("%d. Bad Query \n%v", i, tt.query)
		}

		for ii, r := range q.Results {
			match := false
			for _, m := range tt.matching {
				if reflect.DeepEqual(r.(*vertices.Vertex), m) {
					match = true
					break
				}
			}

			if !match {
				t.Errorf("%d. result %d not matched \n%+v ", i, ii, r)
			}
		}
	}

}
