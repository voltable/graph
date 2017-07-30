package query

// Query is the return object you'll get back from a traversal
type Query struct {
	query   string
	path    QueryPath
	Results []interface{}
}
