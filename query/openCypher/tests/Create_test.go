package features

import (
	"errors"
	"fmt"
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/query"
	"github.com/voltable/graph/query/cypher"
	"github.com/voltable/graph/query/openCypher"
	"github.com/voltable/graph/widecolumnstore"
	"github.com/voltable/graph/widecolumnstore/storage/memorydb"
	"reflect"
	"strconv"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

type graphFeature struct {
	graph       graph.Graph
	storage     widecolumnstore.Storage
	queryResult *graph.Query
}

var (
	storage = func() widecolumnstore.Storage {
		s, _ := memorydb.NewStorageEngine()
		return s
	}()

	options = graph.NewOptions(openCypher.QueryType, memorydb.StorageType)
)

func (s *graphFeature) anyGraph() error {

	cypher.RegisterEngine()

	storage, err := memorydb.NewStorageEngine()
	if err != nil {
		return err
	}

	options := graph.NewOptions(openCypher.QueryType, memorydb.StorageType)
	g, err := query.NewGraphEngineFromStorageEngine(storage, options)
	if err != nil {
		return err
	}

	s.graph = g

	return nil
}

func (s *graphFeature) emptyGraph() error {
	ge, _ := query.NewGraphEngineFromStorageEngine(storage, options)
	s.graph = ge

	return nil
}

func (s *graphFeature) executingQuery(arg1 *gherkin.DocString) error {
	result, err := s.graph.Query(arg1.Content)
	s.queryResult = result
	return err
}

func (s *graphFeature) theResultShouldBeEmpty() error {
	if s.queryResult.Results.IsEmpty() {
		return nil
	}
	return errors.New("Result found")
}


func (s *graphFeature) theResultShouldBe(arg1 *gherkin.DataTable) error {

	transformed := graph.NewTable()
	for r, row := range arg1.Rows {
		if r == 0 {
			for _, cell := range row.Cells {
				transformed.Columns = append(transformed.Columns, graph.Column{
					Field: cell.Value,
					Rows: make([]interface{}, 0),
				})
			}

		} else {
			for c, cell := range row.Cells  {
				transformed.Columns[c].Rows = append(transformed.Columns[c].Rows, cell.Value)
			}
		}
	}

	for c, column := range transformed.Columns  {
		field :=  s.queryResult.Results.Columns[c].Field

		if column.Field != field {
			fmt.Errorf("theResultShouldBe: column %d: field %s but was %s\n", c, column.Field, field)
		}

		for r, row := range column.Rows {
			var value string
			if v, ok := s.queryResult.Results.Columns[c].Rows[r].(float64); ok {
				rowAsFloat, _ := strconv.ParseFloat(row.(string), 64)
				if v != rowAsFloat {
					return fmt.Errorf("theResultShouldBe: column %d, row %d: %+v but was %+v\n", c, r, row, value)
				}

			} else if s.queryResult.Results.Columns[c].Rows[r] == nil {
				value = "null"

				if value != row {
					return fmt.Errorf("theResultShouldBe: column %d, row %d: %+v but was %+v\n", c, r, row, value)
				}
			}  else if _, ok := s.queryResult.Results.Columns[c].Rows[r].(*ir.MapLiteral); ok {
				// Go Maps are unordered so can't easily do the compare
				// TODO need to find a solution
			}else {
				value = fmt.Sprintf("%v", s.queryResult.Results.Columns[c].Rows[r])

				if value != row {
					return fmt.Errorf("theResultShouldBe: column %d, row %d: %+v but was %+v\n", c, r, row, value)
				}
			}
		}
	}

	return nil
}


func (s *graphFeature) theSideEffectsShouldBe(arg1 *gherkin.DataTable) error {
	found := make(map[string]interface{})
	t := s.queryResult.Statistics.DbHits
	e := reflect.ValueOf(&t).Elem()
	typeOfT := e.Type()

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		tag := typeOfT.Field(i).Tag.Get("json")

		found[tag] = f.Interface()
	}

	for i, row := range arg1.Rows {
		key := row.Cells[0].Value
		value, _ := strconv.Atoi(row.Cells[1].Value)

		if found[key] != value {
			return fmt.Errorf("theSideEffectsShouldBe: %d: %s %+v but was %+v\n", i, key, value, found[key])
		}
	}

	return nil
}

func (s *graphFeature) noSideEffects() error {
	t := s.queryResult.Statistics.DbHits
	e := reflect.ValueOf(&t).Elem()
	typeOfT := e.Type()

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		tag := typeOfT.Field(i).Tag.Get("json")

		if value, ok := f.Interface().(int); ok && value != 0 {
			return fmt.Errorf("noSideEffects: %d: %s %+v but was %+v\n", i, tag, 0, value)
		}
	}

	return nil
}

func (s *graphFeature) anEmptyGraph() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {


	cypher.RegisterEngine()

	ge, _ := query.NewGraphEngineFromStorageEngine(storage, options)
	g := graphFeature{
		graph:ge,
	}

	s.Step(`^an empty graph$`, g.emptyGraph)
	s.Step(`^any graph$`, g.anyGraph)
	s.Step(`^executing query:$`, g.executingQuery)
	s.Step(`^the result should be empty$`, g.theResultShouldBeEmpty)
	s.Step(`^the result should be:`, g.theResultShouldBe)
	s.Step(`^the side effects should be:$`, g.theSideEffectsShouldBe)
	s.Step(`^no side effects$`, g.noSideEffects)
	s.Step(`^an empty graph$`, g.anEmptyGraph)
}
