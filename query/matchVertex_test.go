package query

import "testing"
import "github.com/RossMerr/Caudex.Graph/vertices"

func Test_MatchVertex(t *testing.T) {

	it := func() (item interface{}, ok bool) {
		return "", true
	}

	p := VertexPath{Iterate: func() Iterator {
		return it
	}}

	p.MatchVertex(func(v *vertices.Vertex) bool {
		return v.Label() == "foo"
	})
}
