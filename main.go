package main

import (
	"LeetCodeTest/logic"
	"fmt"
)

func main() {

	funcMap := make(map[string]func())

	println("LeetCode Test")

	//练习题
	funcMap["#1"] = logic.TestTwoSum                   //两数之和
	funcMap["#2"] = logic.TestAddTwoNumbers            //两数相加
	funcMap["#3"] = logic.TestLengthOfLongestSubstring //无重复字符的最长子串
	funcMap["#4"] = logic.TestFindMedianSortedArrays   //寻找两个有序数组的中位数
	funcMap["#5"] = logic.TestLongestPalindrome        //最长回文子串
	funcMap["#6"] = logic.TestConvertStrToZ            //Z字形变换
	funcMap["#7"] = logic.TestReverse                  //整数反转
	funcMap["#8"] = logic.TestMyAtoi                   //字符串转换整数 (atoi)
	funcMap["#9"] = logic.TestIsPalindrome             //回文数
	funcMap["#10"] = logic.TeatIsMatch                 //正则表达式匹配
	funcMap["#11"] = logic.TestMaxArea                 //盛最多水的容器
	funcMap["#12"] = logic.TestIntToRoman              //整数转罗马数字
	funcMap["#13"] = logic.TestRomanToInt              //罗马数字转整数
	funcMap["#14"] = logic.TestLongestCommonPrefix     //最长公共前缀
	funcMap["#15"] = logic.TestThreeSum                //三数之和
	funcMap["#16"] = logic.TestThreeSumClosest         //最接近的三数之和
	funcMap["#17"] = logic.TestLetterCombinations      //电话号码的字母组合
	funcMap["#18"] = logic.TestFourSum                 //四数之和
	funcMap["#19"] = logic.TestRemoveNthFromEnd        //删除链表的倒数第N个节点
	funcMap["#20"] = logic.TestIsValid                 //有效的括号
	funcMap["#21"] = logic.TestMergeTwoLists           //合并两个有序链表
	funcMap["#22"] = logic.TestGenerateParenthesis     //括号生成
	funcMap["#23"] = logic.TestMergeKLists             //合并K个排序链表
	funcMap["#24"] = logic.TestSwapPairs               //两两交换链表中的节点
	funcMap["#25"] = logic.TestReverseKGroup           //K个一组翻转链表
	funcMap["#26"] = logic.TestRemoveDuplicates        //删除排序数组中的重复项
	funcMap["#27"] = logic.TestRemoveElement           //移除元素
	funcMap["#28"] = logic.TestStrStr                  //实现 strStr()
	funcMap["#29"] = logic.TestDivide                  //两数相除
	funcMap["#30"] = logic.TestFindSubstring           //串联所有单词的子串
	//面试题
	funcMap["*57_2"] = logic.TestFindContinuousSequence

	//输入你想运行的测试题编号即可返回测试用例
	testStr := "#30"
	fn, ok := funcMap[testStr]
	if ok {
		fn()
	} else {
		fmt.Println("尚且没有该题目解答")
	}
}
