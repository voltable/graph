package expressions

// Comparison operators
type Comparison int

const (
	// EQ equality
	EQ Comparison = iota // =
	// NEQ inequality
	NEQ // <>
	// LT less than
	LT // <
	// LTE less than or equal to
	LTE // <=
	// GT greater than
	GT // >
	// GTE greater than or equal to
	GTE // >=
	// IS_NULL used of IS NULL
	IS_NULL
	// IS_NOT_NULL used for IS NULL
	IS_NOT_NULL
)

// Compare function is a abstract comparison which converts the interfaces to the same type
func Compare(comparison Comparison, x interface{}, y interface{}) bool {
	if comparison == IS_NULL {
		return x == nil
	} else if comparison == IS_NOT_NULL {
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
		case EQ:
			return xs == ys
		case NEQ:
			return xs != ys
		}
	}

	xb, bxok := x.(bool)
	yb, byok := y.(bool)

	if bxok && byok {
		switch comparison {
		case EQ:
			return xb == yb
		case NEQ:
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

func comparisonFloat64(comparison Comparison, x float64, y float64) bool {
	switch comparison {
	case EQ:
		return x == y
	case NEQ:
		return x != y
	case LT:
		return x < y
	case LTE:
		return x <= y
	case GT:
		return x > y
	case GTE:
		return x >= y
	default:
		return false
	}
}
