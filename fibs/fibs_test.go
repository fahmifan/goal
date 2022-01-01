package fibs

import "testing"

func TestFibs(t *testing.T) {
	type testTable struct {
		NFib     int
		Expected int
	}
	table := []testTable{
		{
			NFib:     3,
			Expected: 2,
		},
		{
			NFib:     4,
			Expected: 3,
		},
		{
			NFib:     5,
			Expected: 5,
		},
		{
			NFib:     6,
			Expected: 8,
		},
		{
			NFib:     7,
			Expected: 13,
		},
		{
			NFib:     50,
			Expected: 12586269025,
		},
	}
	for _, tab := range table {
		res := Fib(tab.NFib, nil)
		assert(t, tab.Expected, res)
	}
}

func BenchmarkFibs(b *testing.B) {
	b.Run("Fibs", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Fib(i, nil)
		}
	})
}

func assert(t *testing.T, want, got int) {
	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}
}
