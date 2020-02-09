package logic

import "fmt"

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func TestReverse() {
	num := -1534236469
	fmt.Println("#7 TestReverse Input:")
	fmt.Println(num)
	res := reverse(num)
	fmt.Println("#7 TestReverse Res:")
	fmt.Println(res)
}

func reverse(x int) int {
	//处理特情
	res := 0
	if x == 0 || x < -2147483648 || x > 2147483648 {
		return res
	}
	isNeg := false
	if x < 0 {
		isNeg = true
		x = 0 - x
	}
	//直接解法
	//从第一位循环直到模为0则停止退出
	nRem := 0
	for {
		nRem = x % 10
		res = res*10 + nRem
		x = x / 10
		if x == 0 {
			break
		}
	}
	//出关也不能超出题干范围
	if res > 2147483648 {
		return 0
	}
	if isNeg {
		res = 0 - res
	}
	return res
}
