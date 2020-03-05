package logic

import "fmt"

/*
#18 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestFourSum() {
	nums := []int{0, 0, 0, 0}
	target := 0
	fmt.Println("#18 TestFourSum Input:")
	fmt.Println(nums)
	res := fourSum(nums, target)
	fmt.Println("#18 TestFourSum Res:")
	fmt.Println(res)
}

func fourSum(nums []int, target int) [][]int {
	//特情
	resList := make([][]int, 0, 0)
	if len(nums) < 3 {
		return resList
	}
	//快排一波
	QuickSort(nums, 0, len(nums)-1)

	//四坐标步进法
	for i := 0; i < len(nums); i++ {
		//进行重复跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			//进行重复跳过
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			start := j + 1
			end := len(nums) - 1
			for {
				//进行重复跳过
				if start > j+1 {
					for start < end && nums[start] == nums[start-1] {
						start++
					}
				}
				if end < len(nums)-1 {
					for start < end && nums[end] == nums[end+1] {
						end--
					}
				}
				//测试条件
				if start >= end {
					break
				}
				sum := nums[i] + nums[j] + nums[start] + nums[end]
				if sum == target {
					resList = append(resList, []int{nums[i], nums[j], nums[start], nums[end]})
				}
				if sum > target {
					end--
				} else {
					start++
				}
			}
		}
	}
	return resList
}

//待优化 //TODO
