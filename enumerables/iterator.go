package enumerables

import "github.com/RossMerr/Caudex.Graph/vertices"

// Iterator is an alias for function to iterate over data.
type Iterator func() (item interface{}, ok bool)

// IteratorVertex is an alias for function to iterate over Vertex.
type IteratorVertex func() (item *vertices.Vertex, ok bool)

// IteratorEdge is an alias for function to iterate over Vertex.
type IteratorEdge func() (item *vertices.Edge, ok bool)
