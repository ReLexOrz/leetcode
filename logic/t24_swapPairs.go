package logic

import "fmt"

/*
#24 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestSwapPairs 测试 #24两两交换链表中的节点
func TestSwapPairs() {

	lD := &ListNode{
		Val:  8,
		Next: nil,
	}
	lC := &ListNode{
		Val:  4,
		Next: lD,
	}
	lB := &ListNode{
		Val:  2,
		Next: lC,
	}
	lA := &ListNode{
		Val:  1,
		Next: lB,
	}
	fmt.Println("#24 TestAddTwoNumbers Input:")
	PrintListNode(lA)
	res := swapPairs(lA)
	fmt.Println("#24 TestAddTwoNumbers Res:")
	PrintListNode(res)
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var startNode *ListNode = head.Next
	var ppreNode *ListNode = nil
	var preNode *ListNode = head
	var curNode *ListNode = head.Next

	for {
		preNode.Next = curNode.Next
		curNode.Next = preNode
		if ppreNode == nil {
			ppreNode = curNode
		} else {
			ppreNode.Next = curNode
		}
		if preNode.Next == nil {
			return startNode
		}
		if preNode.Next.Next == nil {
			return startNode
		}
		curNode = preNode.Next.Next
		preNode = preNode.Next
	}
}
