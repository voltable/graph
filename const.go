package graph

// CRUD Create, Read, Update and Delete
type CRUD int

const (
	// Create operation
	Create CRUD = iota
	// Read operation
	Read
	// Update operation
	Update
	// Delete operation
	Delete
)

// Digraph represents the directed or undirected relationship on a Edge
// a character consisting of two joined letters; a ligature.
type Digraph int

const (
	// Directed graphs have edges with direction. The edges indicate a one-way relationship, in that each edge can only be traversed in a single direction.
	Directed Digraph = iota
	// Undirected graphs have edges that do not have a direction. The edges indicate a two-way relationship, in that each edge can be traversed in both directions.
	Undirected
)

// EmptyString a string that is empty
const EmptyString = ""

var (
	vertex                 = []byte("v")
	label                  = []byte("l")
	properties             = []byte("p")
	relationship           = []byte("r")
	relationshipproperties = []byte("k")
	// US unit separator can be used as delimiters to mark fields of data structures. If used for hierarchical levels, US is the lowest level (dividing plain-text data items)
	US = []byte(string('\u241F'))
)
