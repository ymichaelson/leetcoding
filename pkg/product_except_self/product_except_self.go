package product_except_self

/*
238.除自身以外数组的乘积
	题目：
		给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

	示例1:
		输入: [1,2,3,4]
		输出: [24,12,8,6]

	解题思路：
	https://leetcode-cn.com/problems/product-of-array-except-self/solution/chu-zi-shen-yi-wai-shu-zu-de-cheng-ji-by-leetcode-/
*/

func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	results := make([]int, length)

	results[0] = 1
	for i := 1; i < length; i++ {
		results[i] = results[i-1] * nums[i-1]
	}

	right := 1
	for j := length - 1; j >= 0; j-- {
		results[j] = results[j] * right
		right *= nums[j]
	}
	return results
}
