package leetcode

import (
	"fmt"
	"go-leetcode/structures"
	"testing"
)

type question144 struct {
	para144
	ans144
}

// para 是参数
// one 代表第一个参数
type para144 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans144 struct {
	one []int
}

func Test_Problem144(t *testing.T) {

	qs := []question144{

		{
			para144{[]int{}},
			ans144{[]int{}},
		},

		{
			para144{[]int{1}},
			ans144{[]int{1}},
		},

		{
			para144{[]int{1, structures.NULL, 2, 3}},
			ans144{[]int{1, 2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 144------------------------\n")
	tagnum := 0
	for _, q := range qs {
		_, p := q.ans144, q.para144
		fmt.Printf("【input】:%v      \n", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", preorderTraversal(root, tagnum))
	}
	fmt.Printf("\n\n\n")
}
