package caudex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEdge(t *testing.T) {
	v := Vertex{ID: "test", Label: "test"}
	var vertex = &v
	vertex.AddDirectedEdge(vertex)
	assert.Equal(t, 1, len(vertex.edges))
	assert.Equal(t, "test", vertex.Label)
}
