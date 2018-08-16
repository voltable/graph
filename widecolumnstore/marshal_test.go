package widecolumnstore_test

import (
	"reflect"
	"testing"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

func TestMarshalKeyValue(t *testing.T) {
	tests := []struct {
		name   string
		vertex *graph.Vertex
		want   []*widecolumnstore.KeyValue
	}{
		{
			name: "vertex",
			vertex: func() *graph.Vertex {
				v, _ := graph.NewVertex()
				v.SetLabel("person")
				v.SetProperty("name", "john smith")
				t, _ := graph.NewVertex()
				e := graph.NewEdge(v, t)
				e.SetRelationshipType("friend")
				e.SetProperty("years", 10)
				e.Weight = 5
				v.AddEdge(e)
				return v
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, transpose := widecolumnstore.MarshalKeyValue(tt.vertex)
			v := graph.NewEmptyVertex()
			widecolumnstore.UnmarshalKeyValue(v, got)
			if !reflect.DeepEqual(v, tt.vertex) {
				t.Errorf("Marshal() = %v, want %v", v, tt.vertex)
			}

			// The transpose
			v = graph.NewEmptyVertex()
			widecolumnstore.UnmarshalKeyValueTranspose(v, transpose)
			if !reflect.DeepEqual(v, tt.vertex) {
				t.Errorf("Marshal() = %v, want %v", v, tt.vertex)
			}
		})
	}
}
