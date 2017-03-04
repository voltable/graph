package graphs

import "testing"

// CREATE (a { name: 'Andres' })
// RETURN a
func Test_Structure(t *testing.T) {
	g := Graph{}

	v, _ := g.CreateVertex()

	v2, _ := g.CreateVertex()
	e, _ := v2.SetLabel("foo").AddDirectedEdge(v)
	e.SetLabel("bar")

	g.Query(v).Where(func(v *Vertex) bool {
		return v.Label() == "foo"
	})

}

type testObject struct {
	value string
}
