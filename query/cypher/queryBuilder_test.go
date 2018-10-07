package cypher_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore/operators"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore/storage/memorydb"
)

type unaryTest struct {
}

func (s *unaryTest) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
	return i
}

func (s *unaryTest) Op() {

}

func TestQueryBuilder_ToPredicateVertexPath(t *testing.T) {
	storage, _ := memorydb.NewStorageEngine()
	last := &unaryTest{}

	prefix := func(kv widecolumnstore.KeyValue) []byte {
		key := widecolumnstore.Key{}
		key.Unmarshal(kv.Key)
		return widecolumnstore.NewKey(query.TProperties, &widecolumnstore.Column{[]byte("test"), nil, []byte("test")}).Marshal()

	}

	want := operators.NewFilter(storage, last, prefix)
	patn := &ast.VertexPatn{
		Properties: func() map[string]interface{} {
			prop := make(map[string]interface{}, 0)
			prop["test"] = "test"
			return prop
		}(),
	}

	newFilter := func(widecolumnstore.HasPrefix, widecolumnstore.Operator, widecolumnstore.Prefix) *operators.Filter {
		return want
	}

	t.Run("", func(t *testing.T) {
		s := cypher.NewQueryBuilder(storage, newFilter)
		got, err := s.ToPredicateVertexPath(patn, last)
		if err != nil {
			t.Errorf("QueryBuilder.ToPredicateVertexPath() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("QueryBuilder.ToPredicateVertexPath() = got \n%#v, want \n%#v", got, want)
		}
	})
}
