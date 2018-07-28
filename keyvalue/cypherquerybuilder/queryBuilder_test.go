package cypherquerybuilder_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/keyvalue/cypherquerybuilder"
	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/query/cypher/ast"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

type FakeStorage struct {
	tKeyIndex map[int][]byte
	tKey      map[string]*keyvalue.Any
}

func (s FakeStorage) Fetch(string) (*keyvalue.KeyValue, error) {
	return nil, nil
}

func (s FakeStorage) ForEach() query.IteratorUUID {
	return func() (uuid.UUID, bool) {
		return uuid.UUID{}, false
	}
}

func (s FakeStorage) HasPrefix(prefix []byte) query.Iterator {
	position := 0
	length := len(s.tKey)
	return func() (interface{}, bool) {
		for position < length {
			key := s.tKeyIndex[position]
			position = position + 1

			if bytes.HasPrefix(key, prefix) {
				v := s.tKey[string(key)]
				kv := &keyvalue.KeyValue{Key: key, Value: v}
				return kv, true
			}
		}

		return nil, false
	}
}

func (s FakeStorage) Edges(uuid.UUID) query.IteratorUUID {
	return func() (uuid.UUID, bool) {
		return uuid.UUID{}, false
	}
}

func (s FakeStorage) HasPrefixRange([][]byte) query.Iterator {
	return func() (interface{}, bool) {
		return nil, false
	}
}

func NewFakeStorage(triples ...*keyvalue.KeyValue) query.Storage {
	s := &FakeStorage{
		tKeyIndex: make(map[int][]byte),
		tKey:      make(map[string]*keyvalue.Any),
	}

	for i := 0; i < len(triples); i++ {
		triple := triples[i]
		s.tKeyIndex[len(s.tKey)] = triple.Key
		s.tKey[string(triple.Key)] = triple.Value
	}

	return s
}

func TestKeyValueCyperQueryBuilder_ToPredicateVertexPath(t *testing.T) {
	tests := []struct {
		name         string
		storage      func(uuid.UUID) query.Storage
		patn         *ast.VertexPatn
		id           uuid.UUID
		wantVariable string
		wantTraverse query.Traverse
		wantWeight   float64
	}{
		{
			name:         "Empty",
			wantVariable: "",
			wantTraverse: query.Failed,
			wantWeight:   0,
			patn: func() *ast.VertexPatn {
				patn := &ast.VertexPatn{}
				return patn
			}(),
			id: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
			storage: func(id uuid.UUID) query.Storage {
				kv := keyvalue.NewKeyValueProperty(id, "", "")
				return NewFakeStorage(kv)
			},
		},
		{
			name:         "Name",
			wantVariable: "",
			wantTraverse: query.Matched,
			wantWeight:   0,
			patn: func() *ast.VertexPatn {
				patn := &ast.VertexPatn{
					Properties: map[string]interface{}{"name": "John Smith"},
				}
				return patn
			}(),
			id: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
			storage: func(id uuid.UUID) query.Storage {
				kv := keyvalue.NewKeyValueProperty(id, "name", "John Smith")
				return NewFakeStorage(kv)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := cypherquerybuilder.NewKeyValueCyperQueryBuilder(tt.storage(tt.id))
			predicate, err := s.ToPredicateVertexPath(tt.patn)
			if err != nil {
				t.Error(err)
				return
			}

			r1, r2, r3 := predicate(&tt.id, nil, 0)
			if !reflect.DeepEqual(r1, tt.wantVariable) {
				t.Errorf("KeyValueCyperQueryBuilder.ToPredicateVertexPath() = %v, want %v", r1, tt.wantVariable)
			}

			if !reflect.DeepEqual(r2, tt.wantTraverse) {
				t.Errorf("KeyValueCyperQueryBuilder.ToPredicateVertexPath() = %v, want %v", r2, tt.wantTraverse)
			}

			if !reflect.DeepEqual(r3, tt.wantWeight) {
				t.Errorf("KeyValueCyperQueryBuilder.ToPredicateVertexPath() = %v, want %v", r3, tt.wantWeight)
			}
		})
	}
}

func TestKeyValueCyperQueryBuilder_ToPredicateEdgePath(t *testing.T) {
	tests := []struct {
		name         string
		storage      func(*uuid.UUID, *uuid.UUID) query.Storage
		patn         *ast.EdgePatn
		from         *uuid.UUID
		to           *uuid.UUID
		wantVariable string
		wantTraverse query.Traverse
		wantWeight   float64
	}{
		{
			name:         "Empty",
			wantVariable: "",
			wantTraverse: query.Matching,
			wantWeight:   0,
			patn: func() *ast.EdgePatn {
				patn := &ast.EdgePatn{}
				return patn
			}(),
			from: func() *uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return &id
			}(),
			to: func() *uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return &id
			}(),
			storage: func(from, to *uuid.UUID) query.Storage {
				kv := keyvalue.NewKeyValueRelationshipProperty(*from, *to, "", "")
				return NewFakeStorage(kv)
			},
		},
		{
			name:         "Name",
			wantVariable: "",
			wantTraverse: query.Visiting,
			wantWeight:   0,
			patn: func() *ast.EdgePatn {
				patn := &ast.EdgePatn{
					Body: &ast.EdgeBodyStmt{
						Properties: map[string]interface{}{"name": "John Smith"},
					},
				}
				return patn
			}(),
			from: func() *uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return &id
			}(),
			to: func() *uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return &id
			}(),
			storage: func(from, to *uuid.UUID) query.Storage {
				kv := keyvalue.NewKeyValueRelationshipProperty(*from, *to, "name", "John Smith")
				return NewFakeStorage(kv)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := cypherquerybuilder.NewKeyValueCyperQueryBuilder(tt.storage(tt.from, tt.to))
			predicate, err := s.ToPredicateEdgePath(tt.patn)
			if err != nil {
				t.Error(err)
				return
			}

			r1, r2, r3 := predicate(tt.from, tt.to, 0)
			if !reflect.DeepEqual(r1, tt.wantVariable) {
				t.Errorf("KeyValueCyperQueryBuilder.ToPredicateEdgePath() = %v, want %v", r1, tt.wantVariable)
			}

			if !reflect.DeepEqual(r2, tt.wantTraverse) {
				t.Errorf("KeyValueCyperQueryBuilder.ToPredicateEdgePath() = %v, want %v", r2, tt.wantTraverse)
			}

			if !reflect.DeepEqual(r3, tt.wantWeight) {
				t.Errorf("KeyValueCyperQueryBuilder.ToPredicateEdgePath() = %v, want %v", r3, tt.wantWeight)
			}
		})
	}
}
