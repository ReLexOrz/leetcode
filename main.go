package main

import (
	"LeetCodeTest/logic"
	"fmt"
)

func main() {

	funcMap := make(map[int]func())

	println("LeetCode Test")

	funcMap[1] = logic.TestTwoSum
	funcMap[2] = logic.TestAddTwoNumbers
	funcMap[3] = logic.TestLengthOfLongestSubstring
	funcMap[7] = logic.TestReverse
	funcMap[15] = logic.TestThreeSum

	//输入你想运行的测试题编号即可返回测试用例
	testNum := 7
	fn, ok := funcMap[testNum]
	if ok {
		fn()
	} else {
		fmt.Println("尚且没有该题目解答")
	}
}
