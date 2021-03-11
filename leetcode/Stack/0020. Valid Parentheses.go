package Stack

// https://leetcode-cn.com/problems/valid-parentheses/

/*v0
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了17.03%的用户
*/
type ValidStack struct {
	stack []string
	top   int
}

func Constructor() ValidStack {
	return ValidStack{
		[]string{},
		0,
	}
}

func (this *ValidStack) Push(s string) {
	this.stack = append(this.stack, s)
	this.top++
}

func (this *ValidStack) Pop() {
	if this.top != 0 {
		this.top--
		this.stack = this.stack[:this.top]
	}
}

func (this *ValidStack) Top() string {
	if this.top > 0 {
		return this.stack[this.top-1]
	}
	return ""
}

func JudgeIn(s string, lst []string) bool {
	tmp := make(map[string]int)
	for i, v := range lst {
		tmp[v] = i
	}
	if _, ok := tmp[s]; ok {
		return true
	} else {
		return false
	}
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	obj := Constructor()

	left := []string{"(", "{", "["}
	corr := map[string]string{")": "(", "}": "{", "]": "["}
	//right := []string{")","}","]"}

	for _, v := range s {
		if JudgeIn(string(v), left) {
			obj.Push(string(v))
		} else { //因为题目说了只会有6种字符，我们就默认不是左就是右了
			if obj.top == 0 {
				return false
			} else {
				if obj.Top() == corr[string(v)] {
					obj.Pop()
				} else {
					obj.Push(string(v))
				}
			}
		}
	}

	if obj.top != 0 {
		return false
	}
	return true
}

/* v1 参考了一些解法
我的v0存在哪些问题：1、没有转变思维，不是转成string()才能操作的；2、判断字符是不是左，思维也是比较py，go其实判断挺麻烦，我自己也是想得复杂了;
3、使用栈的思维≠实现栈，使用slice活用slice的操作也是一个道理

执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2 MB, 在所有 Go 提交中击败了41.72%的用户
*/

func isValid_v1(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]rune, 0)
	corr := map[string]string{")": "(", "}": "{", "]": "["}
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				//fmt.Printf("type: %v\n",stack[len(stack)-1])
				if string(stack[len(stack)-1]) == corr[string(v)] {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, v)
				}
			}

		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
