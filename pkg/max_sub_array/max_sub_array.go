package max_sub_array

/*
53.最大子序和
	题目：给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	示例：
	     输入: [-2,1,-3,4,-1,2,1,-5,4]
	     输出: 6
		 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

	解题思路：
	f(i)为第i个元素之前的连续子数组的最大和
	f(i-1)为第i-1个元素之前的连续子数组的最大和
	nums[i]为第i个元素

	抽象公式：f(i) = max(f(i-1)+nums[i], nums[i])
*/

func MaxSubArray(nums []int) int {
	var max = nums[0]
	if len(nums) > 1 {
		for i := 1; i < len(nums); i++ {
			if nums[i]+nums[i-1] > nums[i] {
				nums[i] += nums[i-1]
			}

			if nums[i] > max {
				max = nums[i]
			}
		}
	}
	return max
}
