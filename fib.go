package dynamic

/*
Write a function `fib(n)` that takes in a number as an argument.
The function should return the n-th number of the Fibonacci sequence.

The 1st and 2nd number of the sequence is 1.
To generate the next number of the sequence, we sum the previous two.
*/

// Time: O(2^n), space: O(n)
func fib(num int) int {
	if num <= 2 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

// Time: O(n), Space: O(n)
func fib_memo(num int) int {
	memo := map[int]int{}

	var fib func(int) int
	fib = func(n int) int {
		if n <= 2 {
			return 1
		}
		if _, ok := memo[n]; !ok {
			memo[n] = fib(n-1) + fib(n-2)
		}
		return memo[n]
	}

	return fib(num)
}

//TODO: add tabulation with actual table as an example

// Time: O(n), Space: O(1)
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
