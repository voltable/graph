package query

import graph "github.com/RossMerr/Caudex.Graph"

type QueryEngine interface {
	Parse(str string) (*graph.Query, error)
}
