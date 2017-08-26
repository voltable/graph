package comparisons_test

import (
	"math"
	"testing"

	"github.com/RossMerr/Caudex.Graph/comparisons"
)

func Test_Compare(t *testing.T) {

	var tests = []struct {
		comparison comparisons.Comparison
		x          interface{}
		y          interface{}
		result     bool
	}{
		/// LT
		{
			comparison: comparisons.LT,
			x:          math.SmallestNonzeroFloat64,
			y:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: comparisons.LT,
			x:          math.SmallestNonzeroFloat32,
			y:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: comparisons.LT,
			x:          math.MinInt32,
			y:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: comparisons.LT,
			x:          math.MinInt16,
			y:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: comparisons.LT,
			x:          math.MinInt8,
			y:          math.MaxInt8,
			result:     true,
		},

		///// LTE

		{
			comparison: comparisons.LTE,
			x:          math.SmallestNonzeroFloat64,
			y:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: comparisons.LTE,
			x:          math.SmallestNonzeroFloat32,
			y:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: comparisons.LTE,
			x:          math.MinInt32,
			y:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: comparisons.LTE,
			x:          math.MinInt16,
			y:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: comparisons.LTE,
			x:          math.MinInt8,
			y:          math.MaxInt8,
			result:     true,
		},

		///// GT

		{
			comparison: comparisons.GT,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: comparisons.GT,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: comparisons.GT,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: comparisons.GT,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: comparisons.GT,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// GTE

		{
			comparison: comparisons.GTE,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: comparisons.GTE,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: comparisons.GTE,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: comparisons.GTE,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: comparisons.GTE,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// NEQ

		{
			comparison: comparisons.NEQ,
			y:          math.SmallestNonzeroFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: comparisons.NEQ,
			y:          math.SmallestNonzeroFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: comparisons.NEQ,
			y:          math.MinInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: comparisons.NEQ,
			y:          math.MinInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: comparisons.NEQ,
			y:          math.MinInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		///// EQ

		{
			comparison: comparisons.EQ,
			y:          math.MaxFloat64,
			x:          math.MaxFloat64,

			result: true,
		},
		{
			comparison: comparisons.EQ,
			y:          math.MaxFloat32,
			x:          math.MaxFloat32,
			result:     true,
		},
		{
			comparison: comparisons.EQ,
			y:          math.MaxInt32,
			x:          math.MaxInt32,
			result:     true,
		},
		{
			comparison: comparisons.EQ,
			y:          math.MaxInt16,
			x:          math.MaxInt16,
			result:     true,
		},
		{
			comparison: comparisons.EQ,
			y:          math.MaxInt8,
			x:          math.MaxInt8,
			result:     true,
		},

		// string

		{
			comparison: comparisons.EQ,
			y:          "OK",
			x:          "OK",
			result:     true,
		},

		{
			comparison: comparisons.EQ,
			y:          "FOO",
			x:          "BAR",
			result:     false,
		},

		{
			comparison: comparisons.NEQ,
			y:          "OK",
			x:          "OK",
			result:     false,
		},

		{
			comparison: comparisons.NEQ,
			y:          "FOO",
			x:          "BAR",
			result:     true,
		},

		// bool

		{
			comparison: comparisons.EQ,
			y:          true,
			x:          true,
			result:     true,
		},

		{
			comparison: comparisons.EQ,
			y:          true,
			x:          false,
			result:     false,
		},

		{
			comparison: comparisons.NEQ,
			y:          true,
			x:          true,
			result:     false,
		},

		{
			comparison: comparisons.NEQ,
			y:          true,
			x:          false,
			result:     true,
		},
	}

	for i, tt := range tests {
		result := comparisons.Compare(tt.comparison, tt.x, tt.y)
		if result != tt.result {
			t.Errorf("%d. comparison mismatch:\n  exp=%t\n  got=%t\n\n", i, tt.result, result)
		}
	}
}
