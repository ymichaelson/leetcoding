package gcd_strings

/*
1071.题目：对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。
	 返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。

	 示例 1：
	 输入：str1 = "ABCABC", str2 = "ABC"
	 输出："ABC"

	 示例 2：
	 输入：str1 = "ABABAB", str2 = "ABAB"
	 输出："AB"

	 示例 3：
	 输入：str1 = "LEET", str2 = "CODE"
	 输出：""

解题思路：
	如果两个字符串有一个最大公约字符串，那么就会满足str1+str2=str2+str1，如果不满足这个条件则
	没有最大公约字符串。当满足这个条件的时候，则可以使用gcd【欧几里得算法又称辗转相除法】的方法来
	找出最大公约字符串。
*/

func GcdOfStrings(str1 string, str2 string) string {
	if (str1 + str2) != (str2 + str1) {
		return ""
	}

	return str1[0:gcdString(len(str1), len(str2))]
}

func gcdString(a, b int) int {
	if b == 0 {
		return a
	}

	return gcdString(b, a%b)
}
