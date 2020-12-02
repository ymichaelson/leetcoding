package permute_palindrome

/*
面试题 01.04. 题目：给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
				  回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
				  回文串不一定是字典当中的单词。

				  示例1：
				  输入："tactcoa"
				  输出：true（排列有"tacocat"、"atcocta"，等等）

解题思路：
	根据字符串的长度，将字符串分奇数和偶数两种情况看。然后根据消消乐的原理，逐一消减字符串中
	出现两次的字符，使用map来存每一个字符，字符为key，如果遇到已经存在的key，则删除掉对应的key，
	如果字符串长度是偶数，则最后map长度为0，就是回文字符串，如果字符串长度是奇数，则最后map长度
	为1，就是回文字符串。
*/

func CanPermutePalindrome(s string) bool {
	exist := make(map[rune]int)
	for _, i := range s {
		if _, ok := exist[i]; !ok {
			exist[i] = 0
		} else {
			delete(exist, i)
		}
	}

	if len(s)%2 == 0 {
		if len(exist) == 0 {
			return true
		}
	} else {
		if len(exist) == 1 {
			return true
		}
	}
	return false
}
