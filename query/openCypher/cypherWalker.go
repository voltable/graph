package openCypher

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/voltable/graph/query/openCypher/parser"
)

type cypherWalker struct {
	*parser.BaseCypherListener
	errors []antlr.ErrorNode
	stack  []interface{}
}
