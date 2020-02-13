package logic

import "fmt"

/*
#15 三数之和
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestThreeSum 测试 #15三数之和
func TestThreeSum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("#15 TestThreeSum Input:")
	fmt.Println(nums)
	res := threeSum(nums)
	fmt.Println("#15 TestThreeSum Res:")
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	QuickSort(nums, 0, len(nums)-1)
	resList := make([][]int, 0, 0)
	for k, num1 := range nums {
		if num1 > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue // 去重
		}
		left := k + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + num1
			if sum == 0 {
				res := make([]int, 0)
				res = append(res, num1, nums[left], nums[right])
				resList = append(resList, res)
				//左坐标去重
				for left < right && nums[left] == nums[left+1] {
					left = left + 1
				}
				//右坐标去重
				for left < right && nums[right] == nums[right-1] {
					right = right - 1
				}
				left = left + 1
				right = right - 1
			} else if sum > 0 {
				right = right - 1
			} else if sum < 0 {
				left = left + 1
			}
		}
	}
	return resList
}

//QuickSort 快速排序
func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(arr, startIndex, endIndex)
	QuickSort(arr, startIndex, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, endIndex)
}

func partition(arr []int, startIndex, endIndex int) int {
	var (
		pivot = arr[startIndex]
		left  = startIndex
		right = endIndex
	)
	for left != right {
		for left < right && pivot < arr[right] {
			right--
		}
		for left < right && pivot >= arr[left] {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[startIndex], arr[left] = arr[left], arr[startIndex]
	return left
}
