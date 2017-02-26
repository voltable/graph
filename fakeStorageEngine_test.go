package graphs

type fakeStorageEngine struct {
	vertices map[string]Vertex
}

func (f *fakeStorageEngine) Create(c ...Vertex) error {
	for _, v := range c {
		f.vertices[v.ID] = v
	}
	return nil
}

func (f *fakeStorageEngine) Delete(c ...Vertex) error {
	for _, v := range c {
		delete(f.vertices, v.ID)
	}

	return nil
}

func (f *fakeStorageEngine) Find(ID string) (*Vertex, error) {
	v := f.vertices[ID]
	return &v, nil
}

func (f *fakeStorageEngine) Update(c ...Vertex) error {
	f.Create(c...)
	return nil
}
