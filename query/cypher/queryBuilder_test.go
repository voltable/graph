package cypher_test

import (
	"reflect"
	"testing"

	"github.com/voltable/graph/query/cypher"
	"github.com/voltable/graph/query/cypher/ast"
	"github.com/voltable/graph/widecolumnstore"
)

func TestCypherQueryBuilder_Build(t *testing.T) {
	type args struct {
		stmt ast.Clauses
	}
	tests := []struct {
		name    string
		s       *cypher.CypherQueryBuilder
		args    args
		want    widecolumnstore.Iterator
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Match",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Build(tt.args.stmt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CypherQueryBuilder.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CypherQueryBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
