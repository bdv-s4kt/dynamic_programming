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

var globalCount = 0

func allConstruct(target string, wordBank []string) [][]string {
	globalCount++
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

func allConstructMemo(target string, wordBank []string) [][]string {
	memo := map[string][][]string{}

	var f func(string, []string) [][]string

	f = func(t string, wb []string) [][]string {
		globalCount++
		if t == `` {
			return [][]string{{}}
		}

		res, ok := memo[t]
		if ok {
			return res
		}

		res = [][]string{}

		for _, word := range wb {
			if strings.HasPrefix(t, word) {
				for _, newConstuct := range f(t[len(word):], wb) {
					newConstuct = append(newConstuct, word)
					res = append(res, newConstuct)
				}
			}
		}

		memo[t] = res
		return res
	}

	return f(target, wordBank)
}
