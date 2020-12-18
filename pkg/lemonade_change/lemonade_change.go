package lemonade_change

import "github.com/hex-techs/klog"

/*
860.柠檬水找零
	题目：
		在柠檬水摊上，每一杯柠檬水的售价为 5 美元。
		顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
		每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
		注意，一开始你手头没有任何零钱。
		如果你能给每位顾客正确找零，返回 true ，否则返回 false 。

	示例1:
		输入：[5,5,5,10,20]
		输出：true
		解释：
		前 3 位顾客那里，我们按顺序收取 3 张 5 美元的钞票。
		第 4 位顾客那里，我们收取一张 10 美元的钞票，并返还 5 美元。
		第 5 位顾客那里，我们找还一张 10 美元的钞票和一张 5 美元的钞票。
		由于所有客户都得到了正确的找零，所以我们输出 true。

	示例2:
		输入：[5,5,10]
		输出：true

	示例3:
		输入：[5,5,10,10,20]
		输出：false
		解释：
		前 2 位顾客那里，我们按顺序收取 2 张 5 美元的钞票。
		对于接下来的 2 位顾客，我们收取一张 10 美元的钞票，然后返还 5 美元。
		对于最后一位顾客，我们无法退回 15 美元，因为我们现在只有两张 10 美元的钞票。
		由于不是每位顾客都得到了正确的找零，所以答案是 false。

	解题思路：

*/

func LemonadeChange(bills []int) bool {
	if len(bills) < 0 {
		return true
	}

	if bills[0] > 5 {
		return false
	}

	// map: bill -> count
	billMap := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}
	for _, bill := range bills {
		klog.Info(billMap)
		klog.Info(bill)
		if bill == 5 {
			billMap[5] += 1
			continue
		}

		if !change(bill, billMap) {
			return false
		}
		klog.Info(billMap)
	}

	return true
}

func change(bill int, billMap map[int]int) bool {
	if bill == 10 {
		if v := billMap[5]; v >= 1 {
			billMap[5] -= 1
			billMap[10] += 1
			return true
		}
	}

	if bill == 20 {
		v5 := billMap[5]
		v10 := billMap[10]

		if v10 >= 1 && v5 >= 1 {
			billMap[10] -= 1
			billMap[5] -= 1
			billMap[20] += 1
			return true
		}

		if v5 >= 3 {
			billMap[5] -= 3
			billMap[20] += 1
			return true
		}
	}
	return false
}
