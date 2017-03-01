package graphs

import "testing"

// CREATE (a { name: 'Andres' })
// RETURN a
func Test_Structure(t *testing.T) {
	g := Graph{}

	obj := testObject{}
	v, _ := g.AddVertex(obj)

	obj2 := testObject{}
	v2, _ := g.AddVertex(obj2)
	e, _ := v2.SetLabel("foo").AddDirectedEdge(v)
	e.SetLabel("bar")

	all := All{}

	g.Query(v, all, all)

}

type testObject struct {
	value string
}
