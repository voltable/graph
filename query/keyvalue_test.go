package query_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/query"
	"github.com/RossMerr/Caudex.Graph/uuid"
	"github.com/RossMerr/Caudex.Graph/widecolumnstore"
)

func TestUUID(t *testing.T) {
	tests := []struct {
		name  string
		setup func(id uuid.UUID) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue)
		want  uuid.UUID
	}{
		{
			name: "Vertex",
			setup: func(id uuid.UUID) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
				return query.NewKeyValueVertex(id, "person")
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Properties",
			setup: func(id uuid.UUID) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
				return query.NewKeyValueProperty(id, "", "")
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Relationship",
			setup: func(id uuid.UUID) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
				to, _ := uuid.GenerateRandomUUID()
				return query.NewKeyValueRelationship(id, to, "", 5)
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Relationshipproperties",
			setup: func(id uuid.UUID) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
				to, _ := uuid.GenerateRandomUUID()
				return query.NewKeyValueRelationshipProperty(id, to, "", "")
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv, tv := tt.setup(tt.want)
			id, _ := query.UUID(kv)
			if !reflect.DeepEqual(id, tt.want) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, id, tt.want)
			}

			// The transpose
			idTV, _ := query.UUID(tv)
			if !reflect.DeepEqual(idTV, tt.want) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, idTV, tt.want)
			}
		})
	}
}

// func TestUUIDTranspose(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		setup func(id *uuid.UUID) *keyvalue.KeyValue
// 		want  *uuid.UUID
// 	}{
// 		{
// 			name: "Vertex",
// 			setup: func(id *uuid.UUID) *keyvalue.KeyValue {
// 				return keyvalue.NewKeyValueVertexTranspose(id, "")
// 			},
// 			want: func() *uuid.UUID {
// 				id, _ := uuid.GenerateRandomUUID()
// 				return id
// 			}(),
// 		},
// 		{
// 			name: "Properties",
// 			setup: func(id *uuid.UUID) *keyvalue.KeyValue {
// 				return keyvalue.NewKeyValuePropertyTranspose(id, "", "")
// 			},
// 			want: func() *uuid.UUID {
// 				id, _ := uuid.GenerateRandomUUID()
// 				return id
// 			}(),
// 		},
// 		{
// 			name: "Relationship",
// 			setup: func(id *uuid.UUID) *keyvalue.KeyValue {
// 				to, _ := uuid.GenerateRandomUUID()
// 				return keyvalue.NewKeyValueRelationshipTranspose(id, to, "", 5)
// 			},
// 			want: func() *uuid.UUID {
// 				id, _ := uuid.GenerateRandomUUID()
// 				return id
// 			}(),
// 		},
// 		{
// 			name: "Relationshipproperties",
// 			setup: func(id *uuid.UUID) *keyvalue.KeyValue {
// 				to, _ := uuid.GenerateRandomUUID()
// 				return keyvalue.NewKeyValueRelationshipPropertyTranspose(id, to, "", "")
// 			},
// 			want: func() *uuid.UUID {
// 				id, _ := uuid.GenerateRandomUUID()
// 				return id
// 			}(),
// 		},
// 	}
// 	for i, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			kv := tt.setup(tt.want)
// 			if !reflect.DeepEqual(kv.UUID(), tt.want) {
// 				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, kv.UUID(), tt.want)
// 			}
// 		})
// 	}
// }
