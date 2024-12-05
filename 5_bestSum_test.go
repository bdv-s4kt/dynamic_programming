package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

var bestSumTestCases = []struct {
	target int
	nums   []int
	res    []int
}{
	{7, []int{5, 3, 4, 7}, []int{7}},
	{8, []int{2, 3, 5}, []int{5, 3}},
	{8, []int{1, 4, 5}, []int{4, 4}},
	{50, []int{5, 25}, []int{25, 25}},
	{100, []int{1, 2, 5, 25}, []int{25, 25, 25, 25}},
}

func Test_bestSum_recur(t *testing.T) {
	for _, testCase := range bestSumTestCases[:4] {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := bestSum_recur(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Test_bestSum_memo(t *testing.T) {
	for _, testCase := range bestSumTestCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := bestSum_memo(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Test_bestSum_tab(t *testing.T) {
	for _, testCase := range bestSumTestCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := bestSum_tab(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Benchmark_bestSum(b *testing.B) {
	for _, testCase := range bestSumTestCases[:4] {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.nums), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bestSum_recur(testCase.target, testCase.nums)
			}
		})
	}
}

func Benchmark_bestSum_memo(b *testing.B) {
	for _, testCase := range bestSumTestCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.nums), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bestSum_memo(testCase.target, testCase.nums)
			}
		})
	}
}

func Benchmark_bestSum_tab(b *testing.B) {
	for _, testCase := range bestSumTestCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.nums), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bestSum_tab(testCase.target, testCase.nums)
			}
		})
	}
}
