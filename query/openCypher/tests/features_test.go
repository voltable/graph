package features_test

import (
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"

	features "github.com/voltable/graph/query/openCypher/tests"
)

func Test_Features(m *testing.T) {
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		features.FeatureContext(s)
	}, godog.Options{
		Output:        colors.Colored(os.Stdout),
		NoColors:      true,
		Format:        "pretty",
		Paths:         []string{"features"},
		StopOnFailure: true,
		Strict:        true,
	})

	os.Exit(status)
}
