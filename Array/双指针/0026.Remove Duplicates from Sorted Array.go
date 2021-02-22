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

/*TODO v3 还有优化的话*/
// 如果数组都是不重复的 [0,1,2,3] ,判不等时交换其实有点浪费,这是1点. 判断成本vs复制成本?
func removeDuplicates(nums []int) int {
	j := 0
	for i := range nums {
		if nums[i] != nums[j] {
			j++
			if i-j > 1 {
				nums[i], nums[j] = nums[j], nums[i]
			}

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
