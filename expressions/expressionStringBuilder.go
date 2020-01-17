package expressions

import (
	"fmt"
	"reflect"
	"strconv"
)

var _ ExpressionVisitor = (*ExpressionStringBuilder)(nil)

type ExpressionStringBuilder struct {
	ids    map[interface{}]int
	output string
}

func NewExpressionStringBuilder() *ExpressionStringBuilder {
	return &ExpressionStringBuilder{
		ids:    make(map[interface{}]int),
		output: "",
	}
}

func (s *ExpressionStringBuilder) VisitLambda(expr *LambdaExpression) Expression {
	if len(expr.parameters) == 1 {
		s.Visit(expr.parameters[0])

	} else {
		s.VisitExpressions("(", expr.parameters, ")", ", ")

	}
	s.Out(" => ")
	s.Visit(expr.body)

	return expr
}

func (s *ExpressionStringBuilder) VisitExpressions(open string, expressions []*ParameterExpression, close, seperator string) {
	s.Out(open)
	isFirst := true
	for _, expr := range expressions {
		if isFirst {
			isFirst = false
		} else {
			s.Out(seperator)
		}
		s.Visit(expr)
	}
	s.Out(close)
}

func (s *ExpressionStringBuilder) VisitInvocation(expr *InvocationExpression) Expression {
	s.Out("Invoke(")
	s.Visit(expr.Expression())
	sep := ", "

	for _, arg := range expr.Arguments() {
		s.Out(sep)
		s.Visit(arg)
	}
	s.Out(")")
	return expr
}

func (s *ExpressionStringBuilder) VisitBinary(expr BinaryExpression) Expression {
	op := ""
	switch expr.Type() {
	case Binary(add):
		op = "+"
	case Binary(divide):
		op = "/"
	case Binary(modulo):
		op = "%"
	case Binary(multiply):
		op = "*"
	case Binary(power):
		op = "^"
	case Binary(subtract):
		op = "-"
	case Binary(equal):
		op = "="
	case Binary(notEqual):
		op = "<>"
	case Binary(lessThan):
		op = "<"
	case Binary(lessThanOrEqual):
		op = "<="
	case Binary(greaterThan):
		op = ">"
	case Binary(greaterThanOrEqual):
		op = ">="
	case Binary(isNil):
		op = "IS NULL"
	case Binary(isNotNil):
		op = "IS NOT NULL"
	case Binary(and):
		if expr.Kind() == reflect.Bool {
			op = "And"
		} else {
			op = "&"
		}
	case Binary(or):
		if expr.Kind() == reflect.Bool {
			op = "Or"
		} else {
			op = "|"
		}
	case Binary(xor):
		if expr.Kind() == reflect.Bool {
			op = "Xor"
		} else {
			op = "^"
		}
	default:
		panic(InvalidBinaryOperations)
	}

	s.Out("(")
	s.Visit(expr.GetLeft())
	s.Out(" ")
	s.Out(op)
	s.Out(" ")
	s.Visit(expr.GetRight())
	s.Out(")")

	return expr
}

func (s *ExpressionStringBuilder) Visit(expr Expression) Expression {
	return baseVisit(s, expr)
}

func (s *ExpressionStringBuilder) VisitExtension(expr Expression) Expression {
	s.Out(expr.String())
	return expr
}

func (s *ExpressionStringBuilder) VisitParameter(expr *ParameterExpression) Expression {
	if expr.GetName() == emptyString {
		id := s.GetParamId(expr)
		s.Out("Param_" + strconv.Itoa(id))
	} else {
		s.Out(expr.GetName())
	}
	return expr
}

func (s *ExpressionStringBuilder) VisitConstant(expr *ConstantExpression) Expression {
	if expr.GetValue() != nil {
		sValue := fmt.Sprint(expr.GetValue())
		if _, ok := expr.GetValue().(string); ok {
			s.Out("\"")
			s.Out(sValue)
			s.Out("\"")
		} else {
			s.Out(sValue)
		}
	} else {
		s.Out("nil")
	}

	return expr
}

func (s *ExpressionStringBuilder) VisitConditional(expr *ConditionalExpression) Expression {
	s.Out("IF(")
	s.Visit(expr.GetTest())
	s.Out(", ")
	s.Visit(expr.GetIfTrue())
	s.Out(", ")
	s.Visit(expr.GetIfFalse())
	s.Out(")")
	return expr
}

func (s *ExpressionStringBuilder) AddLabel(label LabelTarget) {
	if _, ok := s.ids[label]; !ok {
		s.ids[label] = len(s.ids)
	}
}
func (s *ExpressionStringBuilder) GetLabelId(label LabelTarget) int {
	s.AddLabel(label)
	return s.ids[label]
}

func (s *ExpressionStringBuilder) AddParam(param *ParameterExpression) {
	if _, ok := s.ids[param]; !ok {
		s.ids[param] = len(s.ids)
	}
}

func (s *ExpressionStringBuilder) GetParamId(param *ParameterExpression) int {
	s.AddParam(param)
	return s.ids[param]
}

func (s *ExpressionStringBuilder) Out(text string) {
	s.output = s.output + text
}

func (s ExpressionStringBuilder) String() string {
	return s.output
}

func ExpressionToString(expr Expression) string {
	builder := NewExpressionStringBuilder()
	builder.Visit(expr)
	return builder.String()
}
