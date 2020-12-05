package partition

/*
86.分隔链表
	题目：
		给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
		你应当保留两个分区中每个节点的初始相对位置。

	示例1:
		输入: head = 1->4->3->2->5->2, x = 3
		输出: 1->2->2->4->3->5

	解题思路：
	https://leetcode-cn.com/problems/partition-list/solution/fen-ge-lian-biao-by-leetcode/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func Partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	before := &ListNode{}
	beforeHead := before
	after := &ListNode{}
	afterHead := after

	start := head

	for start != nil {
		if start.Val >= x {
			after.Next = start
			after = after.Next
		} else {
			before.Next = start
			before = before.Next
		}
		start = start.Next
	}

	// merge
	before.Next = afterHead.Next
	after.Next = nil

	newHead := beforeHead.Next
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
