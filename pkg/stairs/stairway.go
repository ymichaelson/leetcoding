package stairs

/*
70.题目：假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
	 注意：给定 n 是一个正整数。

	 输入： 3
	 输出： 3
	 解释： 有三种方法可以爬到楼顶。
	 1.  1 阶 + 1 阶 + 1 阶
	 2.  1 阶 + 2 阶
	 3.  2 阶 + 1 阶

解题思路：
	通过推理，可以发现当需要走上第i个台阶的时候，可以有两种情况，可以是i-1直接到i，
	或者i-2直接到i，可以抽象出一个公式就是走到第i个台阶可以有n(i) = n(i-1) + n(i-2)种方法，
	也就是从第三位开始的斐波那切数列，1，2，3，5，8...

*/

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	var tmp = make([]int, n)
	tmp[0] = 1
	tmp[1] = 2
	for i := 2; i < n; i++ {
		tmp[i] = tmp[i-1] + tmp[i-2]
	}

	return tmp[n-1]
}
