package logic

import "fmt"

/*
#25 K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例：
给你这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5

说明：
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestReverseKGroup 测试 #25 K个一组翻转链表
func TestReverseKGroup() {
	lG := &ListNode{
		Val:  20,
		Next: nil,
	}
	lF := &ListNode{
		Val:  16,
		Next: lG,
	}
	lE := &ListNode{
		Val:  10,
		Next: lF,
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
	k := 4
	fmt.Println("#25 TestReverseKGroup Input:")
	PrintListNode(lA)
	res := reverseKGroup(lA, k)
	fmt.Println("#25 TestReverseKGroup Res:")
	PrintListNode(res)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	//特情
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil    //断后
	var start *ListNode = head //链表反转开始处
	var end *ListNode = head   //链表反转结束处
	var next *ListNode = nil   //前锋

	var newHead *ListNode = nil

	for i := 0; ; {
		if end.Next == nil {
			// subEnd, subHead := reverseOneGroup(start)
			// subEnd.Next = nil
			if pre == nil {
				return start
			}
			//pre.Next = subHead
			return newHead
		}
		end = end.Next
		next = end.Next
		i++
		if i == k-1 {
			end.Next = nil
			subEnd, subHead := reverseOneGroup(start)
			//处理断后
			if pre == nil {
				newHead = subHead
			} else {
				pre.Next = subHead
			}
			//处理前锋
			subEnd.Next = next
			//重新出发
			if next == nil {
				return newHead
			}
			pre = subEnd
			start = next
			end = next
			i = 0
		}
	}
}

//翻转一段链表
func reverseOneGroup(node *ListNode) (*ListNode, *ListNode) {
	if node == nil {
		return nil, nil
	}
	if node.Next == nil {
		return node, node
	}
	newPreNode, newHead := reverseOneGroup(node.Next)
	newPreNode.Next = node
	return node, newHead
}
