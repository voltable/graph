package csr

import "github.com/voltable/graph/storage"

type Row struct {
	Value    interface{}
	Col      int
	RowStart int
}

func NewRowFromTriple(triple storage.Triple) storage.SparseValue {
	return &Row{}
}
