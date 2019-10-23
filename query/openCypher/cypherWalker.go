package openCypher

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/voltable/graph/operators"
	"github.com/voltable/graph/widecolumnstore"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/voltable/graph/query/openCypher/parser"
)


type node struct {
	variable string
	label string
	properties map[string]interface{}
}

type mapLiteral struct {

}

type keyValue struct {
	key string
	value interface{}
}

type cypherWalker struct {
	storage widecolumnstore.Storage
	*parser.BaseCypherListener
	errors []antlr.ErrorNode
	stack  StackExpr
	plan []operators.Operator
}

func NewCypherWalker(storage widecolumnstore.Storage) cypherWalker {
	return cypherWalker{
		storage:storage,
		stack: StackExpr{},
		plan: make([]operators.Operator, 0),
	}
}

func (l *cypherWalker) GetQueryPlan() []operators.Operator {
	return l.plan
}

func (l *cypherWalker) EnterOC_Create(c *parser.OC_CreateContext) {
	l.stack = l.stack.push(&node{})
}

func (l *cypherWalker) ExitOC_Create(c *parser.OC_CreateContext) {
	id := uuid.New()
	var n interface{}
	l.stack, n = l.stack.pop()
	if node, ok := n.(*node); ok {
		 op, _ := operators.NewCreate(l.storage, id, node.variable, node.label, node.properties)
		 l.plan = append(l.plan, op)
	}
}

func (l *cypherWalker) EnterOC_LabelName(c *parser.OC_LabelNameContext) {
	s := c.OC_SchemaName().(*parser.OC_SchemaNameContext)
	sn := s.OC_SymbolicName().(*parser.OC_SymbolicNameContext)
	label := sn.UnescapedSymbolicName().GetText()

	i := l.stack.top()
	if node, ok := i.(*node); ok {
		node.label = label
	}
}

func (l *cypherWalker) EnterOC_Variable(c *parser.OC_VariableContext) {
	s := c.OC_SymbolicName().(*parser.OC_SymbolicNameContext)
	variable := s.UnescapedSymbolicName().GetText()

	i := l.stack.top()
	if node, ok:= i.(*node); ok {
		node.variable = variable
	}
}

func (l *cypherWalker) EnterOC_Properties(c *parser.OC_PropertiesContext) {
	fmt.Printf("EnterOC_Properties : %s\n", c)
}

func (l *cypherWalker) ExitOC_Properties(c *parser.OC_PropertiesContext) {
	fmt.Printf("ExitOC_Properties : %s\n", c)

	var n interface{}
	l.stack, n = l.stack.pop()
	if properties, ok := n.(map[string]interface{}); ok {
		i := l.stack.top()
		if node, ok:= i.(*node); ok {
			node.properties = properties
		}
	}
}

func (l *cypherWalker) EnterOC_MapLiteral(c *parser.OC_MapLiteralContext) {
	fmt.Printf("EnterOC_MapLiteral : %s\n", c.GetText())

	l.stack = l.stack.push(mapLiteral{})
}

func (l *cypherWalker) ExitOC_MapLiteral(c *parser.OC_MapLiteralContext) {
	fmt.Printf("ExitOC_MapLiteral : %s\n", c.GetText())

	var literal interface{}

	properties := make(map[string]interface{}, 0)

	NotMapLiteral := func(n interface{}) bool {
		_, ok := n.(mapLiteral)
		return !ok
	}
	for n := l.stack.top(); NotMapLiteral(n); l.stack, n = l.stack.pop() {
		if kv, ok := n.(keyValue); ok {
			properties[kv.key] = literal
		} else {
			literal = n
		}
	}

	l.stack.push(properties)
}

func (l *cypherWalker) EnterOC_PropertyKeyName(c *parser.OC_PropertyKeyNameContext) {
	fmt.Printf("EnterOC_PropertyKeyName : %s\n", c.GetText())

	l.stack = l.stack.push(keyValue{key:c.GetText()})
}

func (l *cypherWalker) EnterOC_Literal(c *parser.OC_LiteralContext) {
	fmt.Printf("EnterOC_Literal : %s\n", c.GetText())

	if b, ok := c.OC_BooleanLiteral().(*parser.OC_BooleanLiteralContext); ok {
		if t := b.TRUE(); t != nil {
			l.stack = l.stack.push(true)
			return
		} else if t := b.FALSE(); t != nil {
			l.stack = l.stack.push(false)
			return
		}
	}

	if n, ok := c.OC_NumberLiteral().(*parser.OC_NumberLiteralContext); ok {
		if i, ok := n.OC_IntegerLiteral().(*parser.OC_IntegerLiteralContext); ok {
			ii, _ := strconv.Atoi(i.GetText())
			l.stack = l.stack.push(ii)
			return
		}
		if i, ok := n.OC_DoubleLiteral().(*parser.OC_DoubleLiteralContext); ok {
			ii, _ := strconv.ParseFloat(i.GetText(), 64)
			l.stack = l.stack.push(ii)
			return
		}
	}

	l.stack = l.stack.push(c.GetText())
}