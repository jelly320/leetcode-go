package main

import "fmt"

/*寻找一个数（基本的二分搜索）*/
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for {
		mid := left + (right-left)/2
		fmt.Printf("起始区间:[left,right]=[%d,%d],mid: %d\n", left, right, mid)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			//如果小于target,则区间往右移
			left = mid + 1
			fmt.Printf("xiaoyu;区间:[left,right]=[%d,%d]\n", left, right)
		} else if nums[mid] > target {
			//如果大于target,则区间往左移
			right = mid - 1
			fmt.Printf("dayu;区间:[left,right]=[%d,%d]\n", left, right)
		}
		//} else if left >= right {
		//	//循环退出条件
		//	fmt.Printf("left: %d,right: %d",left,right)
		//	break
		//}
		//设置循环的退出条件
		if left > right {
			break
		}
	}
	return -1
}

/*寻找左侧边界的二分搜索*/
func left_bound(nums []int, target int) int {

}

/*寻找右侧边界的二分查找*/
func right_bound(nums []int, target int) int {

}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	/*测试以下内容:*/
	//index := binarySearch(nums,6)
	index := binarySearch(nums, 9)
	fmt.Printf("result: %d", index)
}
