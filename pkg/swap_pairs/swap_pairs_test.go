package swap_pairs

import (
	"leetcoding/utils"
	"testing"

	"github.com/hex-techs/klog"
)

var (
	head1 = ListNode{
		Val:  1,
		Next: &head2,
	}
	head2 = ListNode{
		Val:  2,
		Next: &head3,
	}
	head3 = ListNode{
		Val:  3,
		Next: &head4,
	}
	head4 = ListNode{
		Val:  4,
		Next: nil,
	}

	head5 = ListNode{
		Val:  1,
		Next: nil,
	}
	// head6 = ListNode{
	// 	Val:  1,
	// 	Next: &head7,
	// }
	// head7 = ListNode{
	// 	Val:  1,
	// 	Next: &head8,
	// }
	// head8 = ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }
	input1  = &head1
	expect1 = []int{2, 1, 4, 3}

	input2  = &head5
	expect2 = []int{1}

	expect3 = []int{}
)

func TestSwapPairs(t *testing.T) {
	utils.InitKlog()

	output1 := toList(SwapPairs(input1))
	output2 := toList(SwapPairs(input2))
	output3 := toList(SwapPairs(nil))

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)
	klog.Infof("input: %v; expect: %v; output: %v", nil, expect3, output3)
}
