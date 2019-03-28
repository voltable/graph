package operators

import (
	"testing"

	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore/storage/memorydb"
)

type unaryTest struct {
}

func (s *unaryTest) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
	return i
}

func (s *unaryTest) Op() {

}

func TestFilter_Next(t *testing.T) {
	type fields struct {
		storage  widecolumnstore.Storage
		operator widecolumnstore.Operator
		prefix   widecolumnstore.Prefix
	}
	tests := []struct {
		name   string
		fields fields
		args   func([]widecolumnstore.KeyValue) widecolumnstore.Iterator
		want   []widecolumnstore.KeyValue
	}{
		{
			name: "Filter",
			fields: func() fields {
				storage, _ := memorydb.NewStorageEngine()
				unary := &unaryTest{}
				fields := fields{
					storage:  storage,
					operator: unary,
					prefix: func(widecolumnstore.Key) []byte {
						arr := []byte{}
						return arr
					},
				}
				return fields
			}(),
			args: func(array []widecolumnstore.KeyValue) widecolumnstore.Iterator {
				l := len(array)
				i := 0
				return func() (widecolumnstore.KeyValue, bool) {
					if i < l {
						old := i
						i++
						return array[old], true
					}

					return widecolumnstore.KeyValue{}, false
				}
			},
			want: func() []widecolumnstore.KeyValue {
				want := []widecolumnstore.KeyValue{}
				want = append(want, widecolumnstore.KeyValue{})
				return want
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewFilter(tt.fields.storage,
				tt.fields.operator,
				tt.fields.prefix,
			)
			got := s.Next(tt.args(tt.want))
			for value, ok := got(); ok; value, ok = got() {
				t.Errorf("Filter.Next() = %v", value)
			}
			// if !reflect.DeepEqual(got(), tt.want()) {
			// 	t.Errorf("Filter.Next() = %v, want %v", got, tt.want)
			// }
		})
	}
}
