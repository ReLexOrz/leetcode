package logic

import (
	"fmt"
	"strconv"
)

//PrintListNode 打印int单向列表
func PrintListNode(ln *ListNode) {
	//string := strconv.Itoa(int)
	var resStr = "["
	lp := ln
	for {
		if lp != nil {
			resStr += strconv.Itoa(lp.Val)
		} else {
			resStr += "]"
			break
		}
		lp = lp.Next
		if lp != nil {
			resStr += ","
		}
	}
	fmt.Println(resStr)
}
