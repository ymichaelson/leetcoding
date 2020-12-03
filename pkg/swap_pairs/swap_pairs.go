package swap_pairs

/*
24.两两交换链表中的节点
	题目：给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
		 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

	示例1：
		输入：head = [1,2,3,4]
		输出：[2,1,4,3]
	示例2：
		输入：head = [1]
		输出：[1]
	示例3：
		输入：head = []
		输出：[]

	解题思路：
		交换后的链表头节点是原链表的第二个节点，新链表的第二个节点是原链表的第一个节点，
		至此递归后面的节点两两交换即可。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = SwapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func toList(head *ListNode) []int {
	var list []int
	if head == nil {
		return nil
	}

	list = append(list, head.Val)
	node := head.Next
	for node != nil {
		list = append(list, node.Val)
		node = node.Next
	}

	return list
}
