package dynamic

import (
	"fmt"
	"testing"
)

func Test_CanSum(t *testing.T) {
	testCases := []struct {
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

	for i, testCase := range testCases {
		res := canSum(testCase.target, testCase.arr)
		if res != testCase.res {
			t.Errorf("test %d failed: value [%t], should be [%t]", i, res, testCase.res)
		}
	}
}

func Test_CanSum_memo(t *testing.T) {
	testCases := []struct {
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

	for i, testCase := range testCases {
		res := canSum_memo(testCase.target, testCase.arr)
		if res != testCase.res {
			t.Errorf("test %d failed: value [%t], should be [%t]", i, res, testCase.res)
		}
	}
}

func Test_CanSum_tab(t *testing.T) {
	testCases := []struct {
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

	for i, testCase := range testCases {
		res := canSum_tab(testCase.target, testCase.arr)
		if res != testCase.res {
			t.Errorf("test %d failed: value [%t], should be [%t]", i, res, testCase.res)
		}
	}
}

func Benchmark_canSum(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{2, 3}},
		{7, []int{5, 3, 4, 7}},
		{7, []int{2, 4}},
		{8, []int{2, 3, 5}},
		{300, []int{7, 14}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canSum(testCase.target, testCase.arr)
			}
		})
	}
}

func Benchmark_canSum_memo(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{2, 3}},
		{7, []int{5, 3, 4, 7}},
		{7, []int{2, 4}},
		{8, []int{2, 3, 5}},
		{300, []int{7, 14}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canSum_memo(testCase.target, testCase.arr)
			}
		})
	}
}

func Benchmark_canSum_tab(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{2, 3}},
		{7, []int{5, 3, 4, 7}},
		{7, []int{2, 4}},
		{8, []int{2, 3, 5}},
		{300, []int{7, 14}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canSum_tab(testCase.target, testCase.arr)
			}
		})
	}
}
