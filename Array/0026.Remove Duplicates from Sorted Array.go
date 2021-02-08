//　https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
package main

/*v0*/
//func removeDuplicates(nums []int) int {
//	m := map[int]int{}
//	j := 0
//	for i, v := range nums {
//		if _, ok := m[v]; !ok {
//			if len(m) != 0 {
//				j++
//				nums[i], nums[j] = nums[j], nums[i]
//			}
//			m[v] = i
//		}
//	}
//	return len(m)
//}

/*v1*/
func removeDuplicates(nums []int) int {
	j := 0
	for i := range nums {
		if nums[i] != nums[j] {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return j + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{1, 1, 2, 2}
	aa := removeDuplicates(nums)
	println(aa)

	for _, v := range nums {
		println(v)
	}
}
