package main

/*v0*/
//func twoSum(nums []int, target int) []int {
//	result := make([]int,2,2)
//	for i := 0; i < len(nums); i++ {
//		for y := i + 1; y < len(nums); y++ {
//			if nums[i]+nums[y] == target {
//				result[0] = i
//				result[1] = y
//			}
//		}
//	}
//	return result
//}

/*v1*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		for y := i + 1; y <= len(nums); y++ {
			if nums[y] == t {
				return []int{i, y}
			}
		}
	}
	return nil //没有这个return nil会报错是为啥？
}

func main() {
	var nums = []int{2, 7, 11, 15}
	aa := twoSum(nums, 9)
	println(aa[0], aa[1])
}
