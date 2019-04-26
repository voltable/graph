package query_test

import (
	"reflect"
	"testing"

	"github.com/voltable/graph/query"
	"github.com/voltable/graph/uuid"
	"github.com/voltable/graph/widecolumnstore"
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

func TestRelationshipUUID(t *testing.T) {

	tests := []struct {
		name          string
		setup         func(id uuid.UUID, to uuid.UUID) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue)
		want          uuid.UUID
		wantTranspose uuid.UUID
	}{
		{
			name: "Relationship",
			setup: func(id uuid.UUID, to uuid.UUID) (*widecolumnstore.KeyValue, *widecolumnstore.KeyValue) {
				return query.NewKeyValueRelationship(id, to, "", 5)
			},
			want: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
			wantTranspose: func() uuid.UUID {
				id, _ := uuid.GenerateRandomUUID()
				return id
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv, tv := tt.setup(tt.want, tt.wantTranspose)

			id, _ := query.UUID(kv)
			if !reflect.DeepEqual(id, tt.want) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, id, tt.want)
			}

			// The transpose
			idTV, _ := query.UUID(tv)
			if !reflect.DeepEqual(idTV, tt.wantTranspose) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, idTV, tt.want)
			}
		})
	}
}
