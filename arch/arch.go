// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package arch

import (
	"encoding/binary"
	"fmt"
	"math"
)

// EncodeStringBytes encodes string into bytes
func EncodeStringBytes(v string) []byte {
	return []byte(v)
}

// DecodeStringBytes decodes bytes into a string
func DecodeStringBytes(buf []byte) interface{} {
	return String(buf)
}

// String decodes bytes into a string
func String(buf []byte) string {
	return string(buf)
}

// EncodeBoolBytes encodes bool into bytes
func EncodeBoolBytes(v bool) []byte {
	buf := make([]byte, 1)
	if v {
		buf[0] = 1
	} else {
		buf[0] = 0
	}
	return buf
}

// DecodeBoolBytes decodes bytes into a bool
func DecodeBoolBytes(buf []byte) interface{} {
	return Bool(buf)
}

// Bool decodes bytes into a bool
func Bool(buf []byte) bool {
	return buf[0] != 0
}

// EncodeInt16Bytes encodes int16 into bytes
func EncodeInt16Bytes(v int16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(v))
	return buf
}

// DecodeInt16Bytes decodes bytes into a int16
func DecodeInt16Bytes(buf []byte) interface{} {
	return Int16(buf)
}

// Int16 decodes bytes into a int16
func Int16(buf []byte) int16 {
	return int16(binary.LittleEndian.Uint16(buf))
}

// EncodeInt32Bytes encodes int32 into bytes
func EncodeInt32Bytes(v int32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(v))
	return buf
}

// DecodeInt32Bytes decodes bytes into a int32
func DecodeInt32Bytes(buf []byte) interface{} {
	return Int32(buf)
}

// Int32 decodes bytes into a int32
func Int32(buf []byte) int32 {
	return int32(binary.LittleEndian.Uint32(buf))
}

// EncodeInt64Bytes encodes int64 into bytes
func EncodeInt64Bytes(v int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(v))
	return buf
}

// DecodeInt64Bytes decodes bytes into a int64
func DecodeInt64Bytes(buf []byte) interface{} {
	return Int64(buf)
}

// Int64 decodes bytes into a int64
func Int64(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}

// EncodeUint16Bytes encodes uint16 into bytes
func EncodeUint16Bytes(v uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, v)
	return buf
}

// DecodeUint16Bytes decodes bytes into a uint16
func DecodeUint16Bytes(buf []byte) interface{} {
	return Uint16(buf)
}

// Uint16 decodes bytes into a uint16
func Uint16(buf []byte) uint16 {
	return binary.LittleEndian.Uint16(buf)
}

// EncodeUint32Bytes encodes uint32 into bytes
func EncodeUint32Bytes(v uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, v)
	return buf
}

// DecodeUint32Bytes decodes bytes into a uint32
func DecodeUint32Bytes(buf []byte) interface{} {
	return Uint32(buf)
}

// Uint32 decodes bytes into a uint32
func Uint32(buf []byte) uint32 {
	return binary.LittleEndian.Uint32(buf)
}

// EncodeUint64Bytes encodes uint64 into bytes
func EncodeUint64Bytes(v uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, v)
	return buf
}

// DecodeUint64Bytes decodes bytes into a uint64
func DecodeUint64Bytes(buf []byte) interface{} {
	return Uint64(buf)
}

// Uint64 decodes bytes into a uint64
func Uint64(buf []byte) uint64 {
	return binary.LittleEndian.Uint64(buf)
}

// EncodeFloat32Bytes encodes float32 into bytes
func EncodeFloat32Bytes(v float32) []byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(v))
	return buf[:]
}

// DecodeFloat32Bytes decodes bytes into a float32
func DecodeFloat32Bytes(buf []byte) interface{} {
	return Float32(buf)
}

// Float32 decodes bytes into a float32
func Float32(buf []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(buf))
}

// EncodeFloat64Bytes encode float64 into bytes
func EncodeFloat64Bytes(v float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(v))
	return buf[:]
}

