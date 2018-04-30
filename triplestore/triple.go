package triplestore

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"

	"github.com/RossMerr/Caudex.Graph/container/table"
)

var typeRegister map[string]reflect.Type

func init() {
	typeRegister = make(map[string]reflect.Type)
	typeRegister[reflect.Bool.String()] = reflect.TypeOf(bool(false))
	typeRegister[reflect.Int.String()] = reflect.TypeOf(int(0))
	typeRegister[reflect.Int8.String()] = reflect.TypeOf(int8(0))
	typeRegister[reflect.Int16.String()] = reflect.TypeOf(int16(0))
	typeRegister[reflect.Int32.String()] = reflect.TypeOf(int32(0))
	typeRegister[reflect.Int64.String()] = reflect.TypeOf(int64(0))
	typeRegister[reflect.Uint.String()] = reflect.TypeOf(uint(0))
	typeRegister[reflect.Uint8.String()] = reflect.TypeOf(uint8(0))
	typeRegister[reflect.Uint16.String()] = reflect.TypeOf(uint16(0))
	typeRegister[reflect.Uint32.String()] = reflect.TypeOf(uint32(0))
	typeRegister[reflect.Uint64.String()] = reflect.TypeOf(uint64(0))
	typeRegister[reflect.Float32.String()] = reflect.TypeOf(float32(0))
	typeRegister[reflect.Float64.String()] = reflect.TypeOf(float64(0))
	typeRegister[reflect.Complex64.String()] = reflect.TypeOf(complex64(0))
	typeRegister[reflect.Complex128.String()] = reflect.TypeOf(complex128(0))
	typeRegister[reflect.String.String()] = reflect.TypeOf("")
}

func NewAny(i interface{}) (*Any, error) {
	t := reflect.TypeOf(i)

	// floats are not supported by the binary.Write 😕
	var b [8]byte
	var bs []byte
	n := intDataSize(i)
	if n > len(b) {
		bs = make([]byte, n)
	} else {
		bs = b[:n]
	}

	if f, ok := i.(float32); ok {
		float32ToByte(bs, f)
	} else if f, ok := i.(float64); ok {
		float64ToByte(bs, f)
	} else {
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, i)

		if err != nil {
			return nil, err
		}

		bs = buf.Bytes()
	}

	any := &Any{
		Value: bs,
		Type:  t.Name(),
	}
	return any, nil
}

// NewTriplesFromTable returns a []*Triple
func NewTriplesFromTable(t table.Table) []*Triple {
	tt := make([]*Triple, 0)

	t.ReadAll()

	t.Iterator(func(r, c string, v interface{}) {
		any := &Any{
			Value: []byte(c),
			Type:  "string",
		}

		triple := &Triple{Row: r, Column: c, Value: any}
		tt = append(tt, triple)
	})

	return tt
}

// NewTripleTransposeFromTable returns a []*Triple transposed
func NewTripleTransposeFromTable(t table.Table) []*Triple {
	tt := make([]*Triple, 0)

	t.ReadAll()

	t.Iterator(func(r, c string, v interface{}) {

		t := reflect.TypeOf(v)

		any := &Any{
			Value: []byte(c),
			Type:  t.Name(),
		}

		triple := &Triple{Row: c, Column: r, Value: any}
		tt = append(tt, triple)
	})

	return tt
}

// Transpose swap the row's and column's
func Transpose(tt []*Triple) []*Triple {
	triples := make([]*Triple, 0)

	for _, t := range tt {
		triple := &Triple{Row: t.Column, Column: t.Row, Value: t.Value}
		triples = append(triples, triple)
	}

	return triples
}

func (s *Any) Interface() (interface{}, error) {
	i := typeRegister[s.Type]
	v := reflect.New(i).Elem()

	if i == reflect.TypeOf(float32(0)) {
		return bytesToFloat32(s.Value), nil
	} else if i == reflect.TypeOf(float64(0)) {
		return bytesToFloat64(s.Value), nil
	}

	r := bytes.NewReader(s.Value)
	err := binary.Read(r, binary.LittleEndian, &v)
	return v, err
}

func float64ToByte(b []byte, f float64) {
	binary.LittleEndian.PutUint64(b[:], math.Float64bits(f))
}

func float32ToByte(b []byte, f float32) {
	binary.LittleEndian.PutUint32(b[:], math.Float32bits(f))
}

func bytesToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b))
}

func bytesToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func intDataSize(data interface{}) int {
	switch data := data.(type) {
	case float32, *float32:
		return 4
	case []float32:
		return 4 * len(data)
	case float64, *float64:
		return 8
	case []float64:
		return 8 * len(data)
	}
	return 0
}
