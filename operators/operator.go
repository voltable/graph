package operators

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/widecolumnstore"
)

// Operator usually have 0, 1, or 2 arguments
type Operator interface {
	Op()
}

// Nullary an operation of arity 0, and hence call it nullary
type Nullary interface {
	Operator
	Next() (widecolumnstore.Iterator, graph.Statistics)
}

// Unary operation takes one argument
type Unary interface {
	Operator
	Next(i widecolumnstore.Iterator) widecolumnstore.Iterator
}

// Binary operation takes two arguments
type Binary interface {
	Operator
	Next(x, y widecolumnstore.Iterator) widecolumnstore.Iterator
}
