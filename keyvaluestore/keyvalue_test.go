package keyvaluestore_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/keyvaluestore"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

func TestUUID(t *testing.T) {
	tests := []struct {
		name  string
		setup func(id *uuid.UUID) (*keyvaluestore.KeyValue, *keyvaluestore.KeyValue)
		want  *uuid.UUID
	}{
		{
			name: "Vertex",
			setup: func(id *uuid.UUID) (*keyvaluestore.KeyValue, *keyvaluestore.KeyValue) {
				return keyvaluestore.NewKeyValueVertex(id, "person")
			},
			want: func() *uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Properties",
			setup: func(id *uuid.UUID) (*keyvaluestore.KeyValue, *keyvaluestore.KeyValue) {
				return keyvaluestore.NewKeyValueProperty(id, "", "")
			},
			want: func() *uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Relationship",
			setup: func(id *uuid.UUID) (*keyvaluestore.KeyValue, *keyvaluestore.KeyValue) {
				to, _ := uuid.GenerateRandomUUID()
				return keyvaluestore.NewKeyValueRelationship(id, to, "", 5)
			},
			want: func() *uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Relationshipproperties",
			setup: func(id *uuid.UUID) (*keyvaluestore.KeyValue, *keyvaluestore.KeyValue) {
				to, _ := uuid.GenerateRandomUUID()
				return keyvaluestore.NewKeyValueRelationshipProperty(id, to, "", "")
			},
			want: func() *uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv, tv := tt.setup(tt.want)
			if !reflect.DeepEqual(kv.UUID(), tt.want) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, kv.UUID(), tt.want)
			}

			// The transpose
			if !reflect.DeepEqual(tv.UUID(), tt.want) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, tv.UUID(), tt.want)
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
