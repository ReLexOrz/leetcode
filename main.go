package main

import (
	"LeetCodeTest/logic"
	"fmt"
)

func main() {

	funcMap := make(map[string]func())

	println("LeetCode Test")

	//练习题
	funcMap["#1"] = logic.TestTwoSum
	funcMap["#2"] = logic.TestAddTwoNumbers
	funcMap["#3"] = logic.TestLengthOfLongestSubstring
	funcMap["#4"] = logic.TestFindMedianSortedArrays
	funcMap["#5"] = logic.TestLongestPalindrome
	funcMap["#6"] = logic.TestConvertStrToZ
	funcMap["#7"] = logic.TestReverse
	funcMap["#8"] = logic.TestMyAtoi
	funcMap["#9"] = logic.TestIsPalindrome
	funcMap["#10"] = logic.TeatIsMatch
	funcMap["#11"] = logic.TestMaxArea
	funcMap["#12"] = logic.TestIntToRoman
	funcMap["#13"] = logic.TestRomanToInt
	funcMap["#14"] = logic.TestLongestCommonPrefix
	funcMap["#15"] = logic.TestThreeSum
	funcMap["#16"] = logic.TestThreeSumClosest
	funcMap["#17"] = logic.TestLetterCombinations
	funcMap["#18"] = logic.TestFourSum

	//面试题
	funcMap["*57_2"] = logic.TestFindContinuousSequence

	//输入你想运行的测试题编号即可返回测试用例
	testStr := "*57_2"
	fn, ok := funcMap[testStr]
	if ok {
		fn()
	} else {
		fmt.Println("尚且没有该题目解答")
	}
}
