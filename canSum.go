package dynamic

func canSum(target int, arr []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for _, num := range arr {
		if canSum(target-num, arr) {
			return true
		}
	}
	return false
}

func canSum_memo(t int, a []int) bool {
	memo := map[int]bool{}
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
			res := f(target-num, arr)
			if res {
				memo[target] = true
				return true
			}
		}
		memo[target] = false
		return false
	}
	return f(t, a)
}

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
