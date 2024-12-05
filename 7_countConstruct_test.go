package dynamic

import (
	"fmt"
	"testing"
)

var (
	countConstructTests = []struct {
		target   string
		wordBank []string
		result   int
	}{
		{`abcdef`, []string{`ab`, `abc`, `cd`, `def`, `abcd`}, 1},
		{`purple`, []string{`purp`, `p`, `ur`, `le`, `purpl`}, 2},
		{`skateboard`, []string{`bo`, `rd`, `ate`, `t`, `ska`, `sk`, `boar`}, 0},
		{`enterapotentpot`, []string{`a`, `p`, `ent`, `enter`, `ot`, `o`, `t`}, 4},
		{`eeeeeeeeeeeeeeeeeeeeeeeeeeef`, []string{`e`, `ee`, `eee`, `eeee`, `eeeee`, `eeeeee`}, 0},
	}
)

func Test_countConstruct(t *testing.T) {
	for _, testCase := range countConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := countConstruct(testCase.target, testCase.wordBank)
			if res != testCase.result {
				t.Errorf("[%d] should be [%d]", res, testCase.result)
			}
		})
	}
}

func Test_countConstruct_memo(t *testing.T) {
	for _, testCase := range countConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := countConstruct_memo(testCase.target, testCase.wordBank)
			if res != testCase.result {
				t.Errorf("[%d] should be [%d]", res, testCase.result)
			}
		})
	}
}

func Test_countConstruct_tab(t *testing.T) {
	for _, testCase := range countConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := countConstruct_tab(testCase.target, testCase.wordBank)
			if res != testCase.result {
				t.Errorf("[%d] should be [%d]", res, testCase.result)
			}
		})
	}
}

func Benchmark_countConstruct(b *testing.B) {
	for _, testCase := range countConstructTests {
		b.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countConstruct(testCase.target, testCase.wordBank)
			}
		})
	}
}

func Benchmark_countConstruct_memo(b *testing.B) {
	for _, testCase := range countConstructTests {
		b.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countConstruct_memo(testCase.target, testCase.wordBank)
			}
		})
	}
}

func Benchmark_countConstruct_tab(b *testing.B) {
	for _, testCase := range countConstructTests {
		b.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countConstruct_tab(testCase.target, testCase.wordBank)
			}
		})
	}
}
