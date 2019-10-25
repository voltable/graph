package operators

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Nullary = (*Create)(nil)

// Create fetches all tuples with a specific label.
type Create struct {
	nodes []*ir.Node
	relationships []*ir.Relationship
	storage widecolumnstore.Storage
}

// NewCreate returns a Create
func NewCreate(storage widecolumnstore.Storage, nodes []*ir.Node, relationships []*ir.Relationship) (*Create, error) {
	return &Create{
		storage: storage,
		nodes: nodes,
		relationships:relationships,
	}, nil
}

func (s *Create) Next() (widecolumnstore.Iterator, graph.Statistics) {
	action := &ir.Actions{}
	keyValues := make([]*widecolumnstore.KeyValue, 0)

	for _, n := range s.nodes {
		keyValues = append(keyValues, n.Marshal(action)...)
	}

	for _, n := range s.relationships {
		keyValues = append(keyValues, n.Marshal(action)...)
	}

	statistics := graph.NewStatistics()
	statistics.DbHits.CreateLabels = action.Labels
	statistics.DbHits.CreateNodes = action.Nodes
	statistics.DbHits.CreateProperties = action.Properties
	statistics.DbHits.CreateRelationships = action.Relationships
	statistics.DbHits.CreateTypes = action.Types
	statistics.Rows += len(keyValues)

	_ = s.storage.Create(keyValues...)

	return func() (widecolumnstore.KeyValue, bool) {
		return widecolumnstore.KeyValue{}, false
	}, statistics
}

func (s *Create) Op() {}
