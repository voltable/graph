package query

// Query is the return object you'll get back from a traversal
type Query struct {
	query   string
	path    Path
	Results []interface{}
}

// NewQuery creates the instial Query object with the path create from the query string and path object
func NewQuery(path Path, query string, results []interface{}) *Query {
	return &Query{path: path, query: query, Results: results}
}
