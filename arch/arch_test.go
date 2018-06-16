// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package arch_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/arch"
)

func TestEncodeType(t *testing.T) {
	tests := []struct {
		name string
		arg  interface{}
	}{
		{
			name: "string",
			arg:  "string",
		},
		{
			name: "bool",
			arg:  true,
		},
		{
			name: "uint16",
			arg:  uint16(1),
		},
		{
			name: "uint32",
			arg:  uint32(1),
		},
		{
			name: "uint64",
			arg:  uint64(1),
		},
		{
			name: "int16",
			arg:  int16(1),
		},
		{
			name: "int32",
			arg:  int32(1),
		},
		{
			name: "int64",
			arg:  int64(1),
		},
		{
			name: "float32",
			arg:  float32(1),
		},
		{
			name: "float64",
			arg:  float64(1),
		},
		{
			name: "complex64",
			arg:  complex64(1),
		},
		{
			name: "complex128",
			arg:  complex128(1),
		},
		{
			name: "[]byte",
			arg:  []byte("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := arch.EncodeType(tt.arg)
			if got != tt.name {
				t.Errorf("EncodeType() got = %v, want %v", got, tt.name)
			}
			bytes := arch.DecodeType(got, got1)
			if !reflect.DeepEqual(bytes, tt.arg) {
				t.Errorf("DecodeType() interface = %v, want %v", bytes, tt.arg)
			}
		})
	}
}
