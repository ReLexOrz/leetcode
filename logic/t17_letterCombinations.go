package logic

import (
	"fmt"
	"strconv"
)

/*
#17 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

(1 !@#) (2 abc) (3 def)
(4 ghi) (5 jkl) (6 mno)
(7 pqrs) (8 tuv) (9 wxyz)
示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestLetterCombinations 测试 #16盛最多水的容器
func TestLetterCombinations() {
	digits := "23"
	fmt.Println("#17 TestLetterCombinations Input:")
	fmt.Println(digits)
	res := letterCombinations(digits)
	fmt.Println("#17 TestLetterCombinations Res:")
	fmt.Println(res)
}

func letterCombinations(digits string) []string {

	phoneList := [][]string{
		[]string{""},
		[]string{""},
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
		[]string{"m", "n", "o"},
		[]string{"p", "q", "r", "s"},
		[]string{"t", "u", "v"},
		[]string{"w", "x", "y", "z"},
	}

	//对每个字符做数字转换 并不断追加结果集
	resList := make([]string, 0, 0)
	for i := 0; i < len(digits); i++ {
		numStr := digits[i : i+1]
		num, _ := strconv.Atoi(numStr)
		//对第一个字符循环追加 将结果集扩充
		curResLen := len(resList)
		for j := 0; j < curResLen; j++ {
			for m := 0; m < len(phoneList[num]); m++ {
				resList = append(resList, resList[j]+phoneList[num][m])
			}
		}
		//删除之前的结果集
		resList = resList[curResLen:len(resList)]
		//之前的结果集本来就是空集 初始化一下
		if len(resList) == 0 {
			for _, str := range phoneList[num] {
				resList = append(resList, str)
			}
		}
	}

	return resList
}
