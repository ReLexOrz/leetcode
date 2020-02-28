package logic

import (
	"fmt"
)

/*
#11 盛最多水的容器

给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestMaxArea 测试 #11盛最多水的容器
func TestMaxArea() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("#11 TestMaxArea Input:")
	fmt.Println(nums)
	res := maxArea(nums)
	fmt.Println("#11 TestMaxArea Res:")
	fmt.Println(res)
}

func maxArea(height []int) int {
	wMaxVol := 0
	//双边界向内扫描
	for i, j := 0, len(height)-1; i < j; {
		minH := 0
		if height[i] >= height[j] {
			minH = height[j]
		} else {
			minH = height[i]
		}
		vol := (j - i) * minH
		if vol > wMaxVol {
			wMaxVol = vol
		}

		//比较矮的一边进行丢弃
		//如果都一样 另一边也可以早早弃了 肯定不如原来结果大 不费劲了
		if minH == height[j] {
			j--
		}
		if minH == height[i] {
			i++
		}
	}
	return wMaxVol
}
