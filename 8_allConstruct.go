package dynamic

import (
	"strings"
)

/*
Write a function `allConstruct(target, wordBank)` that accepts
a target string and an array of strings.

The function should return a 2D array containing all of the ways
that the `target` can be constructed by concatenating elements
of the `wordBank` array. Each element of the 2D array should
represent one combination that constructs the `target`.

You may reuse elements of `wordBank` as many times as needed.
*/

func allConstruct(target string, wordBank []string) [][]string {
	if target == `` {
		return [][]string{{}}
	}

	res := [][]string{}

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			for _, newConstuct := range allConstruct(target[len(word):], wordBank) {
				newConstuct = append(newConstuct, word)
				res = append(res, newConstuct)
			}
		}
	}

	return res
}

func allConstruct_memo(target string, wordBank []string) [][]string {
	memo := map[string][][]string{}

	var f func(string, []string) [][]string

	f = func(target string, wordBank []string) [][]string {
		if target == `` {
			return [][]string{{}}
		}

		res, ok := memo[target]
		if ok {
			return res
		}

		res = [][]string{}

		for _, word := range wordBank {
			if strings.HasPrefix(target, word) {
				for _, newConstuct := range f(target[len(word):], wordBank) {
					newConstuct = append(newConstuct, word)
					res = append(res, newConstuct)
				}
			}
		}

		memo[target] = res
		return res
	}

	return f(target, wordBank)
}

func allConstruct_tab(target string, wordBank []string) [][]string {
	tab := make([][][]string, len(target)+1)
	tab[0] = [][]string{{}}

	for i := 0; i < len(tab); i++ {
		if tab[i] == nil {
			continue
		}

		for _, word := range wordBank {
			if i+len(word) <= len(target) && word == target[i:i+len(word)] {
				newComb := [][]string{}
				for _, comb := range tab[i] {
					sl := append(comb, word)
					newComb = append(newComb, sl)
				}
				tab[i+len(word)] = append(tab[i+len(word)], newComb...)
			}
		}
	}

	return tab[len(tab)-1]
}
