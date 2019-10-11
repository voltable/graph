package cypher_test

import (
	"reflect"
	"testing"

	"github.com/voltable/graph/query"
	"github.com/voltable/graph/query/cypher"
	"github.com/voltable/graph/query/cypher/ast"
	"github.com/voltable/graph/widecolumnstore"
	"github.com/voltable/graph/widecolumnstore/storage/memorydb"
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
	tests := []struct {
		name    string
		storage widecolumnstore.Storage
		filter  func(storage widecolumnstore.HasPrefix, operator widecolumnstore.Operator, prefix widecolumnstore.Prefix) (widecolumnstore.Unary, error)
		patn    *ast.VertexPatn
		last    widecolumnstore.Operator
		want    widecolumnstore.Operator
		err     error
	}{
		{
			name: "No Pattern",
			err:  cypher.ErrNoPattern,
		},
		{
			name: "No Operator",
			patn: &ast.VertexPatn{
				Properties: func() map[string]interface{} {
					prop := make(map[string]interface{}, 0)
					prop["key"] = "value"
					return prop
				}(),
			},
			err: cypher.ErrNoLastOperator,
		},
		{
			name: "Properties filter pattern",
			filter: func(h widecolumnstore.HasPrefix, o widecolumnstore.Operator, p widecolumnstore.Prefix) (widecolumnstore.Unary, error) {
				key := widecolumnstore.Key{
					ID: []byte("id"),
				}
				bytes := p(key)
				return &filterMock{
					bytes: bytes,
				}, nil
			},
			storage: func() widecolumnstore.Storage {
				storage, _ := memorydb.NewStorageEngine()
				return storage
			}(),
			patn: &ast.VertexPatn{
				Properties: func() map[string]interface{} {
					prop := make(map[string]interface{}, 0)
					prop["key"] = "value"
					return prop
				}(),
			},
			last: &unaryMock{},
			want: &filterMock{widecolumnstore.NewKey(query.TProperties, &widecolumnstore.Column{[]byte("key"), nil, []byte("id")}).Marshal()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := cypher.NewQueryBuilder(tt.storage, tt.filter)
			got, err := s.ToPredicateVertexPath(tt.patn, tt.last)

			if err != tt.err {
				t.Errorf("QueryBuilder.ToPredicateVertexPath() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBuilder.ToPredicateVertexPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryBuilder_ToPredicateEdgePath(t *testing.T) {

	tests := []struct {
		name    string
		storage widecolumnstore.Storage
		filter  func(storage widecolumnstore.HasPrefix, operator widecolumnstore.Operator, prefix widecolumnstore.Prefix) (widecolumnstore.Unary, error)
		patn    *ast.EdgePatn
		last    widecolumnstore.Operator
		want    widecolumnstore.Operator
		err     error
	}{
		{
			name: "No Pattern",
			err:  cypher.ErrNoPattern,
		},
		{
			name: "No Operator",
			patn: &ast.EdgePatn{},
			err:  cypher.ErrNoLastOperator,
		},
		{
			name: "Properties filter pattern",
			filter: func(h widecolumnstore.HasPrefix, o widecolumnstore.Operator, p widecolumnstore.Prefix) (widecolumnstore.Unary, error) {
				key := widecolumnstore.Key{
					ID: []byte("id"),
				}
				bytes := p(key)
				return &filterMock{
					bytes: bytes,
				}, nil
			},
			storage: func() widecolumnstore.Storage {
				storage, _ := memorydb.NewStorageEngine()
				return storage
			}(),
			patn: &ast.EdgePatn{
				Body: &ast.EdgeBodyStmt{
					Properties: func() map[string]interface{} {
						prop := make(map[string]interface{}, 0)
						prop["key"] = "value"
						return prop
					}(),
				},
			},
			last: &unaryMock{},
			want: &filterMock{widecolumnstore.NewKey([]byte("id"), &widecolumnstore.Column{query.Relationshipproperties, []byte("key"), nil}).Marshal()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := cypher.NewQueryBuilder(tt.storage, tt.filter)

			got, err := s.ToPredicateEdgePath(tt.patn, tt.last)
			if err != tt.err {
				t.Errorf("QueryBuilder.ToPredicateEdgePath() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBuilder.ToPredicateEdgePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryBuilder_Predicate(t *testing.T) {
	tests := []struct {
		name     string
		storage  widecolumnstore.Storage
		filter   func(storage widecolumnstore.HasPrefix, operator widecolumnstore.Operator, prefix widecolumnstore.Prefix) (widecolumnstore.Unary, error)
		patterns []ast.Patn
		want     widecolumnstore.Operator
		err      error
	}{
		{
			name: "No Pattern",
			err:  cypher.ErrNoPattern,
		},
		{
			name: "Edge Pattern",
			patterns: func() []ast.Patn {
				patterns := make([]ast.Patn, 0)
				edge := &ast.EdgePatn{
					Body: &ast.EdgeBodyStmt{
						Properties: func() map[string]interface{} {
							prop := make(map[string]interface{}, 0)
							prop["key"] = "value"
							return prop
						}(),
					},
				}
				patterns = append(patterns, edge)
				return patterns
			}(),
			filter: func(h widecolumnstore.HasPrefix, o widecolumnstore.Operator, p widecolumnstore.Prefix) (widecolumnstore.Unary, error) {
				key := widecolumnstore.Key{
					ID: []byte("id"),
				}
				bytes := p(key)
				return &filterMock{
					bytes: bytes,
				}, nil
			},
			storage: func() widecolumnstore.Storage {
				storage, _ := memorydb.NewStorageEngine()
				return storage
			}(),
			want: &filterMock{widecolumnstore.NewKey([]byte("id"), &widecolumnstore.Column{query.Relationshipproperties, []byte("key"), nil}).Marshal()},
		},

		{
			name: "Vertex Pattern",
			patterns: func() []ast.Patn {
				patterns := make([]ast.Patn, 0)
				vertex := &ast.VertexPatn{
					Properties: func() map[string]interface{} {
						prop := make(map[string]interface{}, 0)
						prop["key"] = "value"
						return prop
					}(),
				}
				patterns = append(patterns, vertex)
				return patterns
			}(),
			filter: func(h widecolumnstore.HasPrefix, o widecolumnstore.Operator, p widecolumnstore.Prefix) (widecolumnstore.Unary, error) {
				key := widecolumnstore.Key{
					ID: []byte("id"),
				}
				bytes := p(key)
				return &filterMock{
					bytes: bytes,
				}, nil
			},
			storage: func() widecolumnstore.Storage {
				storage, _ := memorydb.NewStorageEngine()
				return storage
			}(),
			want: &filterMock{widecolumnstore.NewKey(query.TProperties, &widecolumnstore.Column{[]byte("key"), nil, []byte("id")}).Marshal()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := cypher.NewQueryBuilder(tt.storage, tt.filter)

			got, err := s.Predicate(tt.patterns)
			if err != tt.err {
				t.Errorf("QueryBuilder.Predicate() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryBuilder.Predicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
