package autopilot_test

import (
	"testing"

	"github.com/John-Tonny/lnd/autopilot"
	vclutil "github.com/John-Tonny/vclsuite_vclutil"
)

// TestMedian tests the Median method.
func TestMedian(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		values []vclutil.Amount
		median vclutil.Amount
	}{
		{
			values: []vclutil.Amount{},
			median: 0,
		},
		{
			values: []vclutil.Amount{10},
			median: 10,
		},
		{
			values: []vclutil.Amount{10, 20},
			median: 15,
		},
		{
			values: []vclutil.Amount{10, 20, 30},
			median: 20,
		},
		{
			values: []vclutil.Amount{30, 10, 20},
			median: 20,
		},
		{
			values: []vclutil.Amount{10, 10, 10, 10, 5000000},
			median: 10,
		},
	}

	for _, test := range testCases {
		res := autopilot.Median(test.values)
		if res != test.median {
			t.Fatalf("expected median %v, got %v", test.median, res)
		}
	}
}
