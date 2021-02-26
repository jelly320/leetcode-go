package main

// https://leetcode-cn.com/problems/first-bad-version/

/*v0*/
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
/*模拟isBadVersion,忽略*/
func isBadVersion(version int) bool {
	if version <= 3 {
		return false
	} else {
		return true
	}
}

/*v0
这道题虽然知道是类似找左边界的二分查找,不过题目有点绕
G-正确版本
B-错误版本
如果n=5,第4个版本为首个错误版本,则是:
G,G,G,B,B
这时候,isBadVersion的结果是:
false,false,false,true,true
所以,如果是false,就要一直右移,
如果是true,就要左变不动,不断收缩右边界
*/
func firstBadVersion(n int) int {

	//不要自己构造nums?
	//nums := []int{}
	//for i:=1;i<=n;i++{
	//	nums =append(nums,i)
	//}

	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid + 1) {
			//如果为true
			right = mid - 1
		} else {
			//如果为false
			left = mid + 1
		}
		if left >= n {
			return 0
		}
	}

	return left + 1
}

func main() {
	//假设n=5,则序列就是有序的[1,2,3,4,5]
	aa := firstBadVersion(5)
	println(aa)
}
