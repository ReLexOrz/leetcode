package logic

import "fmt"

/*
#4 寻找两个有序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//TestFindMedianSortedArrays 测试 #4寻找两个有序数组的中位数
func TestFindMedianSortedArrays() {
	nums1 := []int{1, 2, 3, 9}
	nums2 := []int{1}
	fmt.Println("#4 TestFindMedianSortedArrays Input:")
	fmt.Println(nums1)
	fmt.Println(nums2)
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println("#4 TestFindMedianSortedArrays Res:")
	fmt.Println(res)

	fmt.Println("#4 TestFindMedianSortedArrays 第二种解法：")

	nums1 = []int{1, 2, 3, 9}
	nums2 = []int{1}
	fmt.Println("#4 TestFindMedianSortedArrays Input:")
	fmt.Println(nums1)
	fmt.Println(nums2)
	res = findMedianSortedArrays2(nums1, nums2)
	fmt.Println("#4 TestFindMedianSortedArrays Res:")
	fmt.Println(res)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//目前此种方法需要多一块内存用于合并两个数组
	//占用空间较高
	mianNums := make([]int, 0, 0)
	numLen1 := len(nums1)
	numLen2 := len(nums2)
	for i, j := 0, 0; ; {
		if i >= numLen1 {
			mianNums = append(mianNums, nums2[j:]...)
			break
		}
		if j >= numLen2 {
			mianNums = append(mianNums, nums1[i:]...)
			break
		}
		if nums1[i] >= nums2[j] {
			mianNums = append(mianNums, nums2[j])
			j++
		} else {
			mianNums = append(mianNums, nums1[i])
			i++
		}
	}
	return quickCalMNum(mianNums)
}

//快速求取中位取值
func quickCalMNum(nums []int) float64 {
	nLen := len(nums)
	if (nLen & 1) == 0 {
		return float64(nums[nLen>>1-1]+nums[nLen>>1]) / 2.0
	}
	return float64(nums[nLen>>1])
}

//第二种方法 按说内存消耗应该比第一种要小的 但怎么还高出来一点点？
//双有序数组冒泡排序 从最后一个开始往前冒 最后加入nums1
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	for _, num := range nums2 {
		nums1 = append(nums1, num)
		for i := len(nums1) - 1; i > 0; {
			if nums1[i] >= nums1[i-1] {
				break
			} else { //向前交换并步进
				nums1[i-1], nums1[i] = nums1[i], nums1[i-1]
				i--
			}
		}
	}
	return quickCalMNum(nums1)
}
