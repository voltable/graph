package features

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/voltable/graph"
	"github.com/voltable/graph/query"
	"github.com/voltable/graph/query/cypher"
	"github.com/voltable/graph/query/openCypher"
	"github.com/voltable/graph/widecolumnstore"
	"github.com/voltable/graph/widecolumnstore/storage/memorydb"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

type graphFeature struct {
	graph       graph.Graph
	storage     widecolumnstore.Storage
	queryResult *graph.Query
}

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

func (s *graphFeature) executingQuery(arg1 *gherkin.DocString) error {
	result, err := s.graph.Query(arg1.Content)
	s.queryResult = result
	return err
}

func (s *graphFeature) theResultShouldBeEmpty() error {
	if s.queryResult.Results == nil {
		return nil
	}
	return errors.New("Result found")
}

func (s *graphFeature) theSideEffectsShouldBe(arg1 *gherkin.DataTable) error {
	hits := make(map[string]int)

	for _, row := range arg1.Rows  {
		key := row.Cells[0].Value
		value, _ := strconv.Atoi(row.Cells[1].Value)
		hits[key] = value
	}

	t := s.queryResult.Statistics.DbHits
	e := reflect.ValueOf(&t).Elem()
	typeOfT := e.Type()

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		tag := typeOfT.Field(i).Tag.Get("json")

		if hits[tag] != f.Interface() {
			return fmt.Errorf("theSideEffectsShouldBe: %d: %s %+v\n", i, tag, hits[tag])
		} else {
			delete(hits, tag)
		}
	}

	if len(hits) > 0 {
		return godog.ErrPending
	}

	return nil
}

func (s *graphFeature) anEmptyGraph() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	g := graphFeature{}

	s.Step(`^any graph$`, g.anyGraph)
	s.Step(`^executing query:$`, g.executingQuery)
	s.Step(`^the result should be empty$`, g.theResultShouldBeEmpty)
	s.Step(`^the side effects should be:$`, g.theSideEffectsShouldBe)
	s.Step(`^an empty graph$`, g.anEmptyGraph)
}
