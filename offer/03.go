package offer

/*
找出数组中任意一个重复的数字。
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
*/

func findRepeatNumber(nums []int) int {
	dir := map[int]int{}
	for _, v := range nums {
		_, ok := dir[v]
		if !ok {
			dir[v] = v
		} else {
			return v
		}
	}
	return -1
}

func findRepeatNumberv1(nums []int) int {
	//dir :=map[int]int{}
	for i, v := range nums {
		if i == v {
			continue
		} else if nums[v] == v {
			return nums[i]
		} else if nums[v] != v {
			nums[v] = v
		}
	}
	return -1
}
