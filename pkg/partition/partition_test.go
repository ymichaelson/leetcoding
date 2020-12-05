package partition

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	node1 = ListNode{
		Val:  1,
		Next: &node2,
	}
	node2 = ListNode{
		Val:  4,
		Next: &node3,
	}
	node3 = ListNode{
		Val:  3,
		Next: &node4,
	}
	node4 = ListNode{
		Val:  2,
		Next: &node5,
	}
	node5 = ListNode{
		Val:  5,
		Next: &node6,
	}
	node6 = ListNode{
		Val:  2,
		Next: nil,
	}
	input1  = &node1
	expect1 = []int{1, 2, 2, 4, 3, 5}
)

func TestPartition(t *testing.T) {
	utils.InitKlog()

	output1 := Partition(input1, 3)
	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, toList(output1))
}
