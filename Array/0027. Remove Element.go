package main

func removeElement(nums []int, val int) int {
	j := 0
	for i := range nums {
		if nums[i] != val {
			if i != j {
				nums[j], nums[i] = nums[i], nums[j]
			}
			j++
		}
	}
	return j
}

func main() {
	nums := []int{3, 2, 2, 3}
	aa := removeElement(nums, 3)
	println(aa)
}
