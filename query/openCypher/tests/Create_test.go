package features

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func anyGraph() error {
	return godog.ErrPending
}

func executingQuery(arg1 *gherkin.DocString) error {
	return godog.ErrPending
}

func theResultShouldBeEmpty() error {
	return godog.ErrPending
}

func theSideEffectsShouldBe(arg1 *gherkin.DataTable) error {
	return godog.ErrPending
}

func anEmptyGraph() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^any graph$`, anyGraph)
	s.Step(`^executing query:$`, executingQuery)
	s.Step(`^the result should be empty$`, theResultShouldBeEmpty)
	s.Step(`^the side effects should be:$`, theSideEffectsShouldBe)
	s.Step(`^an empty graph$`, anEmptyGraph)
}
