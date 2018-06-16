// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package arch_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/arch"
)

func TestDecodeStringBytes(t *testing.T) {
	tests := []struct {
		name string
		buf  []byte
		want string
	}{
		{
			name: "string",
			buf:  []byte("1"),
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeStringBytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeStringBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeBoolBytes(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		buf  []byte
		want bool
	}{
		{
			name: "bool",
			buf:  []byte(fmt.Sprint(true)),
			want: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeBoolBytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeBoolBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeInt16Bytes(t *testing.T) {

	tests := []struct {
		name string
		buf  []byte
		want int16
	}{
		{
			name: "int16",
			buf:  arch.EncodeInt16Bytes(int16(1)),
			want: int16(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeInt16Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInt16Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeInt32Bytes(t *testing.T) {

	tests := []struct {
		name string
		buf  []byte
		want int32
	}{
		{
			name: "int32",
			buf:  arch.EncodeInt32Bytes(int32(1)),
			want: int32(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeInt32Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInt32Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeInt64Bytes(t *testing.T) {
	tests := []struct {
		name string
		buf  []byte
		want int64
	}{
		{
			name: "int64",
			buf:  arch.EncodeInt64Bytes(int64(1)),
			want: int64(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeInt64Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInt64Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeUint16Bytes(t *testing.T) {

	tests := []struct {
		name string
		buf  []byte
		want uint16
	}{
		{
			name: "uint16",
			buf:  arch.EncodeUint16Bytes(uint16(1)),
			want: uint16(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeUint16Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeUint16Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeUint32Bytes(t *testing.T) {

	tests := []struct {
		name string
		buf  []byte
		want uint32
	}{
		{
			name: "uint32",
			buf:  arch.EncodeUint32Bytes(uint32(1)),
			want: uint32(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeUint32Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeUint32Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeUint64Bytes(t *testing.T) {
	tests := []struct {
		name string
		buf  []byte
		want uint64
	}{
		{
			name: "uint64",
			buf:  arch.EncodeUint64Bytes(uint64(1)),
			want: uint64(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeUint64Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeUint64Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeFloat32Bytes(t *testing.T) {
	tests := []struct {
		name string
		buf  []byte
		want float32
	}{
		{
			name: "float32",
			buf:  arch.EncodeFloat32Bytes(float32(1)),
			want: float32(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeFloat32Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeFloat32Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeFloat64Bytes(t *testing.T) {
	tests := []struct {
		name string
		buf  []byte
		want float64
	}{
		{
			name: "float64",
			buf:  arch.EncodeFloat64Bytes(float64(12.666512)),
			want: float64(12.666512),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arch.DecodeFloat64Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeFloat64Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
