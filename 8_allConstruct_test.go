package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	allConstructTests = []struct {
		target   string
		wordBank []string
		result   [][]string
	}{
		{`abcdef`, []string{`ab`, `abc`, `cd`, `def`, `abcd`}, [][]string{
			{`abc`, `def`},
		}},
		{`purple`, []string{`purp`, `p`, `ur`, `le`, `purpl`}, [][]string{
			{`purp`, `le`},
			{`p`, `ur`, `p`, `le`},
		}},
		{`abcdef`, []string{`ab`, `abc`, `cd`, `def`, `abcd`, `ef`, `c`}, [][]string{
			{`ab`, `cd`, `ef`},
			{`ab`, `c`, `def`},
			{`abc`, `def`},
			{`abcd`, `ef`},
		}},
		// {`skateboard`, []string{`bo`, `rd`, `ate`, `t`, `ska`, `sk`, `boar`}, 0},
		// {`enterapotentpot`, []string{`a`, `p`, `ent`, `enter`, `ot`, `o`, `t`}, 4},
		// {`eeeeeeeeeeeeeeeeeeeeeeeeeeef`, []string{`e`, `ee`, `eee`, `eeee`, `eeeee`, `eeeeee`}, 0},
	}
)

func revertSlSl(sl [][]string) {
	for _, s := range sl {
		revertSlice(s)
	}
}

func revertSlice(sl []string) {
	for i := 0; i < len(sl)/2; i++ {
		buf := sl[i]
		sl[i] = sl[len(sl)-1-i]
		sl[len(sl)-1-i] = buf
	}
}

func Test_revertSlice(t *testing.T) {
	sl := []string{`purp`, `le`, `100`}
	revertSlice(sl)
	fmt.Println(sl)
}

func Test_allConstruct(t *testing.T) {
	for _, testCase := range allConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := allConstruct(testCase.target, testCase.wordBank)
			revertSlSl(res)
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("%v should be %v", res, testCase.result)
			}
		})
	}
}

func Test_allConstruct_memo(t *testing.T) {
	for _, testCase := range allConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := allConstruct_memo(testCase.target, testCase.wordBank)
			revertSlSl(res)
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("%v should be %v", res, testCase.result)
			}
		})
	}
}

func Test_allConstruct_tab(t *testing.T) {
	for _, testCase := range allConstructTests {
		t.Run(fmt.Sprintf("%s %v", testCase.target, testCase.wordBank), func(t *testing.T) {
			res := allConstruct_tab(testCase.target, testCase.wordBank)
			if !reflect.DeepEqual(res, testCase.result) {
				t.Errorf("%v should be %v", res, testCase.result)
			}
		})
	}
}
