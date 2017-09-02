package query

// Query is the return object you'll get back from a traversal
type Query struct {
	Query   string
	Path    Path
	Results []interface{}
}

// NewQuery creates the instial Query object with the path create from the query string and path object
func NewQuery(path Path, query string, results []interface{}) *Query {
	return &Query{Path: path, Query: query, Results: results}
}
