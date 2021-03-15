package Stack

import (
	"strconv"
)

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/

/* v0
执行用时：4 ms, 在所有 Go 提交中击败了89.75%的用户
内存消耗：4.3 MB, 在所有 Go 提交中击败了21.12%的用户
*/
type PolishStack struct {
	stack []string
	top   int
}

func Constructor_0155() PolishStack {
	return PolishStack{
		[]string{},
		0,
	}
}

func evalRPN(tokens []string) int {
	this := Constructor_0155()
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			oper_num1 := this.stack[this.top-2]
			oper_num2 := this.stack[this.top-1]
			res := CalNum(oper_num1, oper_num2, v)
			this.top = this.top - 2
			this.stack = this.stack[:this.top]
			this.stack = append(this.stack, res)
		} else {
			this.stack = append(this.stack, v)
		}
		this.top++
	}
	result, _ := strconv.Atoi(this.stack[this.top-1])
	return result
}

func CalNum(opernum1 string, opernum2 string, oper string) string {
	var res int
	num1, _ := strconv.Atoi(opernum1)
	num2, _ := strconv.Atoi(opernum2)
	if oper == "+" {
		res = num1 + num2
	} else if oper == "-" {
		res = num1 - num2
	} else if oper == "*" {
		res = num1 * num2
	} else if oper == "/" {
		res = num1 / num2
	}
	return strconv.Itoa(res)
}
