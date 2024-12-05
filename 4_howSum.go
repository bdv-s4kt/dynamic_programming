package dynamic

/*
Write a function `howSum(targetSum, numbers)` that takes in
a targetSum and an array of numbers as arguments.

The function should return an array containing any combinations
of elements that add up to exactly the targetSum. If there is no
combination that adds up to the targetSum, then return null.

If there are multiple combinations possible, you may
return any single one.
*/

// m = TargetSum
// n = len(nums)
//
// Recursive solution
// Time: O(n^m), Space: O(m)
func howSum_recur(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, num := range nums {
		res := howSum_recur(target-num, nums)
		if res != nil {
			res = append(res, num)
			return res
		}
	}
	return nil
}

// Memoization
// Time: O(n*m), Space: O(m)
func howSum_memo(t int, n []int) []int {
	memo := map[int]bool{}

	var f func(int, []int) []int

	f = func(target int, nums []int) []int {
		if target == 0 {
			return []int{}
		}
		if target < 0 {
			return nil
		}
		if _, ok := memo[target]; ok {
			return nil
		}

		for _, num := range nums {
			res := f(target-num, nums)
			if res != nil {
				res = append(res, num)
				return res
			}
		}
		memo[target] = false
		return nil
	}
	return f(t, n)
}

// Tabulation
// Time: O(n*m*m), Space: O(m*m)
func howSum_tab(target int, nums []int) []int {
	tab := make([][]int, target+1)
	tab[0] = []int{}
	for i := 0; i < len(tab); i++ {
		if tab[i] == nil {
			continue
		}
		for _, num := range nums {
			if i+num < len(tab) {
				tab[i+num] = append(tab[i][:], num)
			}
		}
	}
	return tab[len(tab)-1]
}
