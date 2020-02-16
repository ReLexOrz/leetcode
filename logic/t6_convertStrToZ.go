package logic

import "fmt"

/*
#6 Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestConvertStrToZ 测试 #6 Z字形变换
func TestConvertStrToZ() {
	s := "PAYPALISHIRING"
	num := 3
	fmt.Println("#6 TestConvertStrToZ Input:")
	fmt.Println(s)
	fmt.Println(3)
	res := convertStrToZ2(s, num)
	fmt.Println("#6 TestConvertStrToZ Res:")
	fmt.Println(res)
}

//在实际题目中func的名字为convert
//非常不好的方法 时间得分60 内存得分17
func convertStrToZ(s string, numRows int) string {
	resStr := ""
	sLen := len(s)
	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		for j := 0; ; j++ {
			fIdx := (numRows-1)*2*j + i
			if fIdx >= sLen {
				break
			}
			resStr = resStr + string(s[fIdx])
			if i > 0 && i < (numRows-1) {
				sIdx := fIdx + (numRows-i-1)*2
				if sIdx >= sLen {
					break
				}
				resStr = resStr + string(s[sIdx])
			}
		}
	}
	return resStr
}

//第二种方法 遍历字串并将合适的Idx丢进合适的数组里 最终整合数组返回
//时间得分79 内存得分25 目前还有优化余地
func convertStrToZ2(s string, n int) string {
	resStr := ""
	sLen := len(s)
	if n == 1 {
		return s
	}

	//创建用于结果的数组
	resList := make([][]byte, 0)
	for i := 0; i < n; i++ {
		resList = append(resList, make([]byte, 0))
	}

	//对字符串进行拆析
	param := 2*n - 2
	for i := 0; i < sLen; i++ {
		subIdx := i % param
		if subIdx < n {
			resList[subIdx] = append(resList[subIdx], s[i])
		} else {
			resList[param-subIdx] = append(resList[param-subIdx], s[i])
		}
	}

	for _, byteList := range resList {
		resStr = resStr + string(byteList)
	}

	return resStr
}
