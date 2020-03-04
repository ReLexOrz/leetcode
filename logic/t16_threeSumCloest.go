package logic

import (
	"fmt"
)

/*
#16 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestThreeSumClosest 测试 #16盛最多水的容器
func TestThreeSumClosest() {
	nums := []int{1, 2, 4, 8, 16, 32, 64, 128}
	target := 82
	fmt.Println("#16 TestThreeSumClosest Input:")
	fmt.Println(nums)
	res := threeSumClosest(nums, target)
	fmt.Println("#16 TestThreeSumClosest Res:")
	fmt.Println(res)
}

func threeSumClosest(nums []int, target int) int {
	//特情
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= 3 {
		res := 0
		for _, num := range nums {
			res += num
		}
		return res
	}

	//快排一波
	QuickSort(nums, 0, len(nums)-1)

	//三坐标步进法
	resSum := nums[0] + nums[1] + nums[2]
	for idx := 0; idx < len(nums); idx++ {
		start := idx + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[idx] + nums[start] + nums[end]
			if calDis(sum, target) <= calDis(resSum, target) {
				resSum = sum
			}
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return sum
			}
		}
	}
	return resSum
}

//求距离
func calDis(sum int, target int) int {
	dis := sum - target
	if dis < 0 {
		return -dis
	}
	return dis
}
