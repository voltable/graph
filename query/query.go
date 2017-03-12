package query

type (
	// Iterator is an alias for function to iterate over data.
	Iterator func() (item interface{}, ok bool)

	// MatchPattern is used for describing the pattern to search for represents a walk through a the graph and consists of a sequence of alternating nodes and relationships.
	MatchPattern map[string]Path

	Path interface {
		//Iterate func() Iterator
		Next() Path
		InsertAfter(Path) Path
		InsertBefore(Path) Path
	}
)

func (vp VertexPath) Next() Path {
	return vp.next
}

func (vp EdgePath) Next() Path {
	return vp.next
}

func (vp VertexPath) InsertAfter(next Path) Path {
	vp.next = next
	return next
}

func (vp EdgePath) InsertAfter(next Path) Path {
	vp.next = next
	return next
}

func (vp VertexPath) InsertBefore(next Path) Path {
	return next.InsertAfter(vp)

}

func (vp EdgePath) InsertBefore(next Path) Path {
	return next.InsertAfter(vp)
}
