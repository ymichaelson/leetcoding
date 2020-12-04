package reverse_words

import (
	"strings"
)

/*
151.翻转字符串里的单词
	题目：
		给定一个字符串，逐个翻转字符串中的每个单词。
		说明：
				无空格字符构成一个 单词 。
				输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
				如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

	示例1:
		输入："the sky is blue"
		输出："blue is sky the"

	示例2:
		输入："  hello world!  "
		输出："world! hello"
		解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

	示例3:
		输入："a good   example"
		输出："example good a"
		解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

	解题思路：

*/

func ReverseWords(s string) string {
	strs := strings.Split(s, " ")
	strMap := make(map[int]string, len(s))
	for i, v := range strs {
		if len(v) > 0 && v != " " {
			strMap[i] = v
		}
	}

	var tmp []string
	for i := len(strs); i >= 0; i-- {
		if v, ok := strMap[i]; ok {
			tmp = append(tmp, v)
		}
	}
	return strings.Join(tmp, " ")
}
