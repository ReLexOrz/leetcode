package logic

import "fmt"

/*
#20 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestIsValid 测试 #20 有效的括号
func TestIsValid() {
	s := "]"
	fmt.Println("#20 TestIsValid Input:")
	fmt.Println(s)
	res := isValid(s)
	fmt.Println("#20 TestIsValid Res:")
	fmt.Println(res)
}

//使用链表做一个堆栈
type CharNode struct {
	Ch   byte
	Next *CharNode
}

func isValid(s string) bool {
	var head *CharNode = nil
	for i := 0; i < len(s); i++ {
		ok, head_ := checkStack(s[i], head)
		head = head_
		if !ok {
			return false
		}
	}
	if head != nil {
		return false
	}
	return true
}

//内存待优化 57
func checkStack(ch byte, head *CharNode) (bool, *CharNode) {
	if head == nil {
		if ch == ')' || ch == ']' || ch == '}' {
			return false, nil
		}
		head = addStack(ch, head)
		return true, head
	}
	if ch == '(' || ch == '[' || ch == '{' {
		head = addStack(ch, head)
		return true, head
	}
	if (ch == ')' && head.Ch == '(') || (ch == ']' && head.Ch == '[') || (ch == '}' && head.Ch == '{') {
		head = popStack(head)
		return true, head
	}
	return false, head
}

//入栈
func addStack(ch byte, head *CharNode) *CharNode {
	if head == nil {
		head = &CharNode{
			Ch:   ch,
			Next: nil,
		}
		return head
	}
	newHead := &CharNode{
		Ch:   ch,
		Next: head,
	}
	return newHead
}

//弹栈
func popStack(head *CharNode) *CharNode {
	if head == nil {
		return head
	}
	head = head.Next
	return head
}
