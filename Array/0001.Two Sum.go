package main

/*v0*/
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		for y := i + 1; y < len(nums); y++ {
//			if nums[i]+nums[y] == target {
//				return []int{i,y}
//			}
//		}
//	}
//	return nil
//}

/*v1*/
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		t := target - nums[i]
//		for y := i + 1; y <= len(nums); y++ {
//			if nums[y] == t {
//				return []int{i, y}
//			}
//		}
//	}
//	return nil //没有这个return nil会报错是为啥？
//}

/*v2 hash*/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if idex, ok := m[target-v]; ok {
			return []int{idex, i}
		}
		m[v] = i //如果有相同在array出现，只能返回最后一个
	}
	return nil
}

func main() {
	var nums = []int{3, 3, 7, 15}
	aa := twoSum(nums, 10)
	println(aa[0], aa[1])
}
