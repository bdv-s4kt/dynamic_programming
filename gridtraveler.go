package dynamic

import "fmt"

/*
Say that you are a traveler on a 2D grid. You begin
in the top-left corner and your goal is to travel to
the bottom-right corner. You may only move down or right.

In how many ways can you travel to the goal on a grid with
dimensions m * n?

Write a function `gridTraveler(m, n)` that calculate this.
*/

func gridTraveler(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	return gridTraveler(m-1, n) + gridTraveler(m, n-1)
}

func gridTraveler_memo(mm, nn int) int {
	memo := map[string]int{}

	var f func(int, int) int

	f = func(m, n int) int {
		if m == 1 || n == 1 {
			return 1
		}
		ind := fmt.Sprintf("%d:%d", m, n)
		res, ok := memo[ind]
		if ok {
			return res
		}

		res = f(m-1, n) + f(m, n-1)
		memo[ind] = res
		return res
	}

	return f(mm, nn)
}

func gridTraveler_tab(m, n int) int {
	tab := make([][]int, m)
	for i := range tab {
		tab[i] = make([]int, n)
	}
	tab[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var left, top int
			if i-1 >= 0 {
				top = tab[i-1][j]
			}
			if j-1 >= 0 {
				left = tab[i][j-1]
			}
			tab[i][j] = top + left
			if tab[i][j] == 0 {
				tab[i][j] = 1
			}
		}
	}
	return tab[m-1][n-1]
}
