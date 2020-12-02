package max_mum_product

import (
	"math"
	"sort"
)

/*
628.三个数的最大乘积
	题目：给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

	示例1：
		输入: [1,2,3]
		输出: 6
	示例2：
		输入: [1,2,3,4]
		输出: 24

	解题思路：

*/

func MaximumProduct(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	return int(math.Max(float64(mum(nums[0], nums[1], nums[len(nums)-1])), float64(mum(nums[len(nums)-1], nums[len(nums)-2], nums[len(nums)-3]))))
}

func mum(nums ...int) int {
	var s = 1
	for _, i := range nums {
		s *= i
	}
	return s
}
