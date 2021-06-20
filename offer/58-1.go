package offer

import (
	"fmt"
	"strings"
)

/*输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/*
这里还是可以熟悉一些Go的strings标准库
strings.Join()
strings.TrimSpace()
strings.Split()
*/
func reverseWords(s string) string {
	i := len(s) - 1
	j := i
	var temp []string

	for i >= 0 && j >= 0 {
		if s[i] == ' ' && s[j] == ' ' { //去除尾部空格
			i--
			j--
		} else if s[i] != ' ' && s[j] != ' ' {

			for {
				if s[i] == ' ' || i == 0 {
					break
				}
				i--
			}
			if i == 0 && s[i] != ' ' {
				temp = append(temp, s[i:j+1])
				break
			} else {
				temp = append(temp, s[i+1:j+1])
			}
			j = i

		}
	}
	var result string

	for x, v := range temp {
		if x == len(temp)-1 {
			result = result + v
		} else {
			result = result + v + " "
		}
	}
	// 优雅输出 []string
	fmt.Printf("%#v", temp)
	return result

}

func reverseWords1(s string) string {
	strList := strings.Split(s, " ")
	var result []string
	for i := len(strList) - 1; i >= 0; i-- {
		//str:= strings.TrimSpace(v)
		if len(strList[i]) > 0 {
			result = append(result, strList[i])
		}
	}
	return strings.Join(result, " ")
}

func reverseWords2(s string) string {
	s = strings.TrimSpace(s) //去除首尾空格
	i := len(s) - 1
	j := i
	var temp []string
	for {
		i--
		if s[i] == ' ' {
			temp = append(temp, s[i+1:j+1])
			j=i
		}
	}
	return strings.Join(temp," ")

}
