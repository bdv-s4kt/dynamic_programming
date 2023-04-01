package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_howSum(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := howSum(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Test_howSum_memo(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := howSum_memo(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Test_howSum_tab(t *testing.T) {
	testCases := []struct {
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

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("test %d %v", testCase.target, testCase.nums), func(t *testing.T) {
			res := howSum_tab(testCase.target, testCase.nums)
			if !reflect.DeepEqual(res, testCase.res) {
				t.Errorf("value %v should be %v", res, testCase.res)
			}
		})
	}
}

func Benchmark_howSum(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{5, 3, 4}},
		{7, []int{2, 3}},
		{7, []int{5, 3, 4, 7}},
		{7, []int{2, 4}},
		{8, []int{2, 3, 5}},
		{300, []int{7, 14}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				howSum(testCase.target, testCase.arr)
			}
		})
	}
}

func Benchmark_howSum_memo(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{5, 3, 4}},
		{7, []int{2, 3}},
		{7, []int{5, 3, 4, 7}},
		{7, []int{2, 4}},
		{8, []int{2, 3, 5}},
		{300, []int{7, 14}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				howSum_memo(testCase.target, testCase.arr)
			}
		})
	}
}

func Benchmark_howSum_tab(b *testing.B) {
	testCases := []struct {
		target int
		arr    []int
	}{
		{7, []int{5, 3, 4}},
		{7, []int{2, 3}},
		{7, []int{5, 3, 4, 7}},
		{7, []int{2, 4}},
		{8, []int{2, 3, 5}},
		{300, []int{7, 14}},
	}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%d, %v", testCase.target, testCase.arr), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				howSum_tab(testCase.target, testCase.arr)
			}
		})
	}
}
