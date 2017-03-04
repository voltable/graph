package graphs

import "errors"

type fakeStorageEngine struct {
	vertices map[string]Vertex
}

func (f *fakeStorageEngine) Create(c ...*Vertex) error {
	for _, v := range c {
		f.vertices[v.ID()] = *v
	}
	return nil
}

func (f *fakeStorageEngine) Delete(c ...*Vertex) error {
	for _, v := range c {
		delete(f.vertices, v.ID())
	}

	return nil
}

func (f *fakeStorageEngine) Find(ID string) (*Vertex, error) {
	if v, ok := f.vertices[ID]; ok {
		return &v, nil
	} else {
		return nil, errors.New("Record Not found")
	}
}

func (f *fakeStorageEngine) Update(c ...*Vertex) error {
	f.Create(c...)
	return nil
}
