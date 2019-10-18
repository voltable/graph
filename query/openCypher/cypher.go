package openCypher

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/voltable/graph/query/openCypher/parser"
)

func Parser(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCypherLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCypherParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.OC_Cypher()

	// Finally parse the expression (by walking the tree)
	var walker cypherWalker
	antlr.ParseTreeWalkerDefault.Walk(&walker, tree)

}
