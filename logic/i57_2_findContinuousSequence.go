package logic

import (
	"fmt"
)

/*
*57_2 和为s的连续正数序列
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

示例 1：
输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：
输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]

限制：
1 <= target <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestFindContinuousSequence 面试*57 和为s的连续正数序列
func TestFindContinuousSequence() {
	target := 9
	fmt.Println("*57_2 TestFindContinuousSequence Input:")
	fmt.Println(target)
	res := findContinuousSequence(target)
	fmt.Println("*57_2 TestFindContinuousSequence Res:")
	fmt.Println(res)
}

func findContinuousSequence(target int) [][]int {
	//将target构成一个序列
	resList := make([][]int, 0, 0)
	//特情
	if target <= 2 {
		return resList
	}
	//连续正整数起点和终点
	start := 1
	end := 2
	for end-start >= 1 && start+end <= target {
		sum := calSum(start, end, target, &resList)
		if sum >= target {
			start++
		} else {
			end++
		}
	}
	return resList
}

//求序列和并将结果集加入
func calSum(start int, end int, target int, resList *[][]int) (res int) {
	for i := start; i <= end; i++ {
		res += i
	}
	if res == target {
		resUnit := make([]int, 0, 0)
		for i := start; i <= end; i++ {
			resUnit = append(resUnit, i)
		}
		*resList = append(*resList, resUnit)
	}
	return
}
