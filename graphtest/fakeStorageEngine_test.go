package graphs

import (
	"errors"

	"github.com/RossMerr/Caudex.Graph/vertices"
)

type fakeStorageEngine struct {
	vertices map[string]vertices.Vertex
}

func (f *fakeStorageEngine) Create(c ...*vertices.Vertex) error {
	for _, v := range c {
		f.vertices[v.ID()] = *v
	}
	return nil
}

func (f *fakeStorageEngine) Delete(c ...*vertices.Vertex) error {
	for _, v := range c {
		delete(f.vertices, v.ID())
	}

	return nil
}

func (f *fakeStorageEngine) Find(ID string) (*vertices.Vertex, error) {
	if v, ok := f.vertices[ID]; ok {
		return &v, nil
	} else {
		return nil, errors.New("Record Not found")
	}
}

func (f *fakeStorageEngine) Update(c ...*vertices.Vertex) error {
	f.Create(c...)
	return nil
}
