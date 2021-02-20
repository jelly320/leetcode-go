package main

// https://leetcode-cn.com/problems/search-insert-position/

/*v0*/
//func searchInsert(nums []int, target int) int {
//	j := 0
//	for i, v := range nums {
//		if v == target {
//			return i
//		}
//		if v < target {
//			j++
//		}
//	}
//	return j
//}

/*v1 二分查找?*/
//题目给定的条件是一个"排好序的数组"
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		mid := left + (right-left)/2

		if left > right {
			break
		}
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	//aa := searchInsert(nums,5)
	//aa := searchInsert(nums,0)
	//aa := searchInsert(nums,2)
	aa := searchInsert(nums, 7)
	println(aa)
}
