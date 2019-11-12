package expressions

import (
	"math"
	"testing"
)

func Test_Compare(t *testing.T) {

	var tests = []struct {
		comparison Logical
		x          interface{}
		y          interface{}
		result     bool
	}{
		/// lessThan
		{
			comparison: lessThan,
			x:          math.SmallestNonzeroFloat64,
			y:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: lessThan,
			x:          math.SmallestNonzeroFloat32,
			y:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: lessThan,
			x:          math.MinInt32,
			y:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: lessThan,
			x:          math.MinInt16,
			y:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: lessThan,
			x:          math.MinInt8,
			y:          math.MaxInt8,
			result:     true,
		},

		///// lessThanOrEqual

		{
			comparison: lessThanOrEqual,
			x:          math.SmallestNonzeroFloat64,
			y:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: lessThanOrEqual,
			x:          math.SmallestNonzeroFloat32,
			y:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: lessThanOrEqual,
			x:          math.MinInt32,
			y:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: lessThanOrEqual,
			x:          math.MinInt16,
			y:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: lessThanOrEqual,
			x:          math.MinInt8,
			y:          math.MaxInt8,
			result:     true,
		},

		///// greaterThan

		{
			comparison: greaterThan,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: greaterThan,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: greaterThan,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: greaterThan,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: greaterThan,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// greaterThanOrEqual

		{
			comparison: greaterThanOrEqual,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: greaterThanOrEqual,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: greaterThanOrEqual,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: greaterThanOrEqual,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: greaterThanOrEqual,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// notEqual

		{
			comparison: notEqual,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: notEqual,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: notEqual,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: notEqual,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: notEqual,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// equal

		{
			comparison: equal,
			y:          math.MaxFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: equal,
			y:          math.MaxFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: equal,
			y:          math.MaxInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: equal,
			y:          math.MaxInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: equal,
			y:          math.MaxInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		// string

		{
			comparison: equal,
			y:          "OK",
			x:          "OK",
			result:     true,
		},

		{
			comparison: equal,
			y:          "FOO",
			x:          "BAR",
			result:     false,
		},

		{
			comparison: notEqual,
			y:          "OK",
			x:          "OK",
			result:     false,
		},

		{
			comparison: notEqual,
			y:          "FOO",
			x:          "BAR",
			result:     true,
		},

		// bool

		{
			comparison: equal,
			y:          true,
			x:          true,
			result:     true,
		},

		{
			comparison: equal,
			y:          true,
			x:          false,
			result:     false,
		},

		{
			comparison: notEqual,
			y:          true,
			x:          true,
			result:     false,
		},

		{
			comparison: notEqual,
			y:          true,
			x:          false,
			result:     true,
		},
	}

	for i, tt := range tests {
		result := Compare(tt.comparison, tt.x, tt.y)
		if result != tt.result {
			t.Errorf("%d. comparison mismatch:\n  exp=%t\n  got=%t\n\n", i, tt.result, result)
		}
	}
}
