package internal

type CRUD int

const (
	Create CRUD = iota
	Read   CRUD = iota
	Update CRUD = iota
	Delete CRUD = iota
)

// Digraph, represents the directed or undirected relationship on a Edge
// a character consisting of two joined letters; a ligature.
type Digraph int

const (
	// Directed graphs have edges with direction. The edges indicate a one-way relationship, in that each edge can only be traversed in a single direction.
	Directed Digraph = iota
	// Undirected graphs have edges that do not have a direction. The edges indicate a two-way relationship, in that each edge can be traversed in both directions.
	Undirected Digraph = iota
)
