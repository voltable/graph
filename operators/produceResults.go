package operators

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Operator = (*ProduceResults)(nil)

// ProduceResults fetches all tuples with a specific label.
type ProduceResults struct {
	storage widecolumnstore.Storage
	items []*ir.ReturnItem
	statistics *graph.Statistics
}

// NewProduceResults returns a Return
func NewProduceResults(storage widecolumnstore.Storage, statistics *graph.Statistics, items []*ir.ReturnItem) (*ProduceResults, error) {
	return &ProduceResults{
		storage: storage,
		statistics: statistics,
		items: items,
	}, nil
}

func (s *ProduceResults) Next() *graph.Table {
	table := graph.NewTable()
	for _, item := range s.items {
		column := graph.Column{
			Field: string(item.Variable),
			Rows: item.Expression.Evaluate(),
		}
		table.Columns = append(table.Columns, column)
	}

	return table
}

func (s *ProduceResults) Op() {}
