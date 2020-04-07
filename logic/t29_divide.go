package logic

import (
	"fmt"
	"math"
)

/*
#29 两数相除

给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:
输入: dividend = 10, divisor = 3
输出: 3

示例 2:
输入: dividend = 7, divisor = -3
输出: -2
说明:

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divide-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestDivide 测试 #29 两数相除
func TestDivide() {
	dividend := -1
	divisor := 1
	fmt.Println("#29 TestDivide Input:")
	fmt.Println(dividend)
	fmt.Println(divisor)
	res := divide(dividend, divisor)
	fmt.Println("#29 TestDivide Res:")
	fmt.Println(res)
}

//高效逼近
func divide(dividend int, divisor int) int {
	isPos := true
	if dividend < 0 && divisor > 0 {
		isPos = false
		dividend = -dividend
	}
	if dividend > 0 && divisor < 0 {
		isPos = false
		divisor = -divisor
	}
	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}

	if dividend >= math.MaxInt32 && divisor == 1 {
		if !isPos {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	if divisor == 1 {
		if !isPos {
			return -dividend
		}
		return dividend
	}
	if divisor == 0 {
		return math.MaxInt32
	}
	if dividend == 0 {
		return 0
	}

	res := CalDivideCount(dividend, divisor)

	if res > math.MaxInt32 {
		if !isPos {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	if !isPos {
		return -res
	}
	return res
}

//CalDivideCount 递归调用
func CalDivideCount(dividend int, divisor int) int {
	//递归退出条件
	if dividend < divisor {
		return 0
	}
	mulDivisor := divisor
	count := 1
	for mulDivisor+mulDivisor < dividend {
		mulDivisor = mulDivisor + mulDivisor
		count = count + count
	}

	return count + CalDivideCount(dividend-mulDivisor, divisor)
}
