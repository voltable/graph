package keyvalue_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.Graph/keyvalue"
	"github.com/RossMerr/Caudex.Graph/uuid"
)

func TestUUID(t *testing.T) {
	tests := []struct {
		name  string
		setup func(id uuid.UUID) *keyvalue.KeyValue
		want  uuid.UUID
	}{
		{
			name: "Vertex",
			setup: func(id uuid.UUID) *keyvalue.KeyValue {
				return keyvalue.NewKeyValue("person", id[:], keyvalue.US, keyvalue.Vertex)
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Properties",
			setup: func(id uuid.UUID) *keyvalue.KeyValue {
				return keyvalue.NewKeyValue("", id[:], keyvalue.US, keyvalue.Properties, keyvalue.US, []byte(""))
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Relationship",
			setup: func(id uuid.UUID) *keyvalue.KeyValue {
				to, _ := uuid.GenerateRandomUUID()
				return keyvalue.NewKeyValue(to[:], id[:], keyvalue.US, keyvalue.Relationship, keyvalue.US, []byte(""), keyvalue.US, to[:])
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Relationshipproperties",
			setup: func(id uuid.UUID) *keyvalue.KeyValue {
				to, _ := uuid.GenerateRandomUUID()
				return keyvalue.NewKeyValue("", id[:], keyvalue.US, keyvalue.Relationshipproperties, keyvalue.US, []byte(""), keyvalue.US, to[:])
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup(tt.want)
			if !reflect.DeepEqual(kv.UUID(), tt.want) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, kv.UUID(), tt.want)
			}
		})
	}
}

func TestUUIDTranspose(t *testing.T) {
	tests := []struct {
		name  string
		setup func(id uuid.UUID) *keyvalue.KeyValue
		want  uuid.UUID
	}{
		{
			name: "Vertex",
			setup: func(id uuid.UUID) *keyvalue.KeyValue {
				return keyvalue.NewKeyValue(id[:], keyvalue.Vertex, keyvalue.US, []byte(""), keyvalue.US, id[:])
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Properties",
			setup: func(id uuid.UUID) *keyvalue.KeyValue {
				return keyvalue.NewKeyValue("", keyvalue.Properties, keyvalue.US, []byte(""), keyvalue.US, id[:])
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Relationship",
			setup: func(id uuid.UUID) *keyvalue.KeyValue {
				to, _ := uuid.GenerateRandomUUID()
				return keyvalue.NewKeyValue(id[:], keyvalue.Relationship, keyvalue.US, []byte(""), keyvalue.US, to[:], keyvalue.US, id[:])
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
		{
			name: "Relationshipproperties",
			setup: func(id uuid.UUID) *keyvalue.KeyValue {
				to, _ := uuid.GenerateRandomUUID()
				return keyvalue.NewKeyValue("", keyvalue.Relationshipproperties, keyvalue.US, []byte(""), keyvalue.US, to[:], keyvalue.US, id[:])
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup(tt.want)
			if !reflect.DeepEqual(kv.UUID(), tt.want) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, kv.UUID(), tt.want)
			}
		})
	}
}
