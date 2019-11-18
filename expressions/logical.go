package expressions

// Logical operators
type Logical Binary

const (
	// equal equality
	equal Logical = Logical(xor) + iota + 1// =
	// notEqual inequality
	notEqual // <>
	// lessThan less than
	lessThan // <
	// lessThanOrEqual less than or equal to
	lessThanOrEqual // <=
	// greaterThan greater than
	greaterThan // >
	// greaterThanOrEqual greater than or equal to
	greaterThanOrEqual // >=
	// isNil used of IS NULL
	isNil
	// isNotNil used for IS NULL
	isNotNil
)

// Compare function is a abstract comparison which converts the interfaces to the same type
func Compare(comparison Logical, x interface{}, y interface{}) bool {
	if comparison == isNil {
		return x == nil
	} else if comparison == isNotNil {
		return x != nil

	}
	x64, xOk := toFloat(x)
	y64, yOk := toFloat(y)

	if xOk && yOk {
		return comparisonFloat64(comparison, x64, y64)
	}

	xs, sxok := x.(string)
	ys, syok := y.(string)

	if sxok && syok {
		switch comparison {
		case equal:
			return xs == ys
		case notEqual:
			return xs != ys
		}
	}

	xb, bxok := x.(bool)
	yb, byok := y.(bool)

	if bxok && byok {
		switch comparison {
		case equal:
			return xb == yb
		case notEqual:
			return xb != yb
		}
	}

	return false
}

func toFloat(value interface{}) (float64, bool) {
	switch i := value.(type) {
	case float64:
		return i, true
	case int:
		return float64(i), true
	default:
		return 0, false
	}
}

func comparisonFloat64(comparison Logical, x float64, y float64) bool {
	switch comparison {
	case equal:
		return x == y
	case notEqual:
		return x != y
	case lessThan:
		return x < y
	case lessThanOrEqual:
		return x <= y
	case greaterThan:
		return x > y
	case greaterThanOrEqual:
		return x >= y
	default:
		return false
	}
}
