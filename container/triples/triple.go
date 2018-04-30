package triples

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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

	buf := new(bytes.Buffer)
	err := write(buf, binary.LittleEndian, i)
	//err := binary.Write(buf, binary.BigEndian, i)
	if err != nil {
		return nil, err
	}

	any := &Any{
		Value: buf.Bytes(),
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
	i := makeInstance(s.Type)
	//i := float64(0)
	r := bytes.NewReader(s.Value)
	err := read(r, binary.LittleEndian, i)
	fmt.Printf("\nfloat64 %+v", i)
	//err := binary.Read(r, binary.BigEndian, i)
	return i, err
}

func makeInstance(name string) interface{} {
	//t := typeRegister[name]
	f := float64(5)
	t := reflect.TypeOf(f)
	v := reflect.New(t)

	if f, err := extractFloat64(v); err == nil {
		return f
	}

	return v.Interface()
}

func extractFloat64(v reflect.Value) (float64, error) {
	if v.Kind() != reflect.Float64 {
		return float64(0), errors.New("Invalid input")
	}
	var floatVal float64
	floatVal = v.Float()
	return floatVal, nil
}

func write(w io.Writer, order binary.ByteOrder, data interface{}) error {
	if n := intDataSize(data); n != 0 {
		var b [8]byte
		var bs []byte
		if n > len(b) {
			bs = make([]byte, n)
		} else {
			bs = b[:n]
		}
		switch v := data.(type) {

		case *float32:
			float32ToByte(bs, float32(*v))
		case float32:
			float32ToByte(bs, float32(v))
		case []float32:
			for i, x := range v {
				float32ToByte(bs[4*i:], float32(x))
			}
		case *float64:
			float64ToByte(bs, float64(*v))
		case float64:
			float64ToByte(bs, float64(v))
		case []float64:
			for i, x := range v {
				float64ToByte(bs[8*i:], float64(x))
			}
		}
		_, err := w.Write(bs)
		return err
	}

	return binary.Write(w, order, data)

}

// A internal read to get support for floats
func read(r io.Reader, order binary.ByteOrder, data interface{}) error {
	if n := intDataSize(data); n != 0 {
		var b [8]byte
		var bs []byte
		if n > len(b) {
			bs = make([]byte, n)
		} else {
			bs = b[:n]
		}
		if _, err := io.ReadFull(r, bs); err != nil {
			return err
		}
		switch data := data.(type) {
		case *float32:
			*data = bytesToFloat32(bs)
		case *float64:
			fmt.Printf("\nfloat64 %+v", data)
			*data = bytesToFloat64(bs)
			fmt.Printf("\nfloat64 %+v", data)
		case float64:
			fmt.Printf("\nfloat64 %+v", data)
			data = bytesToFloat64(bs)
			fmt.Printf("\nfloat64 %+v", data)
		case []float32:
			for i := range data {
				data[i] = bytesToFloat32(bs[4*i:])
			}
		case []float64:
			for i := range data {
				data[i] = bytesToFloat64(bs[8*i:])
			}
		}
		return nil
	}

	return binary.Read(r, order, data)
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
