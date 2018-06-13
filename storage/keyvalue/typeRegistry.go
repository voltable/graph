// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package keyvalue

import (
	"log"
	"reflect"
	"strconv"
)

var typeRegistry = make(map[string]reflect.Type)
var typeDecoder = make(map[string]func(buf []byte) interface{})

func init() {
	RegisterType("", "string", DecodeStringBytes)
	RegisterType(true, "bool", DecodeBoolBytes)
	RegisterType(int8(0), "int8", DecodeInt8Bytes)
	RegisterType(int16(0), "int16", DecodeInt16Bytes)

	// typeRegistry[String.String()] = reflect.TypeOf("")
	// typeRegistry[Bool] = reflect.TypeOf(false)
	// typeRegistry[Int8] = reflect.TypeOf(int8(1))
	// typeRegistry[Int16] = reflect.TypeOf(int16(1))
	// typeRegistry[Int32] = reflect.TypeOf(int32(1))
	// typeRegistry[Int64] = reflect.TypeOf(int64(1))
	// typeRegistry[Uint8] = reflect.TypeOf(uint8(1))
	// typeRegistry[Uint16] = reflect.TypeOf(uint16(1))
	// typeRegistry[Uint32] = reflect.TypeOf(uint32(1))
	// typeRegistry[Uint64] = reflect.TypeOf(uint64(1))
	// typeRegistry[Float32] = reflect.TypeOf(float32(1))
	// typeRegistry[Float64] = reflect.TypeOf(float64(1))
	// typeRegistry[Complex64] = reflect.TypeOf(complex64(1))
	// typeRegistry[Complex128] = reflect.TypeOf(complex128(1))
	// typeRegistry[Byte] = reflect.TypeOf(byte(1))
}

// RegisterType is called from generated code and maps from the fully qualified
func RegisterType(x interface{}, name string, f func(buf []byte) interface{}) {
	if _, ok := typeRegistry[name]; ok {
		log.Printf("keyvalue: duplicate type registered: %s", name)
		return
	}
	t := reflect.TypeOf(x)
	typeRegistry[name] = t
	typeDecoder[name] = f
}

// DecodeBytes returns the decode buf as it's type
func DecodeBytes(name string, buf []byte) interface{} {
	if f, ok := typeDecoder[name]; ok {
		return f(buf)
	}

	log.Printf("keyvalue: type not registered: %s", name)
	return nil
}

// DecodeStringBytes decodes bytes into a string
func DecodeStringBytes(buf []byte) interface{} {
	return string(buf)
}

// DecodeBoolBytes decodes bytes into a bool
func DecodeBoolBytes(buf []byte) interface{} {
	return uint64(buf[0]) != 0
}

// DecodeInt8Bytes decodes bytes into a int8
func DecodeInt8Bytes(buf []byte) interface{} {
	i, _ := strconv.ParseInt(string(buf), 10, 8)
	return int8(i)
}

// DecodeInt16Bytes decodes bytes into a int16
func DecodeInt16Bytes(buf []byte) interface{} {
	i, _ := strconv.ParseInt(string(buf), 10, 16)
	return int16(i)
}