// DecodeFloat64Bytes decodes bytes into a float64
func DecodeFloat64Bytes(buf []byte) interface{} {
	return Float64(buf)
}

// Float64 decodes bytes into a float64
func Float64(buf []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(buf))
}

// EncodeComplex64Bytes encodes a complex64 into bytes
func EncodeComplex64Bytes(v complex64) []byte {
	rpart := float32(real(v))
	ipart := float32(imag(v))

	var buf [8]byte

	binary.LittleEndian.PutUint32(buf[0:4], math.Float32bits(rpart))
	binary.LittleEndian.PutUint32(buf[4:8], math.Float32bits(ipart))

	return buf[:]
}

// DecodeComplex64Bytes decodes bytes into a complex64
func DecodeComplex64Bytes(buf []byte) interface{} {
	return Complex64(buf)
}

// Complex64 decodes bytes into a complex64
func Complex64(buf []byte) complex64 {
	return complex(Float32(buf[0:4]), Float32(buf[4:8]))
}

// EncodeComplex128Bytes encodes a complex128 into bytes
func EncodeComplex128Bytes(v complex128) []byte {
	rpart := float64(real(v))
	ipart := float64(imag(v))

	var buf [16]byte

	binary.LittleEndian.PutUint64(buf[0:8], math.Float64bits(rpart))
	binary.LittleEndian.PutUint64(buf[8:16], math.Float64bits(ipart))

	return buf[:]
}

// DecodeComplex128Bytes decodes bytes into a complex128
func DecodeComplex128Bytes(buf []byte) interface{} {
	return Complex128(buf)
}

// Complex128 decodes bytes into a complex128
func Complex128(buf []byte) complex128 {
	return complex(Float64(buf[0:8]), Float64(buf[8:16]))
}

// EncodeType takes a interface and returns its bytes
func EncodeType(i interface{}) (string, []byte) {
	switch v := i.(type) {
	case bool:
		return "bool", EncodeBoolBytes(v)
	case string:
		return "string", EncodeStringBytes(v)
	case int16:
		return "int16", EncodeInt16Bytes(v)
	case int32:
		return "int32", EncodeInt32Bytes(v)
	case int:
		return "int", EncodeInt32Bytes(int32(v))
	case int64:
		return "int64", EncodeInt64Bytes(v)
	case uint16:
		return "uint16", EncodeUint16Bytes(v)
	case uint32:
		return "uint32", EncodeUint32Bytes(v)
	case uint:
		return "uint", EncodeUint32Bytes(uint32(v))
	case uint64:
		return "uint64", EncodeUint64Bytes(v)
	case float32:
		return "float32", EncodeFloat32Bytes(v)
	case float64:
		return "float64", EncodeFloat64Bytes(v)
	case complex64:
		return "complex64", EncodeComplex64Bytes(v)
	case complex128:
		return "complex128", EncodeComplex128Bytes(v)
	case []byte:
		return "[]byte", v
	}

	return fmt.Sprintf("%T", i), []byte(fmt.Sprint(i))
}

// DecodeType takes a interface and returns its bytes
func DecodeType(t string, buf []byte) interface{} {
	switch t {
	case "bool":
		return DecodeBoolBytes(buf)
	case "string":
		return DecodeStringBytes(buf)
	case "int16":
		return DecodeInt16Bytes(buf)
	case "int32":
		return DecodeInt32Bytes(buf)
	case "int":
		return int(Uint32(buf))
	case "int64":
		return DecodeInt64Bytes(buf)
	case "uint16":
		return DecodeUint16Bytes(buf)
	case "uint32":
		return DecodeUint32Bytes(buf)
	case "uint":
		return uint(Uint32(buf))
	case "uint64":
		return DecodeUint64Bytes(buf)
	case "float32":
		return DecodeFloat32Bytes(buf)
	case "float64":
		return DecodeFloat64Bytes(buf)
	case "complex64":
		return DecodeComplex64Bytes(buf)
	case "complex128":
		return DecodeComplex128Bytes(buf)
	case "[]byte":
		return buf
	}

	return string(buf)
}
