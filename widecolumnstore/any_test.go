package widecolumnstore_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
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
			s := widecolumnstore.NewAny(tt.want)
			if got := widecolumnstore.Unmarshal(s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Any.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
