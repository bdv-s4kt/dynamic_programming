package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_bestSum(t *testing.T) {
	testCases := []struct {
		target int
		nums   []int
		res    []int
	}{
		{7, []int{5, 3, 4, 7}, []int{7}},
		{8, []int{2, 3, 5}, []int{5, 3}},
		{8, []int{1, 4, 5}, []int{4, 4}},
		{50, []int{5, 25}, []int{25, 25}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := bestSum(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Test_bestSum_memo(t *testing.T) {
	testCases := []struct {
		target int
		nums   []int
		res    []int
	}{
		{7, []int{5, 3, 4, 7}, []int{7}},
		{8, []int{2, 3, 5}, []int{5, 3}},
		{8, []int{1, 4, 5}, []int{4, 4}},
		{100, []int{1, 2, 5, 25}, []int{25, 25, 25, 25}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := bestSum_memo(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Test_bestSum_tab(t *testing.T) {
	testCases := []struct {
		target int
		nums   []int
		res    []int
	}{
		{7, []int{5, 3, 4, 7}, []int{7}},
		{8, []int{2, 3, 5}, []int{3, 5}},
		{8, []int{1, 4, 5}, []int{4, 4}},
		{100, []int{1, 2, 5, 25}, []int{25, 25, 25, 25}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := bestSum_tab(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Benchmark_bestSum(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{5, 3, 4, 7}},
		{8, []int{2, 3, 5}},
		{8, []int{1, 4, 5}},
		{100, []int{1, 2, 5, 25}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bestSum(testCase.target, testCase.arr)
			}
		})
	}
}

func Benchmark_bestSum_memo(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{5, 3, 4, 7}},
		{8, []int{2, 3, 5}},
		{8, []int{1, 4, 5}},
		{100, []int{1, 2, 5, 25}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bestSum_memo(testCase.target, testCase.arr)
			}
		})
	}
}

func Benchmark_bestSum_tab(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{5, 3, 4, 7}},
		{8, []int{2, 3, 5}},
		{8, []int{1, 4, 5}},
		{100, []int{1, 2, 5, 25}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bestSum_tab(testCase.target, testCase.arr)
			}
		})
	}
}
