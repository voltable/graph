package openCypher

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/google/uuid"
	"github.com/voltable/graph"
	"github.com/voltable/graph/operators"
	"github.com/voltable/graph/operators/ir"
	"github.com/voltable/graph/query/openCypher/parser"
	"github.com/voltable/graph/widecolumnstore"
	"strconv"
	"strings"
)

const Null = "null"

type cypherWalker struct {
	storage widecolumnstore.Storage
	statistics *graph.Statistics
	*parser.BaseCypherListener
	errors []antlr.ErrorNode
	stack  StackExpr
	plan   []operators.Operator
}

func NewCypherWalker(storage widecolumnstore.Storage, statistics *graph.Statistics) cypherWalker {
	return cypherWalker{
		statistics: statistics,
		storage: storage,
		stack:   StackExpr{},
		plan:    make([]operators.Operator, 0),
	}
}

func (l *cypherWalker) GetQueryPlan() ([]operators.Operator) {
	return l.plan
}

func (l *cypherWalker) EnterOC_Return(c *parser.OC_ReturnContext) {
	l.stack = l.stack.push(&ir.Return{})
}

func (l *cypherWalker) ExitOC_Return(c *parser.OC_ReturnContext) {
	NotReturn := func(n interface{}) bool {
		_, ok := n.(*ir.Return)
		return !ok
	}

	items := make([]*ir.ReturnItem, 0)

	for n := l.stack.top(); NotReturn(n); n = l.stack.top() {
		l.stack, _ = l.stack.pop()

		if k, ok := n.(*ir.ReturnItem); ok {
			items = append(items, k)
		}
	}

	op, _ := operators.NewProduceResults(l.storage, l.statistics, items)
	l.plan = append(l.plan, op)
}

func (l *cypherWalker) EnterOC_ReturnItem(c *parser.OC_ReturnItemContext) {
	l.stack = l.stack.push(&ir.ReturnItem{})
}

func (l *cypherWalker) ExitOC_ReturnItem(c *parser.OC_ReturnItemContext) {
	NotReturnItem := func(n interface{}) bool {
		_, ok := n.(*ir.ReturnItem)
		return !ok
	}

	var variable ir.Variable
	var expression  *ir.Expression

	for n := l.stack.top(); NotReturnItem(n); n = l.stack.top() {
		l.stack, _ = l.stack.pop()
		if k, ok := n.(ir.Variable); ok {
			variable = k
		} else if k, ok := n.(*ir.Expression); ok {
			expression = k
		}
	}

	item := l.stack.top().(*ir.ReturnItem)
	item.Variable = variable
	if expression != nil {
		item.Expression = *expression
	}
}

func (l *cypherWalker) ExitOC_Expression(c *parser.OC_ExpressionContext) {
	var n interface{}
	l.stack, n = l.stack.pop()
	l.stack = l.stack.push(&ir.Expression{Value: n})
}

func (l *cypherWalker) EnterOC_Create(c *parser.OC_CreateContext) {
	l.stack = l.stack.push(&ir.Create{})
}

func (l *cypherWalker) ExitOC_Create(c *parser.OC_CreateContext) {
	NotCreate := func(n interface{}) bool {
		_, ok := n.(*ir.Create)
		return !ok
	}
	nodes := make([]*ir.Node, 0)
	relationships := make([]*ir.Relationship, 0)
	for n := l.stack.top(); NotCreate(n); n = l.stack.top() {
		l.stack, _ = l.stack.pop()
		if k, ok := n.(*ir.Node); ok {
			nodes = append(nodes, k)
		} else if  k, ok := n.(*ir.Relationship); ok {
			relationships = append(relationships, k)
		}
	}

	op, _ := operators.NewCreate(l.storage, l.statistics, nodes, relationships)
	l.plan = append(l.plan, op)
}

func (l *cypherWalker) EnterOC_NodePattern(c *parser.OC_NodePatternContext) {
	l.stack = l.stack.push(&ir.Node{Id: uuid.New()})
}

func (l *cypherWalker) ExitOC_NodePattern(c *parser.OC_NodePatternContext) {
	NotNode := func(n interface{}) bool {
		_, ok := n.(*ir.Node)
		return !ok
	}

	var variable ir.Variable
	var label ir.Label
	var properties *ir.Properties

	for n := l.stack.top(); NotNode(n); n = l.stack.top() {
		l.stack, _ = l.stack.pop()
		if k, ok := n.(ir.Variable); ok {
			variable = k
		} else if k, ok := n.(ir.Label); ok {
			label = k
		} else if k, ok := n.(*ir.Properties); ok{
			properties = k
		}
	}

	node := l.stack.top().(*ir.Node)
	node.Variable = variable
	node.Label = label
	node.Properties = properties
}

func (l *cypherWalker) EnterOC_RelationshipPattern(c *parser.OC_RelationshipPatternContext) {
	l.stack = l.stack.push(&ir.Relationship{Id: uuid.New()})
}

func (l *cypherWalker) ExitOC_RelationshipPattern(c *parser.OC_RelationshipPatternContext) {
	NoRelationship := func(n interface{}) bool {
		_, ok := n.(*ir.Relationship)
		return !ok
	}

	var typeName ir.Type
	var variable ir.Variable
	var properties *ir.Properties

	for n := l.stack.top(); NoRelationship(n); n = l.stack.top() {
		l.stack, _ = l.stack.pop()
		if k, ok := n.(ir.Variable); ok {
			variable = k
		} else if k, ok := n.(ir.Type); ok {
			typeName = k
		}else if k, ok := n.(*ir.Properties); ok{
			properties = k
		}
	}

	relationship := l.stack.top().(*ir.Relationship)
	relationship.Variable = variable
	relationship.Type = typeName
	relationship.Properties = properties
}

