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
	lE := &ListNode{
		Val:  10,
		Next: nil,
	}
	lD := &ListNode{
		Val:  8,
		Next: lE,
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
	fmt.Println("#24 TestSwapPairs Input:")
	PrintListNode(lA)
	res := swapPairs(lA)
	fmt.Println("#24 TestSwapPairs Res:")
	PrintListNode(res)
}

func swapPairs2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var preNode *ListNode
	nodeA := head
	nodeB := head.Next
	for {
		//交换preNode和curNode
		nodeA.Next = nodeB.Next
		nodeB.Next = nodeA
		if preNode != nil {
			preNode.Next = nodeB
		} else {
			head = nodeB
		}
		//跳到下一组
		if nodeA.Next == nil {
			return head
		}
		if nodeA.Next.Next == nil {
			return head
		}
		preNode = nodeA
		nodeB = nodeA.Next.Next
		nodeA = nodeA.Next
	}
}

//递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	nodeA := head
	nodeB := head.Next

	nodeA.Next = swapPairs(nodeB.Next)
	nodeB.Next = nodeA

	return nodeB
}
