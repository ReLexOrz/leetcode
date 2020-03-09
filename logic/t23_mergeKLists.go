package logic

import "fmt"

/*
#23 合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestIsValid 测试 #23 合并K个排序链表
func TestMergeKLists() {
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

	lists := []*ListNode{l1A, l2A}
	fmt.Println("#23 TestMergeKLists Input:")
	for _, pNode := range lists {
		PrintListNode(pNode)
	}
	res := mergeKLists(lists)
	fmt.Println("#23 TestMergeKLists Res:")
	PrintListNode(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//TODO 普通方法 遍历所有List插入排序
func mergeKLists(lists []*ListNode) *ListNode {

	//特情
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	//var preNode *ListNode = nil
	var curNode *ListNode = nil
	//var startNode *ListNode = nil

	for _, nodeN := range lists {
		curNode = mergeTwoNodeList(curNode, nodeN)
	}

	return curNode
}

func mergeTwoNodeList(l1 *ListNode, l2 *ListNode) *ListNode {
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