func (l *cypherWalker) EnterOC_LabelName(c *parser.OC_LabelNameContext) {
	s := c.OC_SchemaName().(*parser.OC_SchemaNameContext)
	sn := s.OC_SymbolicName().(*parser.OC_SymbolicNameContext)
	labelName := sn.UnescapedSymbolicName().GetText()
	l.stack = l.stack.push( ir.Label(labelName))
}

func (l *cypherWalker) EnterOC_RelTypeName(c *parser.OC_RelTypeNameContext) {
	s := c.OC_SchemaName().(*parser.OC_SchemaNameContext)
	sn := s.OC_SymbolicName().(*parser.OC_SymbolicNameContext)
	typeName := sn.UnescapedSymbolicName().GetText()
	l.stack = l.stack.push( ir.Type(typeName))
}

func (l *cypherWalker) EnterOC_Variable(c *parser.OC_VariableContext) {
	s := c.OC_SymbolicName().(*parser.OC_SymbolicNameContext)
	variable := s.UnescapedSymbolicName().GetText()
	l.stack = l.stack.push( ir.Variable(variable))
}

func (l *cypherWalker) EnterOC_MapLiteral(c *parser.OC_MapLiteralContext) {
	l.stack = l.stack.push(ir.NewMapLiteral())
}

func (l *cypherWalker) ExitOC_MapLiteral(c *parser.OC_MapLiteralContext) {
	var expression *ir.Expression
	items := make(map[ir.Key]*ir.Expression, 0)

	NotMapLiteral := func(n interface{}) bool {
		_, ok := n.(*ir.MapLiteral)

		return !ok
	}
	for n := l.stack.top(); NotMapLiteral(n); n = l.stack.top() {
		l.stack, _ = l.stack.pop()
		if key, ok := n.(ir.Key); ok {
			items[key] = expression
		} else if value, ok := n.(*ir.Expression); ok {
			expression = value
		}
	}

	list := l.stack.top().(*ir.MapLiteral)

	list.Items = items
}

func (l *cypherWalker) EnterOC_Properties(c *parser.OC_PropertiesContext) {
	l.stack = l.stack.push(ir.NewProperties())
}

func (l *cypherWalker) ExitOC_Properties(c *parser.OC_PropertiesContext) {
	var mapLiteral *ir.MapLiteral

	NotProperties := func(n interface{}) bool {
		_, ok := n.(*ir.Properties)
		return !ok
	}
	for n := l.stack.top(); NotProperties(n); n = l.stack.top() {
		l.stack, _ = l.stack.pop()
		if m, ok := n.(*ir.MapLiteral); ok {
			mapLiteral = m
		}
	}

	properties := l.stack.top().(*ir.Properties)
	properties.Map = mapLiteral
}


func (l *cypherWalker) EnterOC_PropertyKeyName(c *parser.OC_PropertyKeyNameContext) {
	l.stack = l.stack.push(ir.Key(c.GetText()))
}


func (l *cypherWalker)  EnterOC_UnaryAddOrSubtractExpression(c *parser.OC_UnaryAddOrSubtractExpressionContext) {
	tt := c.GetText()
	if strings.HasPrefix(tt, "-") {
		l.stack = l.stack.push(ir.Subtraction)
	}
}

func (l *cypherWalker) EnterOC_BooleanLiteral(c *parser.OC_BooleanLiteralContext) {
	if t := c.TRUE(); t != nil {
		l.stack = l.stack.push(true)
		return
	} else if t := c.FALSE(); t != nil {
		l.stack = l.stack.push(false)
		return
	}
}

func (l *cypherWalker) EnterOC_NumberLiteral(c *parser.OC_NumberLiteralContext) {
	// We lose the negative number prefix so must prefix it back
	var prefix string
	if a := l.stack.top(); a == ir.Subtraction {
		l.stack, _ = l.stack.pop()
		prefix = "-"
	}
	if i, ok := c.OC_IntegerLiteral().(*parser.OC_IntegerLiteralContext); ok {
		ii, _ := strconv.Atoi(prefix + i.GetText())
		l.stack = l.stack.push(ii)
		return
	}
	if i, ok := c.OC_DoubleLiteral().(*parser.OC_DoubleLiteralContext); ok {
		ii, _ := strconv.ParseFloat(prefix + i.GetText(), 64)
		l.stack = l.stack.push(ii)
		return
	}
}

func (l *cypherWalker) EnterOC_ListLiteral(c *parser.OC_ListLiteralContext) {
	l.stack = l.stack.push(&ir.ListLiteral{})
}

func (l *cypherWalker) ExitOC_ListLiteral(c *parser.OC_ListLiteralContext) {
	items := make([]*ir.Expression, 0)

	NotList := func(n interface{}) bool {
		_, ok := n.(*ir.ListLiteral)
		return !ok
	}
	for n := l.stack.top(); NotList(n); n = l.stack.top() {
		l.stack, _ = l.stack.pop()

		if expression, ok := n.(*ir.Expression); ok {
			items = append(items, expression)
		}

	}

	// list is in the wrong order after coming off the stack so need to reverse
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}

	list := l.stack.top().(*ir.ListLiteral)
	list.Items = items
}


func (l *cypherWalker) EnterOC_Literal(c *parser.OC_LiteralContext) {
	if n := c.StringLiteral(); n != nil {
		stringLiteral := func(r rune) bool {
			return r == '"' || r == '\''
		}

		str := strings.TrimRightFunc(strings.TrimLeftFunc(c.GetText(),stringLiteral),stringLiteral)
		l.stack = l.stack.push(ir.String(str))
		return
	}

	if n := c.NULL(); n != nil {
		l.stack = l.stack.push(nil)
		return
	}
}
