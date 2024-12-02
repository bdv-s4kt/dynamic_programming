package dynamic

/*
Write a function `fib(n)` that takes in a number as an argument.
The function should return the n-th number of the Fibonacci sequence.

The 1st and 2nd number of the sequence is 1.
To generate the next number of the sequence, we sum the previous two.
*/

// Recursive solution, top-down approach
// Time: O(2^n), space: O(n)
func fib_recur(num int) int {
	if num <= 2 {
		return 1
	}
	return fib_recur(num-1) + fib_recur(num-2)
}

// Recursive solution, using memoization
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

// Tabulation solution, bottom-up approach
// Time: O(n), Space: O(n)
func fib_tab(num int) int {
	tab := make([]int, num)
	tab[0], tab[1] = 1, 1

	for i := 2; i < num; i++ {
		tab[i] = tab[i-1] + tab[i-2]
	}
	return tab[num-1]
}

// Tabulation with optimization solution, bottom-up approach,
// but don't store all the values, only those which are necessary for next iterations
// Time: O(n), Space: O(1)
func fib_tab_optim(num int) int {
	if num <= 2 {
		return 1
	}

	a, b := 1, 1
	for i := 0; i < num-2; i++ {
		a, b = b, a+b
	}
	return b
}
