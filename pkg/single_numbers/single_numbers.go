package single_numbers

import "github.com/ymichaelson/klog"

/*
56. 剑指 Offer 56 - I. 数组中数字出现的次数
	题目：
		一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

	示例1:
		输入：nums = [4,1,4,6]
		输出：[1,6] 或 [6,1]

	示例2:
		输入：nums = [1,2,10,4,1,4,3,3]
		输出：[2,10] 或 [10,2]

	解题思路：

*/

func SingleNumbers(nums []int) []int {
	keyMap := make(map[int]int)
	var res []int
	for _, v := range nums {
		if _, ok := keyMap[v]; !ok {
			keyMap[v] = 0
		} else {
			delete(keyMap, v)
		}
	}
	for k, _ := range keyMap {
		res = append(res, k)
	}
	return res
}

func SingleNumbers2(nums []int) []int {
	var a int
	for i := range nums {
		a ^= nums[i]
		klog.Info(a)
	}
	klog.Info(a)
	mask := a & (-a)
	klog.Info(mask)
	res := make([]int, 2)
	for _, v := range nums {
		if (v & mask) == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
