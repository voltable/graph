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


func (s *ExpressionStringBuilder) VisitLambda(expr *LambdaExpression) (Expression, error) {
	if len(expr.parameters) == 1 {
		_, err := s.Visit(expr.parameters[0])
		if err != nil {
			return nil, err
		}

	} else {
		err := s.VisitExpressions("(", expr.parameters, ")", ",")
		if err != nil {
			return nil, err
		}
	}
	s.Out(" => ")
	_, err := s.Visit(expr.body)
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (s *ExpressionStringBuilder) VisitExpressions(open string, expressions []*ParameterExpression, close, seperator string) error {
	s.Out(open)
	isFirst := true
	for _, expr := range expressions {
		if isFirst {
			isFirst = false
		} else {
			s.Out(seperator)
		}
		_, err := s.Visit(expr)
		if err != nil {
			return err
		}
	}
	s.Out(close)
	return nil
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
	_, err:= s.Visit(expr.GetLeft())
	if err != nil {
		return nil, err
	}

	s.Out(" ")
	s.Out(op)
	s.Out(" ")

	_, err = s.Visit(expr.GetRight())
	if err != nil {
		return nil, err
	}

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
	s.Out("IF(")

	_, err := s.Visit(expr.GetTest())
	if err != nil {
		return nil, err
	}

	s.Out(", ")

	_, err = s.Visit(expr.GetIfTrue())
	if err != nil {
		return nil, err
	}

	s.Out(", ")

	_, err = s.Visit(expr.GetIfFalse())
	if err != nil {
		return nil, err
	}

	s.Out(")")
	return expr, nil
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

