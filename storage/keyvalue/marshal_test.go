package keyvalue_test

import (
	"reflect"
	"strings"
	"testing"

	graph "github.com/RossMerr/Caudex.Graph"
	"github.com/RossMerr/Caudex.Graph/storage/keyvalue"
)

func TestProperty(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		want1   string
		wantErr bool
	}{
		{
			name:  "Property",
			want:  "123-321",
			want1: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := keyvalue.PropertiesID(tt.want, tt.want1)
			split := strings.Split(string(buf), keyvalue.US)
			got, got1, err := keyvalue.Property(split)
			if (err != nil) != tt.wantErr {
				t.Errorf("Property() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Property() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Property() got = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRelationship(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		want1   string
		wantErr bool
	}{
		{
			name:  "Relationship",
			want:  "123-321",
			want1: "friend",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := keyvalue.RelationshipID(tt.want, tt.want1)
			split := strings.Split(string(buf), keyvalue.US)
			got, got1, err := keyvalue.Relationship(split)
			if (err != nil) != tt.wantErr {
				t.Errorf("Relationship() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Relationship() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Relationship() got = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRelationshipProperty(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		want1   string
		want2   string
		wantErr bool
	}{
		{
			name:  "RelationshipProperty",
			want:  "123-321",
			want1: "321-123",
			want2: "key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := keyvalue.RelationshipPropertiesID(tt.want, tt.want1, tt.want2)
			split := strings.Split(string(buf), keyvalue.US)
			got, got1, got2, err := keyvalue.RelationshipProperties(split)
			if (err != nil) != tt.wantErr {
				t.Errorf("Relationship() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Relationship() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Relationship() got = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Relationship() got = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		name   string
		vertex *graph.Vertex
		want   []*keyvalue.KeyValue
	}{
		{
			name: "vertex",
			vertex: func() *graph.Vertex {
				v, _ := graph.NewVertex()
				v.SetLabel("person")
				v.SetProperty("name", "john smith")
				t, _ := graph.NewVertex()
				id, _ := graph.ParseUUID(t.ID())
				e, _ := graph.NewEdgeFromID(id)
				e.SetRelationshipType("friend")
				e.SetProperty("years", 10)
				v.AddEdge(e)
				return v
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := keyvalue.Marshal(tt.vertex)
			v := keyvalue.Unmarshal(got...)
			if !reflect.DeepEqual(v, tt.vertex) {
				t.Errorf("Marshal() = %v, want %v", v, tt.vertex)
			}
		})
	}
}

func TestMarshalTranspose(t *testing.T) {
	tests := []struct {
		name   string
		vertex *graph.Vertex
		want   []*keyvalue.KeyValue
	}{
		{
			name: "vertex",
			vertex: func() *graph.Vertex {
				v, _ := graph.NewVertex()
				v.SetLabel("person")
				v.SetProperty("name", "john smith")
				t, _ := graph.NewVertex()
				id, _ := graph.ParseUUID(t.ID())
				e, _ := graph.NewEdgeFromID(id)
				e.SetRelationshipType("friend")
				e.SetProperty("years", 10)
				v.AddEdge(e)
				return v
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := keyvalue.MarshalTranspose(tt.vertex)
			v := keyvalue.UnmarshalTranspose(got...)
			if !reflect.DeepEqual(v, tt.vertex) {
				t.Errorf("Marshal() = %v, want %v", v, tt.vertex)
			}
		})
	}
}
