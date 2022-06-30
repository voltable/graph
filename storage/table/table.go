package table

import (
	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/constraints"
	"github.com/voltable/graph/storage"
	"github.com/voltable/graph/storage/csr"
)

type Table[T constraints.Number] struct {
	row    map[string]int
	column map[string]int
	matrix graphblas.MatrixCompressed[T]
}

type Vector[T constraints.Number] struct {
	arr    map[string]int
	vector graphblas.Vector[T]
}

func (v *Vector[T]) At(key string) T {
	index := v.arr[key]
	return v.vector.AtVec(index)
}

func NewTable[T constraints.Number](r, c int) *Table[T] {
	return &Table[T]{
		row:    make(map[string]int),
		column: make(map[string]int),
		matrix: graphblas.NewCSRMatrix[T](r, c),
	}
}

func (t *Table[T]) ColumnsAt(subject string) *Vector[T] {
	c := t.column[subject]
	return &Vector[T]{
		vector: t.matrix.ColumnsAt(c),
	}

}

func (t *Table[T]) RowsAt(predicate string) *Vector[T] {
	c := t.row[predicate]
	return &Vector[T]{
		vector: t.matrix.RowsAt(c),
	}
}

func (t *Table[T]) Create(triples ...storage.Triple[T]) []*csr.Row {
	rows := []*csr.Row{}
	for i := 0; i < len(triples); i++ {
		triple := triples[i]
		if c, ok := t.column[triple.Subject]; ok {
			if r, ok := t.row[triple.Predicate]; ok {
				pointer, start := t.matrix.SetReturnPointer(r, c, triple.Object)
				rows = append(rows, &csr.Row{
					Value:    triple.Object,
					Col:      pointer,
					RowStart: start,
				})
			}
		}
	}

	return rows
}
