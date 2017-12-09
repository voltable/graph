package memorydb_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/storage/memorydb"
)

func Test_Query(t *testing.T) {
	cypher.RegisterEngine()
	options := graph.NewOptions()

	var tests = []struct {
		expecting    []interface{}
		uninterested []*graph.Vertex
		query        string
	}{
		// 0
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)
				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "foo bar")
				arr = append(arr, v2)
				return arr
			}(),
			query: "MATCH (n:person) WHERE n.name = 'john smith'",
		},
		// 1
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)
				fmt.Printf("%+v", v1.ID())
				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)
				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "foo bar")
				arr = append(arr, v2)
				return arr
			}(),
			query: "MATCH (n:person) WHERE n.name = 'john smith' RETURN *",
		},
		// 2
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "john smith")
				arr = append(arr, v2)
				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)
				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "foo bar")
				arr = append(arr, v2)
				return arr
			}(),
			query: "MATCH (n:person) WHERE n.name = 'john smith' RETURN *",
		},
		// 3
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("location", "london")
				arr = append(arr, v2)
				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)
				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "foo bar")
				arr = append(arr, v2)
				return arr
			}(),
			query: "MATCH (n:person) WHERE n.name = 'john smith' OR n.location = 'london' RETURN *",
		},
		// 4
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)

				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)

				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "foo bar")
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()
				v2.SetLabel("location")
				v2.SetProperty("address", "london")
				arr = append(arr, v2)
				return arr
			}(),
			query: "MATCH (n:person) WHERE n.name = 'john smith' OR n.location = 'london' RETURN *",
		},
		// 5
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("age", 18)
				arr = append(arr, v2)

				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)

				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "foo bar")
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()
				v2.SetLabel("location")
				v2.SetProperty("address", "london")
				arr = append(arr, v2)
				return arr
			}(),
			query: "MATCH (n:person) WHERE n.name = 'john smith' OR n.age = 18 RETURN *",
		},
		// 6
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()

				v2.SetLabel("person")
				v2.SetProperty("name", "max power")
				arr = append(arr, v2)

				edge, _ := v1.AddDirectedEdge(v2)
				edge.SetRelationshipType("knows")
				arr = append(arr, edge)

				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)

				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "foo bar")
				arr = append(arr, v1)

				return arr
			}(),
			query: "MATCH (n:person)-[:knows]->(m:person) RETURN *",
		},
		// 7
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)

				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "max power")
				arr = append(arr, v2)

				edge, _ := v1.AddDirectedEdge(v2)
				edge.SetRelationshipType("knows")

				v3, _ := graph.NewVertex()
				v3.SetLabel("person")
				v3.SetProperty("name", "foo bar")
				arr = append(arr, v3)

				return arr
			}(),
			query: "MATCH (n:person)-[:notknows]->(m:person) RETURN *",
		},
		// 8
		{
			expecting: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				arr = append(arr, v1)

				return arr
			}(),
			uninterested: func() []*graph.Vertex {
				arr := make([]*graph.Vertex, 0, 0)

				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "foo bar")
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()
				v2.SetLabel("location")
				v2.SetProperty("address", "london")
				arr = append(arr, v2)
				return arr
			}(),
			query: "MATCH (n:person{name:'john smith'}) RETURN *",
		},
	}
	for i, tt := range tests {
		g, err := memorydb.NewStorageEngine(options)
		if err != nil {
			t.Errorf("Failed to create the storageEngine %v", err)
		}

		for _, i := range tt.expecting {
			if v, ok := i.(*graph.Vertex); ok {
				g.Create(v)
			}
		}
		g.Create(tt.uninterested...)
		q, err := g.Query(tt.query)

		if err != nil {
			t.Errorf("%d. Bad Query \n%v", i, tt.query)
		}

		if len(q.Results) != len(tt.expecting) {
			t.Errorf("%d. expected %d got %d", i, len(tt.expecting), len(q.Results))
		}

		for ii, r := range q.Results {
			match := false
			for _, m := range tt.expecting {
				if reflect.DeepEqual(r, m) {
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

func Test_QueryRelationships(t *testing.T) {
	cypher.RegisterEngine()
	options := graph.NewOptions()

	var tests = []struct {
		setup     []interface{}
		expecting func([]interface{}) []interface{}
		query     string
	}{
		// 0
		{
			expecting: func(in []interface{}) []interface{} {
				arr := make([]interface{}, 0, 0)
				arr = append(arr, in[0])
				return arr
			},
			setup: func() []interface{} {
				arr := make([]interface{}, 0, 0)
				v1, _ := graph.NewVertex()
				v1.SetLabel("person")
				v1.SetProperty("name", "john smith")
				//v1.SetProperty("male", true)
				arr = append(arr, v1)

				v2, _ := graph.NewVertex()
				v2.SetLabel("person")
				v2.SetProperty("name", "max power")
				arr = append(arr, v2)

				edge, _ := v1.AddDirectedEdge(v2)
				edge.SetRelationshipType("knows")
				arr = append(arr, edge)

				v3, _ := graph.NewVertex()
				v3.SetLabel("person")
				v3.SetProperty("name", "foo bar")
				arr = append(arr, v3)

				return arr
			}(),
			query: "MATCH (n:person)-[*]->(m:person) RETURN n",
		},
		// // 1
		// {
		// 	expecting: func(in []*vertices.Vertex) []*vertices.Vertex {
		// 		arr := make([]*vertices.Vertex, 0, 0)
		// 		arr = append(arr, in[0])
		// 		arr = append(arr, in[1])
		// 		return arr
		// 	},
		// 	setup: func() []*vertices.Vertex {
		// 		arr := make([]*vertices.Vertex, 0, 0)

		// 		v1, _ := vertices.NewVertex()
		// 		v1.SetLabel("person")
		// 		v1.SetProperty("name", "john smith")
		// 		arr = append(arr, v1)

		// 		v2, _ := vertices.NewVertex()
		// 		v2.SetLabel("person")
		// 		v2.SetProperty("name", "max power")
		// 		arr = append(arr, v2)

		// 		edge, _ := v1.AddDirectedEdge(v2)
		// 		edge.SetRelationshipType("knows")

		// 		v3, _ := vertices.NewVertex()
		// 		v3.SetLabel("person")
		// 		v3.SetProperty("name", "foo bar")
		// 		arr = append(arr, v3)

		// 		return arr
		// 	}(),
		// 	query: "MATCH (n:person)-[*1]->(m:person) ",
		// },
		// // 2
		// {
		// 	expecting: func(in []*vertices.Vertex) []*vertices.Vertex {
		// 		arr := make([]*vertices.Vertex, 0, 0)
		// 		arr = append(arr, in[0])
		// 		arr = append(arr, in[2])
		// 		return arr
		// 	},
		// 	setup: func() []*vertices.Vertex {
		// 		arr := make([]*vertices.Vertex, 0, 0)

		// 		v1, _ := vertices.NewVertex()
		// 		v1.SetLabel("person")
		// 		v1.SetProperty("name", "john smith")
		// 		arr = append(arr, v1)

		// 		v2, _ := vertices.NewVertex()
		// 		v2.SetLabel("person")
		// 		v2.SetProperty("name", "max power")
		// 		arr = append(arr, v2)

		// 		edge2, _ := v1.AddDirectedEdge(v2)
		// 		edge2.SetRelationshipType("notknows")

		// 		v3, _ := vertices.NewVertex()
		// 		v3.SetLabel("person")
		// 		v3.SetProperty("name", "foo bar")
		// 		arr = append(arr, v3)

		// 		edge3, _ := v2.AddDirectedEdge(v3)
		// 		edge3.SetRelationshipType("knows")

		// 		return arr
		// 	}(),
		// 	query: "MATCH (n:person)-[*2]->(m:person) ",
		// },
		// // 3
		// {
		// 	expecting: func() []*vertices.Vertex {
		// 		arr := make([]*vertices.Vertex, 0, 0)
		// 		v1, _ := vertices.NewVertex()
		// 		v1.SetLabel("person")
		// 		v1.SetProperty("name", "john smith")
		// 		arr = append(arr, v1)

		// 		v2, _ := vertices.NewVertex()
		// 		v2.SetLabel("person")
		// 		v2.SetProperty("name", "max power")
		// 		arr = append(arr, v2)

		// 		edge, _ := v1.AddDirectedEdge(v2)
		// 		edge.SetRelationshipType("knows")

		// 		return arr
		// 	}(),
		// 	uninterested: func() []*vertices.Vertex {
		// 		arr := make([]*vertices.Vertex, 0, 0)

		// 		v1, _ := vertices.NewVertex()
		// 		v1.SetLabel("person")
		// 		v1.SetProperty("name", "foo bar")
		// 		arr = append(arr, v1)

		// 		return arr
		// 	}(),
		// 	query: "MATCH (n:person)-[:knows]->(m:person) WHERE n.name = 'john smith' OR m.name = 'max power'",
		// },
	}

	for i, tt := range tests {
		g, err := memorydb.NewStorageEngine(options)
		if err != nil {
			t.Errorf("Failed to create the storageEngine %v", err)
		}

		for _, i := range tt.setup {
			if v, ok := i.(*graph.Vertex); ok {
				g.Create(v)
			}
		}

		q, err := g.Query(tt.query)

		if err != nil {
			t.Errorf("%d. Bad Query \n%v", i, tt.query)
		}

		expecting := tt.expecting(tt.setup)
		if len(q.Results) != len(expecting) {
			t.Errorf("%d. expected %d got %d", i, len(expecting), len(q.Results))
		}

		for ii, r := range q.Results {
			match := false
			for _, m := range expecting {

				// if strings.Compare(r, m) == 0 {
				// 	match = true
				// 	break
				// }
				fmt.Printf("%+v", reflect.TypeOf(r))
				if reflect.DeepEqual(r, m) {
					match = true
					break
				}
			}

			if !match {
				t.Errorf("%d. result %d not matched got \n%+v", i, ii, r)
			}
		}
	}

}
