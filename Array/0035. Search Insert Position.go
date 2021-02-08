package main

// https://leetcode-cn.com/problems/search-insert-position/

/*v0*/
func searchInsert(nums []int, target int) int {
	j := 0
	for i, v := range nums {
		if v == target {
			return i
		}
		if v < target {
			j++
		}
	}
	return j
}

func main() {
	nums := []int{1, 3, 5, 6}
	//aa := searchInsert(nums,5)
	aa := searchInsert(nums, 7)
	//aa := searchInsert(nums,0)
	//aa := searchInsert(nums,2)
	println(aa)
}
