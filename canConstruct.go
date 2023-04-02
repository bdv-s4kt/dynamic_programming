package dynamic

import "strings"

/*
Write a function `canConstruct(target, wordBank)` that accepts
a target string and an array of strings.

The function should return a boolean indicating whether or not
the `target` can be constructed by concatenating elements
of the `wordBank` array.

You may reuse elements of `wordBank` as many times as needed.
*/

func canConstruct(target string, wordBank []string) bool {
	if target == `` {
		return true
	}

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) && canConstruct(target[len(word):], wordBank) {
			return true
		}
	}

	return false
}

func canConstruct_memo(target string, wordBank []string) bool {
	memo := map[string]bool{}

	var f func(string, []string) bool

	f = func(target string, worldBank []string) bool {
		if target == `` {
			return true
		}
		res, ok := memo[target]
		if ok {
			return res
		}

		for _, word := range worldBank {
			if strings.HasPrefix(target, word) && f(target[len(word):], wordBank) {
				memo[target] = true
				return true
			}
		}

		memo[target] = false
		return false
	}

	return f(target, wordBank)
}

func canConstruct_tab(target string, wordBank []string) bool {
	tab := make([]bool, len(target)+1)
	tab[0] = true
	for i := 0; i < len(tab); i++ {
		if !tab[i] {
			continue
		}

		for _, word := range wordBank {
			if i+len(word) <= len(target) && word == target[i:i+len(word)] {
				tab[i+len(word)] = true
			}
		}
	}
	return tab[len(tab)-1]
}
