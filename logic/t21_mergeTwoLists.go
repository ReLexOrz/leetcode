package logic

import "fmt"

/*
#21 合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestMergeTwoLists 测试 #21合并两个有序链表
func TestMergeTwoLists() {

	l1C := &ListNode{
		Val:  4,
		Next: nil,
	}
	l1B := &ListNode{
		Val:  3,
		Next: l1C,
	}
	l1A := &ListNode{
		Val:  1,
		Next: l1B,
	}

	l2D := &ListNode{
		Val:  8,
		Next: nil,
	}
	l2C := &ListNode{
		Val:  4,
		Next: l2D,
	}
	l2B := &ListNode{
		Val:  2,
		Next: l2C,
	}
	l2A := &ListNode{
		Val:  1,
		Next: l2B,
	}
	fmt.Println("#2 TestAddTwoNumbers Input:")
	PrintListNode(l1A)
	PrintListNode(l2A)
	res := mergeTwoLists(l1A, l2A)
	fmt.Println("#2 TestAddTwoNumbers Res:")
	PrintListNode(res)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//特情
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	//以l1做主链
	var preNode *ListNode = nil
	curNode := l1
	for curNode != nil && l2 != nil {
		if curNode.Val >= l2.Val {
			tpl2 := l2
			l2 = l2.Next
			tpl2.Next = curNode
			if preNode == nil {
				l1 = tpl2
				preNode = l1
			} else {
				preNode.Next = tpl2
				preNode = tpl2
			}
		} else {
			preNode = curNode
			curNode = curNode.Next
		}
	}
	if curNode == nil && l2 != nil {
		if preNode != nil {
			preNode.Next = l2
		}
	}
	return l1
}
