package operators

import (
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/widecolumnstore"
)

var _ Operator = (*ProduceResults)(nil)

// ProduceResults fetches all tuples with a specific label.
type ProduceResults struct {
	storage    widecolumnstore.Storage
	statistics *graph.Statistics
	r          *ir.Return
}

// NewProduceResults returns a Return
func NewProduceResults(storage widecolumnstore.Storage, statistics *graph.Statistics, r *ir.Return) (*ProduceResults, error) {
	return &ProduceResults{
		storage:    storage,
		statistics: statistics,
		r:          r,
	}, nil
}

func (s *ProduceResults) Next(iterator widecolumnstore.Iterator) *graph.Table {

	table := graph.NewTable()

	for _, item := range s.r.Items {
		column := graph.Column{
			Field: string(item.Variable),
			Rows:  item.Expression.Evaluate(),
		}
		table.Columns = append(table.Columns, column)
	}

	// for kv, ok := iterator(); ok; kv, ok = iterator() {
	// 	fmt.Printf("%v", kv)
	// }

	return table
}

func (s *ProduceResults) Op() {}
