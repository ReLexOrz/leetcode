package logic

import "fmt"

/*
#22 删除链表的倒数第N个节点
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestGenerateParenthesis 测试 #22 括号生成
func TestGenerateParenthesis() {
	n := 3
	fmt.Println("#22 TestGenerateParenthesis Input:")
	fmt.Println(n)
	res := generateParenthesis(n)
	fmt.Println("#22 TestGenerateParenthesis Res:")
	fmt.Println(res)
}

//递归生成方式
func generateParenthesis(n int) []string {
	resList := make([]string, 0, 0)
	AddChar("(", "", 0, 0, &resList, n)
	return resList
}

func AddChar(addChar string, curString string, posCount int, naviCount int, resList *[]string, n int) {
	//统计之前的字符串有多少
	curString = curString + addChar
	if addChar == "(" {
		posCount += 1
	}
	if addChar == ")" {
		naviCount += 1
	}
	if posCount > n || naviCount > n {
		return
	}
	if naviCount > posCount {
		return
	}
	if len(curString) == 2*n {
		if posCount == naviCount {
			*resList = append(*resList, curString)
		}
		return
	}
	//两种追加方式继续追加
	AddChar("(", curString, posCount, naviCount, resList, n)
	AddChar(")", curString, posCount, naviCount, resList, n)
}
