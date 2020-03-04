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

//快速排序
func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(arr, startIndex, endIndex)
	QuickSort(arr, startIndex, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, endIndex)
}

func partition(arr []int, startIndex, endIndex int) int {
	var (
		pivot = arr[startIndex]
		left  = startIndex
		right = endIndex
	)
	for left != right {
		for left < right && pivot < arr[right] {
			right--
		}
		for left < right && pivot >= arr[left] {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[startIndex], arr[left] = arr[left], arr[startIndex]
	return left
}
