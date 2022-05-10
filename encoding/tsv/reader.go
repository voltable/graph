package tsv

import (
	"encoding/binary"
	"encoding/csv"
	"io"
	"math"
	"strconv"

	"github.com/rossmerr/graphblas"
	GraphBLAS "github.com/rossmerr/graphblas"
	triples "github.com/voltable/graph/triplestore/store64"
)

// Reader Tab-Separated Values (TSV) file format
// (Row, Col, Value) tuple describing the adjacency matrix of the graph.
type Reader struct {
	csv *csv.Reader
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	reader := &Reader{
		csv: csv.NewReader(r),
	}

	reader.csv.Comma = '\t'
	return reader
}

// Read reads one record (a slice of fields) from r.
func (s *Reader) read() (r, c int, value float64, err error) {
	record, err := s.csv.Read()

	if err != nil {
		return
	}

	r, err = strconv.Atoi(record[0])
	if err != nil {
		return
	}

	c, err = strconv.Atoi(record[1])
	if err != nil {
		return
	}

	value, err = strconv.ParseFloat(record[2], 64)
	return
}

// ReadToMatrix reads all records from r.
func (s *Reader) ReadToMatrix() (graphblas.Matrix[float64], error) {
	columnMax := 0
	matrix := [][]float64{}
	for {
		r, c, value, err := s.read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if columnMax < c {
			columnMax = c
		}

		if len(matrix) < r {
			count := r - len(matrix)
			for i := 0; i < count; i++ {
				matrix = append(matrix, make([]float64, columnMax))
			}
		}

		if len(matrix[r-1]) < c {
			count := columnMax - len(matrix[r-1])
			for i := 0; i < count; i++ {
				matrix[r-1] = append(matrix[r-1], 0)
			}
		}

		matrix[r-1][c-1] = value
	}

	// Set all zero elements in the matrix
	for r := range matrix {
		if len(matrix[r]) < columnMax {
			count := columnMax - len(matrix[r])
			for i := 0; i < count; i++ {
				matrix[r] = append(matrix[r], 0)
			}
		}
	}

	graph := GraphBLAS.NewDenseMatrixFromArray(matrix)

	return graph, nil
}

// ReadToTriples reads all records from r and returns a Triples
func (s *Reader) ReadToTriples() ([]*triples.Triple, error) {
	tt := make([]*triples.Triple, 0)
	for {
		r, c, value, err := s.read()

		if err == io.EOF {
			return tt, nil
		} else if err != nil {
			return tt, err
		}

		tt = append(tt, &triples.Triple{
			Row:    strconv.Itoa(r),
			Column: strconv.Itoa(c),
			Value:  value,
		})

	}
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}
