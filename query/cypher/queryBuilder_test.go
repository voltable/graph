package cypher_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore/storage/memorydb"
)

type unaryMock struct {
}

func (s *unaryMock) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
	return i
}

func (s *unaryMock) Op() {

}

type filterMock struct {
	bytes []byte
}

func (s *filterMock) Op() {}

func (s *filterMock) Next(i widecolumnstore.Iterator) widecolumnstore.Iterator {
	return i
}

func TestQueryBuilder_ToPredicateVertexPath(t *testing.T) {
	storage, _ := memorydb.NewStorageEngine()
	last := &unaryMock{}

	want := &filterMock{widecolumnstore.NewKey(query.TProperties, &widecolumnstore.Column{[]byte("key"), nil, []byte("id")}).Marshal()}

	patn := &ast.VertexPatn{
		Properties: func() map[string]interface{} {
			prop := make(map[string]interface{}, 0)
			prop["key"] = "value"
			return prop
		}(),
	}

	newFilter := func(h widecolumnstore.HasPrefix, o widecolumnstore.Operator, p widecolumnstore.Prefix) widecolumnstore.Unary {
		bytes := p(widecolumnstore.KeyValue{Key: []byte("id")})
		return &filterMock{
			bytes: bytes,
		}
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
