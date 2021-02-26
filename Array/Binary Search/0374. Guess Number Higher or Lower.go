package main

//  https://leetcode-cn.com/problems/guess-number-higher-or-lower/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
/*guess的返回情况:
-1：我选出的数字比你猜的数字小 pick < num
1：我选出的数字比你猜的数字大 pick > num
0：我选出的数字和你猜的数字一样。恭喜！你猜对了！pick == num

前提条件:
    1 <= n <= 231 - 1
    1 <= pick <= n
	还有一个重要前提：每轮游戏，我都会从 1 到 n 随机选择一个数字。　即意味着pick这个数字的范围是在１－n之间的ｓ. 所以依旧还是个范围确定的二分查找...
*/

/*模拟guess*/
func guess(num int) int {
	pick := 6
	if pick < num {
		return -1
	} else if pick > num {
		return 1
	} else if pick == num {
		return 0
	}
	return 0
}

/*v0*/
func guessNumber(n int) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if guess(mid+1) == 0 {
			return mid + 1
		} else if guess(mid+1) == 1 {
			//nums[mid]<pick,右移
			left = mid + 1
		} else if guess(mid+1) == -1 {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	//aa := guessNumber(6) //pick:=6
	aa := guessNumber(60) //pick:=6
	println(aa)
}
