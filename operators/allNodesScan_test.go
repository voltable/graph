package operators

import (
	"reflect"
	"testing"

	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/widecolumnstore"
)

func TestAllNodesScan_Next(t *testing.T) {
	type fields struct {
		storage    widecolumnstore.Storage
		statistics *graph.Statistics
		variable   ir.Variable
	}
	tests := []struct {
		name    string
		fields  fields
		want    widecolumnstore.Iterator
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AllNodesScan{
				storage:    tt.fields.storage,
				statistics: tt.fields.statistics,
				variable:   tt.fields.variable,
			}
			got, err := s.Next()
			if (err != nil) != tt.wantErr {
				t.Errorf("AllNodesScan.Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllNodesScan.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
