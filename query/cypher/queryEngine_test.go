package cypher_test

import (
	"reflect"
	"testing"

	graph "github.com/voltable/graph"
	"github.com/voltable/graph/query/cypher"
)

func TestQueryEngine_Parse(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		qe      cypher.QueryEngine
		args    args
		want    *graph.Query
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.qe.Parse(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryEngine.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryEngine.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
