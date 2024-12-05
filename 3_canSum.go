package dynamic

/*
Write a function `canSum(target, numbers)` that takes in
a targetSum and an array of numbers as arguments.

The function should return a boolean indicating whether or not it is
possible to generate the targetSum using numbers from the array.

You may use an element of the array as many times as needed.

You may assume that all input numbers are non-negative.
*/

// Recursive solution
// m = target
// n = len(arr)
// Time: O(n^m), Space: O(m)
func canSum_recur(target int, arr []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for _, num := range arr {
		if canSum_recur(target-num, arr) {
			return true
		}
	}
	return false
}

// Memoization
// Time: O(m*n), Space: O(m)
func canSum_memo(t int, a []int) bool {
	memo := make(map[int]bool, t)

	var f func(int, []int) bool

	f = func(target int, arr []int) bool {
		if target == 0 {
			return true
		}
		if target < 0 {
			return false
		}

		val, ok := memo[target]
		if ok {
			return val
		}

		for _, num := range arr {
			if f(target-num, arr) {
				return true
			}
		}

		memo[target] = false
		return false
	}

	return f(t, a)
}

// Tabulation
// Time: O(m*n), Space: O(m)
func canSum_tab(target int, arr []int) bool {
	tab := make([]bool, target+1)
	tab[0] = true

	for i := 0; i < len(tab); i++ {
		if !tab[i] {
			continue
		}
		for _, num := range arr {
			if i+num < len(tab) {
				tab[i+num] = true
			}
		}
	}
	return tab[len(tab)-1]
}
