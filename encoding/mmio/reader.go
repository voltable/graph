package mmio

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/rossmerr/graphblas"
	triples "github.com/voltable/graph/triplestore/store64"
)

const (
	stringEmpty string = ""
	comment     string = "%"

	// start of header
	header string = comment + comment + "MatrixMarket"

	// object
	matrix     string = "matrix"
	coordinate string = "coordinate"

	// data
	real    string = "real"
	complex string = "complex"
	integer string = "integer"
	pattern string = "pattern"

	// symmetry
	general       string = "general"
	symmetric     string = "symmetric"
	skewSymmetric string = "skew-symmetric"
	hermitian     string = "hermitian"
)

// http://math.nist.gov/MatrixMarket/

// Reader Matrix Market I/O (MMIO) file format
type Reader struct {
	text *bufio.Reader
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	reader := &Reader{
		text: bufio.NewReader(r),
	}

	return reader
}

func (s *Reader) readLine() (line string, err error) {
	b, _, err := s.text.ReadLine()

	if err != nil {
		return stringEmpty, err
	}

	return string(b), nil
}

// Read reads one record (a slice of fields) from r.
func (s *Reader) read() (r, c int, value float64, err error) {
	line, err := s.readLine()

	record := strings.Split(strings.TrimSpace(line), " ")
	if err != nil {
		return
	}

	if strings.HasPrefix(record[0], comment) {
		return s.read()
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

func (s *Reader) header() (string, error) {
	line, err := s.readLine()

	record := strings.Split(strings.TrimSpace(line), " ")

	if err != nil {
		return stringEmpty, err
	}

	if len(record) != 5 {
		return stringEmpty, fmt.Errorf("Invalid header")
	}

	if record[0] != header {
		return stringEmpty, fmt.Errorf("No Matrix Market header")
	}

	if record[1] != matrix {
		return stringEmpty, fmt.Errorf("Unknown object '%+v' only 'matrix' supported", record[1])
	}

	if record[2] != coordinate {
		return stringEmpty, fmt.Errorf("Unknown representation '%+v' only 'coordinate' supported", record[2])
	}

	if record[3] != real && record[3] != complex && record[3] != integer {
		return stringEmpty, fmt.Errorf("Unknown data type '%+v' only 'integer', real' or 'complex'", record[3])
	}

	if record[4] != general && record[4] != symmetric && record[4] != skewSymmetric && record[4] != hermitian {
		return stringEmpty, fmt.Errorf("Unknown symmetry '%+v' only 'general', 'symmetric', 'skew-symmetric' or 'hermitian'", record[4])
	}

	return record[2], nil
}

// ReadToMatrix reads all records from r and returns a Matrix
func (s *Reader) ReadToMatrix() (graphblas.Matrix[float64], error) {

	_, err := s.header()

	if err != nil {
		return nil, err
	}

	r, c, value, err := s.read()

	if err != nil {
		return nil, err
	}

	matrix := graphblas.NewCSCMatrix[float64](r, c)

	for {
		r, c, value, err = s.read()

		if err == io.EOF {
			return matrix, nil
		} else if err != nil {
			return nil, err
		}

		matrix.Set(r-1, c-1, value)
	}
}

// ReadToTriples reads all records from r and returns a Triples
func (s *Reader) ReadToTriples() ([]*triples.Triple, error) {
	_, err := s.header()

	if err != nil {
		return nil, err
	}

	r, c, value, err := s.read()

	if err != nil {
		return nil, err
	}

	tt := make([]*triples.Triple, 0)

	for {
		r, c, value, err = s.read()

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
