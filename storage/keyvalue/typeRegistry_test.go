// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package keyvalue

import (
	"fmt"
	"reflect"
	"testing"
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
			if got := DecodeStringBytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
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
			if got := DecodeBoolBytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeBoolBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeInt8Bytes(t *testing.T) {

	tests := []struct {
		name string
		buf  []byte
		want int8
	}{
		{
			name: "int8",
			buf:  []byte(fmt.Sprint(int8(1))),
			want: int8(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeInt8Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInt8Bytes() = %v, want %v", got, tt.want)
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
			buf:  []byte(fmt.Sprint(int16(1))),
			want: int16(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeInt16Bytes(tt.buf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeInt16Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
