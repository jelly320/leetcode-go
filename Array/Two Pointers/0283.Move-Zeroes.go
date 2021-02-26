package main

// https://leetcode-cn.com/problems/move-zeroes/

//func moveZeroes(nums []int)  {
//	j:=0
//	for _,v := range nums {
//		if v==0 {
//			nums=append(nums,0)
//			fmt.Printf("tmp=o：%d\n",nums )
//			//nums[len(nums)-1]=0
//		} else {
//			nums[j]=v
//			j++
//			fmt.Printf("tmp!=0：%d\n",nums )
//		}
//	}
//}

func moveZeroes(nums []int) {
	//i指针1，j是指针2，
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for _, v := range nums {
		println(v)
	}
}
