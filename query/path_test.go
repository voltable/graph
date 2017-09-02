package query_test

import "github.com/RossMerr/Caudex.Graph/query"

// Root is the start of the path
type Root struct {
	next query.Path
}

var _ query.Path = (*Root)(nil)

// NewPath creates a Path object used to walk the graph
func NewPath() (query.Path, error) {
	path := &Root{}
	return path, nil
}

// Next returns the next Vertex or Edge in the Path
func (p *Root) Next() query.Path {
	return p.next
}

// SetNext sets the next Edge in the QueryPath
func (p *Root) SetNext(path query.Path) {
	if v, ok := path.(*query.PredicateVertexPath); ok {
		p.next = v
	}
}
