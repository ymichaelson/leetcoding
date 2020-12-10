package maximum_gap

import "sort"

/*
164.最大间距
	题目：
		给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。如果数组元素个数小于 2，则返回 0。

	示例1:
		输入: [3,6,9,1]
		输出: 3
		解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。

	示例2:
		输入: [10]
		输出: 0
		解释: 数组元素个数小于 2，因此返回 0。

	解题思路：

*/

func MaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Sort(sort.IntSlice(nums))

	var max int
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i-1]
		if tmp > max {
			max = tmp
		}
	}

	return max
}
