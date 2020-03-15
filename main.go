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
	funcMap["#19"] = logic.TestRemoveNthFromEnd
	funcMap["#20"] = logic.TestIsValid
	funcMap["#21"] = logic.TestMergeTwoLists
	funcMap["#22"] = logic.TestGenerateParenthesis
	funcMap["#23"] = logic.TestMergeKLists
	funcMap["#24"] = logic.TestSwapPairs
	funcMap["#25"] = logic.TestReverseKGroup
	funcMap["#26"] = logic.TestRemoveDuplicates
	funcMap["#27"] = logic.TestRemoveElement
	funcMap["#28"] = logic.TestStrStr
	//面试题
	funcMap["*57_2"] = logic.TestFindContinuousSequence

	//输入你想运行的测试题编号即可返回测试用例
	testStr := "#28"
	fn, ok := funcMap[testStr]
	if ok {
		fn()
	} else {
		fmt.Println("尚且没有该题目解答")
	}
}
