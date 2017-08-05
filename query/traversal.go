package query

// Traversal decides how to excute the query
type Traversal struct {
}

// NewTraversal create a Traversal object used to run the query over the graph
func NewTraversal() *Traversal {
	return &Traversal{}
}

// Travers run's the query over the graph
func (t *Traversal) Travers(query *Query) error {
	// var next func() Next
	// next = query.path.Next

	return nil
}
