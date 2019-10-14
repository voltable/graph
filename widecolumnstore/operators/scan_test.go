package operators_test

import (
	"reflect"
	"testing"

	"github.com/voltable/graph/widecolumnstore"
	"github.com/voltable/graph/widecolumnstore/operators"
	"github.com/voltable/graph/widecolumnstore/storage/memorydb"
)

func TestScan_Next(t *testing.T) {
	type fields struct {
		storage widecolumnstore.Storage
		prefix  []byte
	}

	tests := []struct {
		name   string
		fields fields
		want   []*widecolumnstore.KeyValue
	}{
		{
			name: "Each",
			fields: func() fields {
				storage, _ := memorydb.NewStorageEngine()
				fields := fields{
					storage: storage,
					prefix:  nil,
				}
				return fields
			}(),
			want: func() []*widecolumnstore.KeyValue {
				want := []*widecolumnstore.KeyValue{}
				want = append(want, &widecolumnstore.KeyValue{})
				return want
			}(),
		},
		{
			name: "Prefix",
			fields: func() fields {
				storage, _ := memorydb.NewStorageEngine()
				fields := fields{
					storage: storage,
					prefix:  []byte{},
				}
				return fields
			}(),
			want: func() []*widecolumnstore.KeyValue {
				want := []*widecolumnstore.KeyValue{}
				want = append(want, &widecolumnstore.KeyValue{})
				return want
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.fields.storage.Create(tt.want...)
			s := operators.NewScan(
				tt.fields.storage,
				tt.fields.prefix,
			)
			iterator := s.Next()

			got := []*widecolumnstore.KeyValue{}
			for kv, ok := iterator(); ok; kv, ok = iterator() {
				got = append(got, &kv)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%d. %q: Next() = got %v, want %v", i, tt.name, got, tt.want)
			}
		})
	}
}
