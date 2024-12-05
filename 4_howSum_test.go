package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

var howSumTestCases = []struct {
	target int
	nums   []int
	res    []int
}{
	{7, []int{5, 3, 4}, []int{4, 3}},
	{7, []int{2, 3}, []int{3, 2, 2}},
	{7, []int{5, 3, 4, 7}, []int{4, 3}},
	{7, []int{2, 4}, nil},
	{8, []int{2, 3, 5}, []int{2, 2, 2, 2}},
	{300, []int{7, 14}, nil},
}

func Test_howSum_recur(t *testing.T) {
	for _, testCase := range howSumTestCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := howSum_recur(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Test_howSum_memo(t *testing.T) {
	for _, testCase := range howSumTestCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := howSum_memo(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Test_howSum_tab(t *testing.T) {
	for _, testCase := range howSumTestCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := howSum_tab(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Benchmark_howSum(b *testing.B) {
	for _, testCase := range howSumTestCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.nums), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				howSum_recur(testCase.target, testCase.nums)
			}
		})
	}
}

func Benchmark_howSum_memo(b *testing.B) {
	for _, testCase := range howSumTestCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.nums), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				howSum_memo(testCase.target, testCase.nums)
			}
		})
	}
}

func Benchmark_howSum_tab(b *testing.B) {
	for _, testCase := range howSumTestCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.nums), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				howSum_tab(testCase.target, testCase.nums)
			}
		})
	}
}
