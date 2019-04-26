package expressions_test

import (
	"math"
	"testing"

	"github.com/voltable/graph/expressions"
)

func Test_Compare(t *testing.T) {

	var tests = []struct {
		comparison expressions.Comparison
		x          interface{}
		y          interface{}
		result     bool
	}{
		/// LT
		{
			comparison: expressions.LT,
			x:          math.SmallestNonzeroFloat64,
			y:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: expressions.LT,
			x:          math.SmallestNonzeroFloat32,
			y:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: expressions.LT,
			x:          math.MinInt32,
			y:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: expressions.LT,
			x:          math.MinInt16,
			y:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: expressions.LT,
			x:          math.MinInt8,
			y:          math.MaxInt8,
			result:     true,
		},

		///// LTE

		{
			comparison: expressions.LTE,
			x:          math.SmallestNonzeroFloat64,
			y:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: expressions.LTE,
			x:          math.SmallestNonzeroFloat32,
			y:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: expressions.LTE,
			x:          math.MinInt32,
			y:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: expressions.LTE,
			x:          math.MinInt16,
			y:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: expressions.LTE,
			x:          math.MinInt8,
			y:          math.MaxInt8,
			result:     true,
		},

		///// GT

		{
			comparison: expressions.GT,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: expressions.GT,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: expressions.GT,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: expressions.GT,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: expressions.GT,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// GTE

		{
			comparison: expressions.GTE,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: expressions.GTE,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: expressions.GTE,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: expressions.GTE,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: expressions.GTE,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// NEQ

		{
			comparison: expressions.NEQ,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: expressions.NEQ,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: expressions.NEQ,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: expressions.NEQ,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: expressions.NEQ,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// EQ

		{
			comparison: expressions.EQ,
			y:          math.MaxFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: expressions.EQ,
			y:          math.MaxFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: expressions.EQ,
			y:          math.MaxInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: expressions.EQ,
			y:          math.MaxInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: expressions.EQ,
			y:          math.MaxInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		// string

		{
			comparison: expressions.EQ,
			y:          "OK",
			x:          "OK",
			result:     true,
		},

		{
			comparison: expressions.EQ,
			y:          "FOO",
			x:          "BAR",
			result:     false,
		},

		{
			comparison: expressions.NEQ,
			y:          "OK",
			x:          "OK",
			result:     false,
		},

		{
			comparison: expressions.NEQ,
			y:          "FOO",
			x:          "BAR",
			result:     true,
		},

		// bool

		{
			comparison: expressions.EQ,
			y:          true,
			x:          true,
			result:     true,
		},

		{
			comparison: expressions.EQ,
			y:          true,
			x:          false,
			result:     false,
		},

		{
			comparison: expressions.NEQ,
			y:          true,
			x:          true,
			result:     false,
		},

		{
			comparison: expressions.NEQ,
			y:          true,
			x:          false,
			result:     true,
		},
	}

	for i, tt := range tests {
		result := expressions.Compare(tt.comparison, tt.x, tt.y)
		if result != tt.result {
			t.Errorf("%d. comparison mismatch:\n  exp=%t\n  got=%t\n\n", i, tt.result, result)
		}
	}
}
