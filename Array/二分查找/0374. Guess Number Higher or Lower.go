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
*/

/*模拟guess*/
func guess(num int) int {
	pick := 10
	if pick < num {
		return -1
	} else if pick > num {
		return 1
	} else if pick == num {
		return 0
	}
	return 0
}

/*v0
这个题的难点在于,怎么划定边界呢...
*/

func guessNumber(n int) int {
	if guess(n) == 0 {
		return n
	}
	return -1
}

func main() {
	//aa := guessNumber(10) //pick:=6
	aa := guessNumber(10) //pick:=10
	println(aa)
}
