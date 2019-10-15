package query_test

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/voltable/graph/encoding/wcs"
	"github.com/voltable/graph/query"
	"github.com/voltable/graph/widecolumnstore"
)

func TestRelationshipUUID(t *testing.T) {

	tests := []struct {
		name          string
		setup         func(id uuid.UUID, to uuid.UUID) widecolumnstore.KeyValue
		want          uuid.UUID
		wantTranspose uuid.UUID
	}{
		{
			name: "Relationship",
			setup: func(id uuid.UUID, to uuid.UUID) widecolumnstore.KeyValue {
				return wcs.NewKeyValueRelationship(id, to, "", 5)
			},
			want: func() uuid.UUID {
				id := uuid.New()
				return id
			}(),
			wantTranspose: func() uuid.UUID {
				id := uuid.New()
				return id
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := tt.setup(tt.want, tt.wantTranspose)

			id, _ := query.UUID(&kv)
			if !reflect.DeepEqual(id, tt.want) {
				t.Errorf("%d. %q: UUID() = %v, want %v", i, tt.name, id, tt.want)
			}
		})
	}
}
