package query

// Query is the return object you'll get back from a traversal
type Query struct {
	query   string
	path    Path
	Results []interface{}
}

func NewQuery(path Path) *Query {
	return &Query{path: path}
}
