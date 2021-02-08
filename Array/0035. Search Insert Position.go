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

/*v1 二分查找?*/
//题目给定的条件是一个"排好序的数组"
func searchInsert(nums []int, target int) int {
	left := 0
	right := (len(nums) - 1) / 2
	if nums[right] == target {
		return right
	} else if nums[right] < target {

	} else if nums[right] > target {
		right := right / 2

	}

}

func main() {
	nums := []int{1, 3, 5, 6}
	//aa := searchInsert(nums,5)
	aa := searchInsert(nums, 7)
	//aa := searchInsert(nums,0)
	//aa := searchInsert(nums,2)
	println(aa)
}
