package logic

import "fmt"

/*
#10 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TeatIsMatch #10 正则表达式匹配
func TeatIsMatch() {
	s := "bbbba"
	p := ".*a*a"
	fmt.Println("#10 TeatIsMatch Input:")
	fmt.Println(s)
	fmt.Println(p)
	res := isMatch2(s, p)
	fmt.Println("#10 TeatIsMatch Res:")
	fmt.Println(res)
}

func isMatch(s string, p string) bool {
	star := byte('*')
	dot := byte('.')
	pLen := len(p)
	sLen := len(s)
	if pLen == 0 && sLen == 0 {
		return true
	}
	if pLen == 0 {
		return false
	}
	if p == ".*" {
		return true
	}
	for si, pi := sLen-1, pLen-1; ; {
		//缺乏匹配器 但是字符串仍有多余
		if pi < 0 && si >= 0 {
			return false
		}
		//都没有了 返回匹配成功
		if pi < 0 && si < 0 {
			break
		}
		//此种情况匹配器为重复字串匹配
		if pi-1 >= 0 && p[pi] == star {
			mateByte := p[pi-1]
			//直到所有相同字符匹配完毕
			for ; si >= 0; si-- {
				if s[si] != mateByte {
					break
				}
			}
			//pi要跳两个字
			pi = pi - 2
		} else if p[pi] == dot { //此种情况匹配器为单一任意字符匹配
			if si < 0 {
				return false
			}
			si--
			pi--
		} else { //此种情况为匹配器为特殊字符 必须匹配
			if si < 0 {
				return false
			}
			if p[pi] != s[si] {
				return false
			}
			si--
			pi--
		}
	}
	return true
}

//递归解法
func isMatch2(s string, p string) bool {
	pLen := len(p)
	sLen := len(s)
	if pLen == 0 && sLen == 0 {
		return true
	}
	if pLen == 0 {
		return false
	}
	if p == ".*" {
		return true
	}
	return matchSwitchState(0, 0, s, p)
}

func matchSwitchState(si int, pi int, s string, p string) bool {
	pLen := len(p)
	sLen := len(s)
	//结果集
	//到达尾端
	if pi >= pLen && si >= sLen {
		return true
	}
	//匹配器没有了 但是字符还有
	if pi >= pLen && si < sLen {
		return false
	}
	//分情况向下递归
	if pi+1 < len(p) && p[pi+1] == byte('*') {
		mateByte := p[pi]
		//匹配超出 但是可以继续
		if si >= sLen {
			return matchSwitchState(si, pi+2, s, p)
		}
		//匹配到了有三种种选择 注意，匹配有两种情况 比较特殊的一种是'.*'可以和任意字符匹配
		//三种步进方式都有可能
		if s[si] == mateByte || mateByte == byte('.') {
			return matchSwitchState(si, pi+2, s, p) || matchSwitchState(si+1, pi, s, p) || matchSwitchState(si+1, pi+2, s, p)
		}
		//没有匹配到则只能p继续下去
		return matchSwitchState(si, pi+2, s, p)
	} else if p[pi] == byte('.') {
		//点号 可匹配任何字符
		if si >= sLen {
			return false
		}
		return matchSwitchState(si+1, pi+1, s, p)
	} else {
		//具体字符 只能严格匹配
		if si >= sLen {
			return false
		}
		if p[pi] != s[si] {
			return false
		}
		return matchSwitchState(si+1, pi+1, s, p)
	}
}
