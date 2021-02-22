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
//func searchInsert(nums []int, target int) int {
//	left := 0
//	right := len(nums) - 1
//
//	for {
//		mid := left + (right-left)/2
//
//		if left > right {
//			break
//		}
//		if nums[mid] == target {
//			return mid
//		} else if nums[mid] > target {
//			right = mid - 1
//		} else if nums[mid] < target {
//			left = mid + 1
//		}
//	}
//	return left
//}

/*参考 halfrost LeetCode-Go*/
//这个是beat 100%的
//func searchInsert(nums []int, target int) int {
//	low, high := 0, len(nums)-1
//	for low <= high {
//		mid := low + (high-low)>>1
//		if nums[mid] >= target {
//			high = mid - 1
//		} else {
//			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
//				return mid + 1
//			}
//			low = mid + 1
//		}
//	}
//	return 0
//}
/*参考后写的*/
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}

func main() {
	//nums := []int{1, 3, 5, 6}
	//aa := searchInsert(nums,5)
	//aa := searchInsert(nums,0)
	//aa := searchInsert(nums,2)

	nums := []int{1, 3, 3, 3, 5, 6}
	//aa := searchInsert(nums, 3)
	//aa := searchInsert(nums, 7)
	//aa := searchInsert(nums, 0)
	aa := searchInsert(nums, 4)
	println(aa)
}
