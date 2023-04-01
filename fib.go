package dynamic

func fib(num int) int {
	if num <= 2 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

func fib_memo(num int) int {
	memo := map[int]int{}

	var fib func(int) int
	fib = func(n int) int {
		if n <= 2 {
			return 1
		}
		res, ok := memo[n]
		if ok {
			return res
		}
		res = fib(n-1) + fib(n-2)
		memo[n] = res
		return res
	}

	return fib(num)
}

func fib_tab(num int) int {
	if num <= 2 {
		return 1
	}

	a, b := 1, 1
	for i := 0; i < num-2; i++ {
		a, b = b, a+b
	}
	return b
}
