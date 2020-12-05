package find_median_sorted_arrays

/*
4.寻找两个正序数组的中位数
	题目：
		给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

	示例1:
		输入：nums1 = [1,3], nums2 = [2]
		输出：2.00000
		解释：合并数组 = [1,2,3] ，中位数 2

	示例2:
		输入：nums1 = [1,2], nums2 = [3,4]
		输出：2.50000
		解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

	示例3:
		输入：nums1 = [0,0], nums2 = [0,0]
		输出：0.00000

	示例4:
		输入：nums1 = [], nums2 = [1]
		输出：1.00000

	示例5:
		输入：nums1 = [2], nums2 = []
		输出：2.00000

	解题思路：

*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenght1 := len(nums1)
	lenght2 := len(nums2)
	lp := 0
	rp := 0
	var merge []int

	for {
		if lp >= lenght1 {
			merge = append(merge, nums2[rp:]...)
			break
		}

		if rp >= lenght2 {
			merge = append(merge, nums1[lp:]...)
			break
		}

		if i, ok := compare(nums1[lp], nums2[rp]); ok {
			merge = append(merge, i)
			rp++
		} else {
			merge = append(merge, i)
			lp++
		}
	}

	var median float64
	if len(merge)%2 == 0 {
		median = float64(merge[len(merge)/2-1]+merge[len(merge)/2]) / 2
	} else {
		median = float64(merge[len(merge)/2])
	}

	return median
}

// a left, b right; false means need move left pointer,
// or move right pointer.
func compare(a, b int) (int, bool) {
	if a > b {
		return b, true
	}
	return a, false
}
