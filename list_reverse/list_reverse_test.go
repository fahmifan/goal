package stats

import "testing"

func TestMode(t *testing.T) {
	type testCase struct {
		Nums     []int
		Expected int
	}
	tableTest := []testCase{
		{
			Nums:     []int{1, 2, 3, 8, 9, 3, 2, 1},
			Expected: 3,
		},
		{
			Nums:     []int{1, 2, 1, 2, 2, 1},
			Expected: 2,
		},
		{
			Nums:     []int{7, 1, 2, 9, 7, 2, 1},
			Expected: 2,
		},
		{
			Nums:     []int{1, 1, 2, 9, 7, 2, 1, 1},
			Expected: 2,
		},
		{
			Nums:     []int{4, 1, 2, 7},
			Expected: 0,
		},
	}

	for _, ts := range tableTest {
		assert(t, ts.Expected, Run(ts.Nums))
	}
}

func assert(t *testing.T, want, got int) {
	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}
}

func TestReverse(t *testing.T) {
	ok := checkReverse([]int{1, 2, 3}, []int{3, 2, 1})
	if !ok {
		t.FailNow()
	}
}
