package Stack

import "strconv"

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/

/* v0
*/
type PolishStack struct {
	stack []string
	top int
	oper int
}

func Constructor() PolishStack {
	return PolishStack{
		[]string,
		0,
		0,
	}
}

func evalRPN(tokens []string) int {
	this := Constructor()
	for _,v:=range tokens {
		if v=="+" || v=="-" || v=="*" ||v=="/" {
			this.oper=this.top
			oper_num1:=this.stack[this.top-2]
			oper_num2 := this.stack[this.top-1]
			res := CalNum(oper_num1,oper_num2,v)
			this.stack[this.top-1]=res
		}
		this.stack=append(this.stack,v)
		this.top++
	}

}

func CalNum(opernum1 string,opernum2 string, oper string) string {
	var res int
	num1,_ := strconv.Atoi(opernum1)
	num2,_ := strconv.Atoi(opernum2)
	if oper == "+" {
		res=num1+num2
	} else if oper =="-" {
		res=num1 - num2
	} else if oper == "*" {
		res=num1*num2
	} else if oper == "/" {
		res=num1/num2
	}
	return string(res)
}
