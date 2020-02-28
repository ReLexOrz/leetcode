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
	funcMap[4] = logic.TestFindMedianSortedArrays
	funcMap[5] = logic.TestLongestPalindrome
	funcMap[6] = logic.TestConvertStrToZ
	funcMap[7] = logic.TestReverse
	funcMap[8] = logic.TestMyAtoi
	funcMap[9] = logic.TestIsPalindrome
	funcMap[10] = logic.TeatIsMatch
	funcMap[11] = logic.TestMaxArea
	funcMap[12] = logic.TestIntToRoman
	funcMap[13] = logic.TestRomanToInt
	funcMap[15] = logic.TestThreeSum

	//输入你想运行的测试题编号即可返回测试用例
	testNum := 13
	fn, ok := funcMap[testNum]
	if ok {
		fn()
	} else {
		fmt.Println("尚且没有该题目解答")
	}
}
