package leetcode_go

import (
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_Problem(t *testing.T) {

	qs := []question{

		{
			para{321},
			ans{123},
		},

		{
			para{-123},
			ans{-321},
		},

		{
			para{120},
			ans{21},
		},

		{
			para{1534236469},
			ans{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem ------------------------\n")

	for _, q := range qs {
		_, p := q.ans, q.para
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, reverse(p.one))
	}
	fmt.Printf("\n\n\n")
}
