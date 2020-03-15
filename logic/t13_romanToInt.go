package logic

import "fmt"

/*
#13 罗马数字转整数

罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/roman-to-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

//TestRomanToInt 测试 #13罗马数字转整数
func TestRomanToInt() {
	s := "MMMXLV"
	fmt.Println("#13 TestRomanToInt Input:")
	fmt.Println(s)
	res := romanToInt(s)
	fmt.Println("#13 TestRomanToInt Res:")
	fmt.Println(res)
}

func romanToInt2(s string) int {
	res := 0
	sLen := len(s)

	for i := sLen - 1; i >= 0; {
		if s[i] == 'I' {
			res = res + 1
			i--
			continue
		}
		if s[i] == 'V' {
			if (i-1) >= 0 && s[i-1] == 'I' {
				res = res + 4
				i = i - 2
				continue
			} else {
				res = res + 5
				i--
				continue
			}
		}
		if s[i] == 'X' {
			if (i-1) >= 0 && s[i-1] == 'I' {
				res = res + 9
				i = i - 2
				continue
			} else {
				res = res + 10
				i--
				continue
			}
		}
		if s[i] == 'L' {
			if (i-1) >= 0 && s[i-1] == 'X' {
				res = res + 40
				i = i - 2
				continue
			} else {
				res = res + 50
				i--
				continue
			}
		}

		if s[i] == 'C' {
			if (i-1) >= 0 && s[i-1] == 'X' {
				res = res + 90
				i = i - 2
				continue
			} else {
				res = res + 100
				i--
				continue
			}
		}

		if s[i] == 'D' {
			if (i-1) >= 0 && s[i-1] == 'C' {
				res = res + 400
				i = i - 2
				continue
			} else {
				res = res + 500
				i--
				continue
			}
		}

		if s[i] == 'M' {
			if (i-1) >= 0 && s[i-1] == 'C' {
				res = res + 900
				i = i - 2
				continue
			} else {
				res = res + 1000
				i--
				continue
			}
		}
	}

	return res
}

//

//解法2 可以看下一位是否小于上一位 如果小于则做减法 速度81 内存99
func romanToInt(s string) int {
	res := 0
	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		num := getValueByRoman(s[i])
		if pre == 0 {
			pre = num
		} else {
			pre = getValueByRoman(s[i+1])
		}
		if pre > num {
			res -= num
		} else {
			res += num
		}
	}
	return res
}

func getValueByRoman(romanChar byte) int {
	switch romanChar {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}
