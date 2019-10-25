package graph

import (
	"time"
)

// Query is the return object you'll get back from a traversal
type Query struct {
	Query      string
	Results   Table
	Length     time.Time
	Statistics Statistics
}

// NewQuery creates the initial Query object with the path create from the query string and path object
func NewQuery(query string, results *Table, statistics *Statistics) *Query {
	return &Query{
		Query: query,
		Results: *results,
		Statistics: *statistics,
	}
}

type Table struct {
	Columns []Column
}

func NewTable() *Table {
	return &Table {
		Columns: make([]Column, 0),
	}
}

func (s *Table) IsEmpty() bool {
	if len(s.Columns) > 0 {
		return false
	}
	return true
}

type Column struct {
	Field string
	Rows []interface{}
}