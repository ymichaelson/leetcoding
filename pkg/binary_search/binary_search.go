package binary_search

/*
704.二分查找
	题目：给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，
	     写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

	示例1：
		输入: nums = [-1,0,3,5,9,12], target = 9
		输出: 4
		解释: 9 出现在 nums 中并且下标为 4
	示例2：
		输入: nums = [-1,0,3,5,9,12], target = 2
		输出: -1
		解释: 2 不存在 nums 中因此返回 -1

	解题思路：
		使用二分查找的方法即可
*/

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := (right + left) / 2

	for left <= right {
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}

		middle = (right + left) / 2
	}
	return -1
}
