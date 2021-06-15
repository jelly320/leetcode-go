package offer

import "math"

/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
*/

func printNumbers(n int) []int {

	end := int(math.Pow(10, float64(n))) - 1
	var temp []int
	for i := 1; i <= end; i++ {
		temp = append(temp, i)
	}
	return temp
}
