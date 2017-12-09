package graph

// Query is the return object you'll get back from a traversal
type Query struct {
	Query   string
	Results []interface{}
}

// NewQuery creates the instial Query object with the path create from the query string and path object
func NewQuery(query string, results []interface{}) *Query {
	return &Query{Query: query, Results: results}
}
