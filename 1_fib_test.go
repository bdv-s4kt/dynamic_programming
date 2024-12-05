package dynamic

import (
	"fmt"
	"testing"
)

func Test_fib_recur(t *testing.T) {
	if fib_recur(8) != 21 {
		t.Error(`fib_recur(8) != 21`)
	}
}

func Test_fib_memo(t *testing.T) {
	if fib_memo(8) != 21 {
		t.Error(`fib_memo(8) != 21`)
	}
}

func Test_fib_tab(t *testing.T) {
	if fib_tab(8) != 21 {
		t.Error(`fib_tab(8) != 21`)
	}
}

func Test_fib_tab_optim(t *testing.T) {
	if fib_tab_optim(8) != 21 {
		t.Error(`fib_tab_optim(8) != 21`)
	}
}

func Benchmark_fib_recur(b *testing.B) {
	testCases := []int{10, 20, 30, 40, 45}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("fib_recur(%d)", testCase), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fib_recur(testCase)
			}
		})
	}
}
func Benchmark_fib_memo(b *testing.B) {
	testCases := []int{10, 20, 30, 40, 45, 500, 5000, 50000}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("fib_memo(%d)", testCase), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fib_memo(testCase)
			}
		})
	}
}

func Benchmark_fib_tab(b *testing.B) {
	testCases := []int{10, 20, 30, 40, 45, 500, 5000, 50000}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("fib_tab(%d)", testCase), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fib_tab(testCase)
			}
		})
	}
}

func Benchmark_fib_tab_optim(b *testing.B) {
	testCases := []int{10, 20, 30, 40, 45, 500, 5000, 50000}
	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("fib_tab_optim(%d)", testCase), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fib_tab_optim(testCase)
			}
		})
	}
}
