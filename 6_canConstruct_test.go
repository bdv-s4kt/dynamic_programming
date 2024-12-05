package dynamic

import (
	"fmt"
	"testing"
)

var (
	canConstructTests = []struct {
		target   string
		wordBank []string
		result   bool
	}{
		{`abcdef`, []string{`ab`, `abc`, `cd`, `def`, `abcd`}, true},
		{`skateboard`, []string{`bo`, `rd`, `ate`, `t`, `ska`, `sk`, `boar`}, false},
		{``, []string{`cat`, `dog`, `mouse`}, true},
		{`enterapotentpot`, []string{`a`, `p`, `ent`, `enter`, `ot`, `o`, `t`}, true},
		{`eeeeeeeeeeeeeeeeeeeeeeeeeeef`, []string{`e`, `ee`, `eee`, `eeee`, `eeeee`, `eeeeee`}, false},
	}
)

func Test_canConstruct(t *testing.T) {
	for _, testCase := range canConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := canConstruct(testCase.target, testCase.wordBank)
			if res != testCase.result {
				t.Errorf("[%t] should be [%t]", res, testCase.result)
			}
		})
	}
}

func Test_canConstruct_memo(t *testing.T) {
	for _, testCase := range canConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := canConstruct_memo(testCase.target, testCase.wordBank)
			if res != testCase.result {
				t.Errorf("[%t] should be [%t]", res, testCase.result)
			}
		})
	}
}

func Test_canConstruct_tab(t *testing.T) {
	for _, testCase := range canConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := canConstruct_tab(testCase.target, testCase.wordBank)
			if res != testCase.result {
				t.Errorf("[%t] should be [%t]", res, testCase.result)
			}
		})
	}
}

func Benchmark_canConstruct(b *testing.B) {
	for _, testCase := range canConstructTests {
		b.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canConstruct(testCase.target, testCase.wordBank)
			}
		})
	}
}

func Benchmark_canConstruct_memo(b *testing.B) {
	for _, testCase := range canConstructTests {
		b.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canConstruct_memo(testCase.target, testCase.wordBank)
			}
		})
	}
}

func Benchmark_canConstruct_tab(b *testing.B) {
	for _, testCase := range canConstructTests {
		b.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				canConstruct_tab(testCase.target, testCase.wordBank)
			}
		})
	}
}
