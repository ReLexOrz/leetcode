package logic

import "fmt"

/*
#9 回文数
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

//TestIsPalindrome 测试 #9 回文数
func TestIsPalindrome() {
	num := 142241
	fmt.Println("#9 TestIsPalindrome Input:")
	fmt.Println(num)
	res := isPalindrome2(num)
	fmt.Println("#9 TestIsPalindrome Res:")
	fmt.Println(res)
}

//优化方法
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	orgx := x
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y == orgx {
		return true
	}
	return false
}

//自写方法
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	numList := make([]int, 0)
	for x != 0 {
		numList = append(numList, x%10)
		x = x / 10
	}
	for i, j := 0, len(numList)-1; j >= i; {
		if numList[i] != numList[j] {
			return false
		}
		i++
		j--
	}
	return true
}
