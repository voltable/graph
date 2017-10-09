package ast

// // Digraph represents the directed or undirected relationship on a Edge
// // a character consisting of two joined letters; a ligature.
// type Digraph int

// const (
// 	// Undirected graphs have edges that do not have a direction. The edges indicate a two-way relationship, in that each edge can be traversed in both directions.
// 	Undirected Digraph = iota
// 	Inbound
// 	Outbound
// )

type Parentheses int

const (
	RPAREN Parentheses = iota // )
	LPAREN                    // (
)
