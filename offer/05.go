package offer

import "strings"

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/

func replaceSpace(s string) string {
	tmp := make([]string,len(s)*3)
	for i, v := range s {
		if string(v) == " " {
			tmp[i] = "%20"
		} else {
			tmp[i] = string(v)
		}
	}
	return strings.Join(tmp, "")
}


//	bytes:=[]byte(s)
//	for i,v:=range bytes{
//		//println(i,v)
//		if string(v) == " " {
//			append(bytes, "%20")
//		} else {
//			bytes[i]=v
//		}
//	}
//	return "aa"
