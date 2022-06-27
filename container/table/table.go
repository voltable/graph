package table

import (
	"bufio"
	"io"
	"strings"

	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/constraints"
	"github.com/voltable/graph"
)

const (
	stringEmpty string = ""
)

// Table is a set of data elements using a model of columns and rows
type Table[T constraints.Number] interface {
	ReadAll() error
	Iterator(i func(string, string, interface{})) bool
	Columns() int
	Rows() int
	get(r, c string) interface{}
	Get(r, c string) T
}

type table[T constraints.Number] struct {
	matrix        graphblas.Matrix[T]
	rowIndices    []string
	columnIndices []string
	columns       map[string]int
	delimiter     rune
	reader        *container
}

// NewTableFromReader returns a table.Table
func NewTableFromReader[T constraints.Number](r, c int, reader io.Reader) Table[T] {
	return &table[T]{
		matrix:        graphblas.NewCSCMatrix[T](r, c),
		rowIndices:    make([]string, r),
		columnIndices: make([]string, c),
		columns:       make(map[string]int, c),
		delimiter:     '|',
		reader: &container{
			text: bufio.NewReader(reader),
		},
	}
}

func (s *table[T]) read(header []string, r int, row []string) {
	indice := header[0]
	s.rowIndices[r] = indice + string(s.delimiter) + row[0]

	for i := 1; i < len(row); i++ {
		// Column header name
		uniqueTypeValuePair := header[i] + string(s.delimiter) + row[i]
		v := graph.Default[T]() + 1

		if c, ok := s.columns[uniqueTypeValuePair]; ok {
			v += s.matrix.At(r, c)
			s.matrix.Set(r, c, v)
		} else {
			c = len(s.columns)
			s.columns[uniqueTypeValuePair] = c
			s.columnIndices[c] = uniqueTypeValuePair
			s.matrix.Set(r, c, v)
		}
	}
}

func (s *table[T]) ReadAll() error {
	// Read the header
	line, err := s.reader.Read()
	if err != nil {
		return err
	}

	header := line

	// Read the body
	count := 0
	for {
		line, err := s.reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		s.read(header, count, line)
		count++
	}

	return nil
}

// Columns the number of columns of the matrix
func (s *table[T]) Columns() int {
	return s.matrix.Columns()
}

// Rows the number of rows of the matrix
func (s *table[T]) Rows() int {
	return s.matrix.Rows()
}

// Get (unoptimized) returns the value of a table element at r-th, c-th
func (s *table[T]) get(r, c string) interface{} {
	cPointer := s.columns[c]
	rPointer := -1
	for i, value := range s.rowIndices {
		if value == r {
			rPointer = i
			break
		}
	}

	return s.matrix.At(rPointer, cPointer)
}

func (s *table[T]) Get(r, c string) T {
	v := s.get(r, c)
	if value, ok := v.(T); ok {
		return value
	}
	return graph.Zero[T]()
}

// Iterator iterates through all non-zero elements, order is not guaranteed
func (s *table[T]) Iterator(i func(string, string, interface{})) bool {
	enumerator := s.matrix.Enumerate()
	if enumerator.HasNext() {
		r, c, v := enumerator.Next()
		i(s.rowIndices[r], s.columnIndices[c], v)
		return true
	}

	return false
}

type container struct {
	text *bufio.Reader
}

func (s *container) readLine() (line string, err error) {
	b, _, err := s.text.ReadLine()

	if err != nil {
		return stringEmpty, err
	}

	return string(b), nil
}

func (s *container) Read() (record []string, err error) {
	line, err := s.readLine()
	split := strings.Split(line, " ")
	return split, err
}
