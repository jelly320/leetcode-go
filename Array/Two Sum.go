package main

func twoSum(nums []int, target int) []int {
	result := make([]int,2,2)
	for i := 0; i < len(nums); i++ {
		for y := i + 1; y < len(nums); y++ {
			if nums[i]+nums[y] == target {
				result[0] = i
				result[1] = y
			}
		}
	}
	return result
}

func main() {
	var nums = []int{2,7,11,15}
	aa := twoSum(nums,9)
	println(aa[0],aa[1])
}
