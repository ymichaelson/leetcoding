package search_range

/*
34.在排序数组中查找元素的第一个和最后一个位置
	题目：
		给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。如果数组中不存在目标值 target，返回 [-1, -1]。

	示例1:
		输入：nums = [5,7,7,8,8,10], target = 8 输出：[3,4]

	示例2:
		输入：nums = [5,7,7,8,8,10], target = 6 输出：[-1,-1]

	示例3:
		输入：nums = [], target = 0 输出：[-1,-1]

	解题思路：

*/

func SearchRange(nums []int, target int) []int {
	first, last := -1, -1
	for i, v := range nums {
		if v == target {
			if first == -1 {
				first = i
			}

			last = i
		}
	}
	return []int{first, last}
}
