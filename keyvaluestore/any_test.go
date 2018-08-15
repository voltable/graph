package keyvaluestore_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/keyvaluestore"
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
			s := keyvaluestore.NewAny(tt.want)
			if got := keyvaluestore.Unmarshal(s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Any.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
