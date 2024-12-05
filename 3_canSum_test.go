package dynamic

import (
	"fmt"
	"testing"
)

var canSumTestCases = []struct {
	target int
	arr    []int
	res    bool
}{
	{7, []int{2, 3}, true},
	{7, []int{5, 3, 4, 7}, true},
	{7, []int{2, 4}, false},
	{8, []int{2, 3, 5}, true},
	{300, []int{7, 14}, false},
}

func Test_CanSum_recur(t *testing.T) {
	for i, testCase := range canSumTestCases {
		res := canSum_recur(testCase.target, testCase.arr)
		if res != testCase.res {
			t.Errorf("test %d failed: value [%t], should be [%t]", i, res, testCase.res)
		}
	}
}

func Test_CanSum_memo(t *testing.T) {
	for i, testCase := range canSumTestCases {
		res := canSum_memo(testCase.target, testCase.arr)
		if res != testCase.res {
			t.Errorf("test %d failed: value [%t], should be [%t]", i, res, testCase.res)
		}
	}
}

func Test_CanSum_tab(t *testing.T) {
	for i, testCase := range canSumTestCases {
		res := canSum_tab(testCase.target, testCase.arr)
		if res != testCase.res {
			t.Errorf("test %d failed: value [%t], should be [%t]", i, res, testCase.res)
		}
	}
}

func Benchmark_canSum_recur(b *testing.B) {
	for _, testCase := range canSumTestCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canSum_recur(testCase.target, testCase.arr)
			}
		})
	}
}

func Benchmark_canSum_memo(b *testing.B) {
	for _, testCase := range canSumTestCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canSum_memo(testCase.target, testCase.arr)
			}
		})
	}
}

func Benchmark_canSum_tab(b *testing.B) {
	for _, testCase := range canSumTestCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canSum_tab(testCase.target, testCase.arr)
			}
		})
	}
}
