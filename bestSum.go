package dynamic

/*
Write a function `bestSum(targetSum, numbers)` that takes in
a targetSum and an array of numbers as arguments.

The function should an array containing the shortest combination
of numbers that add up to exactly to targetSum.

If there is a tie for the shortest combination, you may return
any one of the shortest combinations.
*/

func bestSum(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	var best []int

	for _, num := range nums {
		res := bestSum(target-num, nums)
		res = append(res, num)
		if best == nil || (res != nil && len(res) < len(best)) {
			best = res
		}
	}

	return best
}

func bestSum_memo(t int, n []int) []int {
	memo := map[int][]int{}

	var f func(int, []int) []int

	f = func(target int, nums []int) []int {
		if target == 0 {
			return []int{}
		}
		if target < 0 {
			return nil
		}
		res, ok := memo[target]
		if ok {
			return res
		}

		var best []int

		for _, num := range nums {
			res := f(target-num, nums)
			res = append(res, num)
			if best == nil || (res != nil && len(res) < len(best)) {
				best = res
			}
		}

		memo[target] = best
		return best
	}

	return f(t, n)
}

func bestSum_tab(target int, nums []int) []int {
	tab := make([][]int, target+1)
	tab[0] = []int{}

	for i := 0; i < len(tab); i++ {
		if tab[i] == nil {
			continue
		}
		for _, num := range nums {
			index := i + num
			if index < len(tab) && (tab[index] == nil || len(tab[i])+1 < len(tab[index])) {
				tab[index] = append(tab[i], num)
			}
		}
	}
	return tab[len(tab)-1]
}
