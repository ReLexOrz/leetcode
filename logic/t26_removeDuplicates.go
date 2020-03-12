package logic

import "fmt"

/*
#26 删除排序数组中的重复项
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。



示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。


说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestRemoveDuplicates 测试 #26删除排序数组中的重复项
func TestRemoveDuplicates() {
	nums := []int{1, 1, 2, 3, 3}
	fmt.Println("#26 TestRemoveDuplicates Input:")
	fmt.Println(nums)
	res := removeDuplicates(nums)
	fmt.Println("#26 TestRemoveDuplicates Res:")
	fmt.Println(nums)
	fmt.Println(res)
}

//速度10 空间100
func removeDuplicates3(nums []int) int {
	//特情 无需任何操作
	if len(nums) <= 1 {
		return 1
	}
	for i := 1; ; {
		if i >= len(nums) {
			break
		}
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func removeDuplicates2(nums []int) int {
	//特情 无需任何操作
	for i, j := 0, 0; ; {
		if i+1 >= len(nums) {
			return len(nums)
		}
		for {
			j++
			if j >= len(nums) {
				nums = nums[:i+1]
				return len(nums)
			}
			if nums[j] == nums[i] {
				continue
			} else {
				nums = append(nums[:i+1], nums[j:]...)
				break
			}
		}
		i++
		j = i
	}
}

//直接操作数 不考虑切割 分值不稳定
func removeDuplicates(nums []int) int {
	//特情 无需任何操作
	listLen := len(nums)
	if listLen == 0 {
		return 0
	}
	i := 0
	for j := 1; j < listLen; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
