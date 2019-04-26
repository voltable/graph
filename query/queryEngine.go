package query

import graph "github.com/voltable/graph"

type QueryEngine interface {
	Parse(str string) (*graph.Query, error)
}
