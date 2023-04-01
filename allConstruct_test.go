package dynamic

import (
	"fmt"
	"testing"
	"time"
)

func Test_AllConstruct(t *testing.T) {
	runTest(allConstruct, `purple`, []string{`purp`, `p`, `ur`, `le`, `purpl`})
	runTest(allConstruct, `abcdef`, []string{`ab`, `abc`, `cd`, `def`, `abcd`, `ef`, `c`})
	runTest(allConstruct, `skateboard`, []string{`bo`, `rd`, `ate`, `t`, `ska`, `sk`, `boar`})
	runTest(allConstruct, `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz`, []string{`a`, `aa`, `aaa`, `aaaa`, `aaaaa`})

	runTest(allConstructMemo, `purple`, []string{`purp`, `p`, `ur`, `le`, `purpl`})
	runTest(allConstructMemo, `abcdef`, []string{`ab`, `abc`, `cd`, `def`, `abcd`, `ef`, `c`})
	runTest(allConstructMemo, `skateboard`, []string{`bo`, `rd`, `ate`, `t`, `ska`, `sk`, `boar`})
	runTest(allConstructMemo, `aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz`, []string{`a`, `aa`, `aaa`, `aaaa`, `aaaaa`})
}

func runTest(f func(string, []string) [][]string, w string, wb []string) {
	globalCount = 0
	t1 := time.Now()
	res := f(w, wb)
	t2 := time.Now()
	fmt.Println("iters [", globalCount, "] time [", t2.Sub(t1), "]", res)
}
