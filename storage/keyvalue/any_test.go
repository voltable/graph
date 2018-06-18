// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package keyvalue_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/storage/keyvalue"
)

func TestAny_Unmarshal(t *testing.T) {

	tests := []struct {
		name string
		want interface{}
	}{
		{
			name: "string",
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := keyvalue.NewAny(tt.want)
			if got := s.Unmarshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Any.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
