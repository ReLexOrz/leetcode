package logic

import "fmt"

/*
#19 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
进阶：
你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestRemoveNthFromEnd 测试 #19删除链表的倒数第N个节点
func TestRemoveNthFromEnd() {
	l1E := &ListNode{
		Val:  5,
		Next: nil,
	}
	l1D := &ListNode{
		Val:  4,
		Next: l1E,
	}
	l1C := &ListNode{
		Val:  3,
		Next: l1D,
	}
	l1B := &ListNode{
		Val:  2,
		Next: l1C,
	}
	l1A := &ListNode{
		Val:  1,
		Next: l1B,
	}
	n := 2
	fmt.Println("#19 TestRemoveNthFromEnd Input:")
	PrintListNode(l1A)
	res := removeNthFromEnd(l1A, n)
	fmt.Println("#19 TestRemoveNthFromEnd Res:")
	PrintListNode(res)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//找一个List随着遍历记录链表地址
	nodeList := make([]*ListNode, 0, 0)
	curNode := head
	for curNode != nil {
		//记录位置
		nodeList = append(nodeList, curNode)
		curNode = curNode.Next
	}

	//根据n删除相应位置上的node
	if n > len(nodeList) {
		return nil
	}
	//n正好是第一个节点
	if n == len(nodeList) {
		return head.Next
	}
	//删哪个
	i := len(nodeList) - n

	preNode := nodeList[i-1]
	preNode.Next = nodeList[i].Next

	return head
}
