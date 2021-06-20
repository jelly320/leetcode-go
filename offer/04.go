package offer

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

func findNumberIn2DArray(matrix [][]int, target int) bool {

	i:= 0
	//j:= len(matrix[0]) - 1
	j:= len(matrix[i]) - 1
	println(j)

	for i<len(matrix) && j>=0 {
		flag:=matrix[i][j]

		if target < flag {
			j--
		} else if target > flag {
			i++
		} else if target == flag {
			return true
		}
	}
	return false

}