package main

import (
	"github.com/RossMerr/Caudex.Graph"
)

func main() {
	var g graph.Graph
	var err error
	if g, err = graph.NewGraph("bolt", &graph.Options{Name: "test"}); err != nil {
		panic(err)
	}
	//	q, err := g.Query("")

	// slice := q.Node(func(v *vertices.Vertex) bool {
	// 	return v.Label() == "foo"
	// }).Relationship(func(e *vertices.Edge) bool {
	// 	return true
	// }).Node(func(v *vertices.Vertex) bool {
	// 	return v.Label() == "foo"
	// }).ToSlice()

	// for _, v := range slice {
	// 	fmt.Println(v)
	// }

	g.Close()
}
