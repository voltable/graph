package expressions

import (
	"fmt"
	"reflect"
	"strconv"
)

type ExpressionStringBuilder struct {
	ids map[interface{}]int
	output string
}


var _ ExpressionVisitor = (*ExpressionStringBuilder)(nil)

func NewExpressionStringBuilder() *ExpressionStringBuilder{
	return &ExpressionStringBuilder{
		ids:    make(map[interface{}]int),
		output: "",
	}
}

func (s *ExpressionStringBuilder) VisitBinary(expr BinaryExpression) (Expression, error) {
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
		return nil, InvalidBinaryOperations
	}

	s.Out("(")
	s.Visit(expr.GetLeft())
	s.Out(" ")
	s.Out(op)
	s.Out(" ")
	s.Visit(expr.GetRight())
	s.Out(")")

	return expr, nil
}

func (s *ExpressionStringBuilder) Visit(expr Expression) (Expression, error) {
	return baseVisit(s, expr)
}

func (s *ExpressionStringBuilder) VisitExtension(expr Expression) (Expression, error) {
	s.Out(expr.String())
	return expr, nil
}

func (s *ExpressionStringBuilder) VisitParameter(expr *ParameterExpression) (Expression, error) {
	if expr.GetName() == emptyString {
		id := s.GetParamId(expr)
		s.Out("Param_" + strconv.Itoa(id))
	} else {
		s.Out(expr.GetName())
	}
	return expr, nil
}

func (s *ExpressionStringBuilder) VisitConstant(expr *ConstantExpression) (Expression, error) {
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

	return expr, nil
}

func (s *ExpressionStringBuilder) VisitConditional(expr *ConditionalExpression) (Expression, error) {
	s.Out("IIF(")
	s.Visit(expr.GetTest())
	s.Out(", ")
	s.Visit(expr.GetIfTrue())
	s.Out(", ")
	s.Visit(expr.GetIfFalse())
	s.Out(")")
	return baseVisitConditional(s, expr)
}


func (s *ExpressionStringBuilder)AddLabel(label LabelTarget) {
	if _, ok := s.ids[label]; !ok {
		s.ids[label] = len(s.ids)
	}
}
func (s *ExpressionStringBuilder)GetLabelId(label LabelTarget) int {
	s.AddLabel(label)
	return s.ids[label]
}

func (s *ExpressionStringBuilder) AddParam(param *ParameterExpression) {
	if _, ok := s.ids[param]; !ok {
		s.ids[param] = len(s.ids)
	}
}

func (s *ExpressionStringBuilder)GetParamId(param *ParameterExpression) int {
	s.AddParam(param)
	return s.ids[param]
}

func (s *ExpressionStringBuilder) Out(text string) {
	s.output = s.output + text
}

func (s ExpressionStringBuilder)String() string {
	return s.output
}

func ExpressionToString(expr Expression) string {
	builder := NewExpressionStringBuilder()
	builder.Visit(expr)
	return builder.String()
}

