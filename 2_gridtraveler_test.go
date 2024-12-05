package dynamic

import (
	"fmt"
	"testing"
)

func Test_gridTraveler(t *testing.T) {
	if gridTraveler_recur(3, 3) != 6 {
		t.Error(`should be 6`)
	}
	x := gridTraveler_recur(4, 3)
	if x != 10 {
		t.Errorf("value [%d], should be [10]", x)
	}
}

func Test_gridTraveler_memo(t *testing.T) {
	if gridTraveler_memo(3, 3) != 6 {
		t.Error(`should be 6`)
	}
	x := gridTraveler_memo(4, 3)
	if x != 10 {
		t.Errorf("value [%d], should be [10]", x)
	}
}

func Test_gridTraveler_memo_optim(t *testing.T) {
	if gridTraveler_memo_optim(3, 3) != 6 {
		t.Error(`should be 6`)
	}
	x := gridTraveler_memo_optim(4, 3)
	if x != 10 {
		t.Errorf("value [%d], should be [10]", x)
	}
}

func Test_gridTraveler_tab(t *testing.T) {
	if gridTraveler_tab(3, 3) != 6 {
		t.Error(`should be 6`)
	}
	x := gridTraveler_tab(4, 3)
	if x != 10 {
		t.Errorf("value [%d], should be [10]", x)
	}
}

func Test_gridTraveler_tab_optim(t * testing.T) {
	if gridTraveler_tab_optim(3, 3) != 6 {
		t.Error(`should be 6`)
	}
	x := gridTraveler_tab_optim(4, 3)
	if x != 10 {
		t.Errorf("value [%d], should be [10]", x)
	}
}

// benchmarks

func Benchmark_gridTraveler_recur(b *testing.B) {
	testCases := [][]int{{2, 3}, {3, 3}, {4, 4}, {18, 18}}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("array %d x %d", testCase[0], testCase[1]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gridTraveler_recur(testCase[0], testCase[1])
			}
		})
	}
}

func Benchmark_gridTraveler_memo(b *testing.B) {
	testCases := [][]int{{2, 3}, {3, 3}, {4, 4}, {18, 18}, {50, 50}}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("array %d x %d", testCase[0], testCase[1]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gridTraveler_memo(testCase[0], testCase[1])
			}
		})
	}
}

// 
func Benchmark_gridTraveler_memo_optim(b *testing.B) {
	testCases := [][]int{{2, 3}, {3, 3}, {4, 4}, {18, 18}, {50, 50}}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("array %d x %d", testCase[0], testCase[1]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gridTraveler_memo_optim(testCase[0], testCase[1])
			}
		})
	}
}

func Benchmark_gridTraveler_tab(b *testing.B) {
	testCases := [][]int{{2, 3}, {3, 3}, {4, 4}, {18, 18}, {50, 50}}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("array %d x %d", testCase[0], testCase[1]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gridTraveler_tab(testCase[0], testCase[1])
			}
		})
	}
}

func Benchmark_gridTraveler_tab_optim(b *testing.B) {
	testCases := [][]int{{2, 3}, {3, 3}, {4, 4}, {18, 18}, {50, 50}}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("array %d x %d", testCase[0], testCase[1]), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gridTraveler_tab_optim(testCase[0], testCase[1])
			}
		})
	}
}
