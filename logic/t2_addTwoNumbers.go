package logic

import "fmt"

/*
#2 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestAddTwoNumbers 测试 #2两数相加
func TestAddTwoNumbers() {

	l1C := &ListNode{
		Val:  3,
		Next: nil,
	}
	l1B := &ListNode{
		Val:  4,
		Next: l1C,
	}
	l1A := &ListNode{
		Val:  2,
		Next: l1B,
	}

	l2C := &ListNode{
		Val:  6,
		Next: nil,
	}
	l2B := &ListNode{
		Val:  4,
		Next: l2C,
	}
	l2A := &ListNode{
		Val:  5,
		Next: l2B,
	}
	fmt.Println("#2 TestAddTwoNumbers Input:")
	PrintListNode(l1A)
	PrintListNode(l2A)
	res := addTwoNumbers(l1A, l2A)
	fmt.Println("#2 TestAddTwoNumbers Res:")
	PrintListNode(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode 用于int单向链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carNum := 0
	po1 := l1
	po2 := l2
	for true {
		add1 := 0
		add2 := 0
		if po1 != nil {
			add1 = po1.Val
		}
		if po2 != nil {
			add2 = po2.Val
		}
		addRes := add1 + add2 + carNum
		if addRes >= 10 {
			carNum = 1
			addRes = addRes - 10
		} else {
			carNum = 0
		}
		//赋值
		if po1 != nil {
			po1.Val = addRes
		}
		if po2 != nil {
			po2.Val = addRes
		}

		//两个表一样长 追加一个然后退出
		if po1 != nil && po2 != nil && po1.Next == nil && po2.Next == nil {
			if carNum > 0 {
				po1.Next = &ListNode{
					Val: carNum,
				}
			}
			return l1
		}

		//两个表不一样长分为两种情况 l1长
		if po2 == nil && po1 != nil && po1.Next == nil {
			if carNum > 0 {
				po1.Next = &ListNode{
					Val: carNum,
				}
			}
			return l1
		}

		//两个表不一样长分为两种情况 l2长
		if po1 == nil && po2 != nil && po2.Next == nil {
			if carNum > 0 {
				po2.Next = &ListNode{
					Val: carNum,
				}
			}
			return l2
		}

		//指针前进
		if po1 != nil {
			po1 = po1.Next
		}
		if po2 != nil {
			po2 = po2.Next
		}
	}
	return nil
}
