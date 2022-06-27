package operators

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Nullary = (*Create)(nil)

// Create fetches all tuples with a specific label.
type Create struct {
	nodes         []*ir.Node
	relationships []*ir.Relationship
	storage       widecolumnstore.Storage
	statistics    *graph.Statistics
}

// NewCreate returns a Create
func NewCreate(storage widecolumnstore.Storage, statistics *graph.Statistics, nodes []*ir.Node, relationships []*ir.Relationship) (*Create, error) {
	return &Create{
		storage:       storage,
		statistics:    statistics,
		nodes:         nodes,
		relationships: relationships,
	}, nil
}

func (s *Create) Next() (widecolumnstore.Iterator, error) {
	action := &ir.Actions{}
	keyValues := make([]*widecolumnstore.KeyValue, 0)

	for _, n := range s.nodes {
		keyValues = append(keyValues, n.Marshal(action)...)
	}

	for _, n := range s.relationships {
		keyValues = append(keyValues, n.Marshal(action)...)
	}

	s.statistics.DbHits.CreateLabels += action.Labels
	s.statistics.DbHits.CreateNodes += action.Nodes
	s.statistics.DbHits.CreateProperties += action.Properties
	s.statistics.DbHits.CreateRelationships += action.Relationships
	s.statistics.DbHits.CreateTypes += action.Types
	s.statistics.Rows += len(keyValues)

	err := s.storage.Create(keyValues...)

	return func() (widecolumnstore.KeyValue, bool) {
		return widecolumnstore.KeyValue{}, false
	}, err
}

func (s *Create) Op() {}
