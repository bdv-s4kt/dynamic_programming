package dynamic

import "strings"

/*
Write a function `countConstruct(target, wordBank)` that accepts
a target string and an array of strings.

The function should return the number of ways that the `target`
can be constructed by concatenating elements of the `wordBank` array.

You may reuse elements of `wordBank` array as many times as neeeded.
*/

// Recursive solution
// m = len(target)
// n = len(wordBank)
// Time: O(n^m*m), Space: O(m^2)
func countConstruct(target string, wordBank []string) int {
	if target == `` {
		return 1
	}

	count := 0
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			count += countConstruct(target[len(word):], wordBank)
		}
	}

	return count
}

// Memoization
// Time: O(n*m^2), Space: O(m^2)
func countConstruct_memo(target string, wordBank []string) int {
	memo := map[string]int{}

	var f func(string, []string) int

	f = func(target string, wordBank []string) int {
		if target == `` {
			return 1
		}
		res, ok := memo[target]
		if ok {
			return res
		}

		count := 0
		for _, word := range wordBank {
			if strings.HasPrefix(target, word) {
				count += f(target[len(word):], wordBank)
			}
		}

		memo[target] = count
		return count
	}

	return f(target, wordBank)
}

// Tabulation
// Time: O(n*m^2), O(m)
func countConstruct_tab(target string, wordBank []string) int {
	tab := make([]int, len(target)+1)
	tab[0] = 1

	for i := 0; i < len(tab); i++ {
		if tab[i] == 0 {
			continue
		}

		for _, word := range wordBank {
			if i+len(word) <= len(target) && word == target[i:i+len(word)] {
				tab[i+len(word)] += tab[i]
			}
		}
	}

	return tab[len(tab)-1]
}
