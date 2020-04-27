package logic

import "fmt"

/*
#30 串联所有单词的子串

给定一个字符串 s 和一些长度相同的单词 words。
找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestFindSubstring 测试 #30 串联所有单词的子串
func TestFindSubstring() {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	fmt.Println("#30 TestFindSubstring Input:")
	fmt.Println(s)
	fmt.Println(words)
	res := findSubstring(s, words)
	fmt.Println("#30 TestFindSubstring Res:")
	fmt.Println(res)
}

// func findSubstring2(s string, words []string) []int {
// 	res := make([]int, 0)
// 	if len(words) == 0 {
// 		return res
// 	}
// 	//单字符串长度
// 	wordLen := len(words[0])
// 	twordLen := len(words) * wordLen
// 	wordsMap := make(map[string]int, 0)
// 	for i := 0; i < len(words); i++ {
// 		wordsMap[words[i]] = 0
// 	}
// 	if len(s) < wordLen {
// 		return res
// 	}

// 	for i, j := 0, wordLen; j <= len(s); {
// 		subStr := s[i:j]
// 		v, ok := wordsMap[subStr]
// 		if ok && v == 0 {
// 			wordsMap[subStr] = 1
// 			for m, n := j, j+wordLen; ; {
// 				v, ok := wordsMap[s[m:n]]
// 				if !ok || v == 1 {
// 					//ClearMapValue()
// 					break
// 				}
// 				if v == 0 {

// 				}
// 			}
// 		}
// 		i++
// 		j++
// 	}
// 	return res
// }

// func ClearMapValue(strMap map[string]int) {
// 	//for
// }

//方法超时
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(words) == 0 {
		return res
	}
	//列出所有单词的组合
	wordLen := len(words[0])
	twordLen := len(words) * wordLen
	numCombinList := GetAllNumCombineList(len(words) - 1)
	strMap := make(map[string]int, 0)
	for i := 0; i < len(numCombinList); i++ {
		var subs string = ""
		for j := 0; j < len(numCombinList[i]); j++ {
			subs = subs + words[numCombinList[i][j]]
		}
		strMap[subs] = 1
	}
	//遍历字符串
	for i := 0; i+twordLen <= len(s); i++ {
		_, ok := strMap[s[i:i+twordLen]]
		if ok {
			res = append(res, i)
		}
	}
	return res
}

func GetAllNumCombineList(maxNum int) [][]int {
	if maxNum == 0 {
		return [][]int{[]int{0}}
	}
	numCombinList := GetAllNumCombineList(maxNum - 1)
	newCombinList := make([][]int, 0, 0)
	for i := 0; i < len(numCombinList); i++ {
		for j := 0; j < len(numCombinList[i]); j++ {
			newList := make([]int, 0, 0)
			newList = append(newList, numCombinList[i][:j]...)
			newList = append(newList, maxNum)
			newList = append(newList, numCombinList[i][j:]...)
			newCombinList = append(newCombinList, newList)
		}
		newCombinList = append(newCombinList, append(numCombinList[i], maxNum))
	}
	return newCombinList
}
