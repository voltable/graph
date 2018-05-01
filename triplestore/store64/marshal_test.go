package store64_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/triplestore/store64"

	graph "github.com/RossMerr/Caudex.Graph"
)

func TestMarshal(t *testing.T) {
	type args struct {
		c []*graph.Vertex
	}
	tests := []struct {
		name string
		args args
		want func([]*graph.Vertex) []*store64.Triple
	}{
		{
			args: args{
				c: func() []*graph.Vertex {
					vertices := make([]*graph.Vertex, 0)
					vertex, _ := graph.NewVertexWithLabel("person")
					vertex.SetProperty("name", "john")
					target, _ := graph.NewVertexWithLabel("person")
					edge, _, _ := vertex.AddEdgeWeight(target, 0.5)
					edge.SetRelationshipType("friend")
					edge.SetProperty("known", "4 years")
					vertices = append(vertices, vertex)
					return vertices
				}(),
			},
			want: func(in []*graph.Vertex) []*store64.Triple {
				tt := make([]*store64.Triple, 0)
				delimiter := string(store64.Delimiter)
				tt = append(tt, store64.NewTriple(store64.Vertex+delimiter+in[0].ID(), "person", float64(1)))
				tt = append(tt, store64.NewTriple(store64.VertexProperties+delimiter+"name"+delimiter+in[0].ID(), "john", float64(1)))
				tt = append(tt, store64.NewTriple(store64.Edge+delimiter+"friend"+delimiter+in[0].ID(), in[0].Edges()[0].ID(), float64(0.5)))
				tt = append(tt, store64.NewTriple(store64.EdgeProperties+delimiter+"known"+delimiter+"friend"+delimiter+in[0].ID(), "4 years", float64(1)))
				return tt
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triples := tt.want(tt.args.c)
			if got := store64.Marshal(tt.args.c...); !reflect.DeepEqual(got, triples) {
				t.Errorf("Marshal() = \n\tgot \n%v, \n\twant \n%v", got, triples)
			}
		})
	}
}
