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

func howSum(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for _, num := range nums {
		res := howSum(target-num, nums)
		if res != nil {
			res = append(res, num)
			return res
		}
	}
	return nil
}

func howSum_memo(t int, n []int) []int {
	memo := map[int][]int{}

	var f func(int, []int) []int
	f = func(target int, nums []int) []int {
		if target == 0 {
			return []int{}
		}
		if target < 0 {
			return nil
		}
		val, ok := memo[target]
		if ok {
			return val
		}

		for _, num := range nums {
			res := f(target-num, nums)
			if res != nil {
				res = append(res, num)
				memo[target] = res
				return res
			}
		}
		memo[target] = nil
		return nil
	}
	return f(t, n)
}

func howSum_tab(target int, nums []int) []int {
	tab := make([][]int, target+1)
	tab[0] = []int{}
	for i := 0; i < len(tab); i++ {
		if tab[i] == nil {
			continue
		}
		for _, num := range nums {
			if i+num < len(tab) {
				res := append(tab[i][:], num)
				tab[i+num] = res
			}
		}
	}
	return tab[len(tab)-1]
}
