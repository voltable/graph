package graphs

import "testing"

func Test_CreateVertex(t *testing.T) {
	se := fakeStorageEngine{}
	g := GraphOperation{&se}
	to := testObject{}
	g.CreateVertex(&to)
}

type testObject struct {
}

type fakeStorageEngine struct {
}

func (f *fakeStorageEngine) Create(c ...Vertex) error {
	return nil
}

func (f *fakeStorageEngine) Delete(c ...Vertex) error {
	return nil
}

func (f *fakeStorageEngine) Find(string) (*Vertex, error) {
	return nil, nil
}

func (f *fakeStorageEngine) Update(c ...Vertex) error {
	return nil
}
