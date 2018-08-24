package cypher_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

func TestQueryBuilder_ToPredicateVertexPath(t *testing.T) {
	type fields struct {
		storage widecolumnstore.Storage
	}
	type args struct {
		patn *ast.VertexPatn
		last widecolumnstore.Operator
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    widecolumnstore.Operator
		wantErr bool
	}{
		{
			args: args{
				patn: &ast.VertexPatn{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := cypher.NewQueryBuilder(tt.fields.storage)
			got, err := s.ToPredicateVertexPath(tt.args.patn, tt.args.last)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryBuilder.ToPredicateVertexPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBuilder.ToPredicateVertexPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
