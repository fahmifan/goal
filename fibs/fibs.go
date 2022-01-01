package fibs

type memoizer map[int]int

func Fib(n int, memo memoizer) int {
	if memo == nil {
		memo = make(memoizer)
	}

	if val, ok := memo[n]; ok {
		return val
	}

	if n <= 2 {
		return 1
	}

	memo[n] = Fib(n-1, memo) + Fib(n-2, memo)
	return memo[n]
}
