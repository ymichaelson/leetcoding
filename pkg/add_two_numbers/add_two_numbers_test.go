package add_two_numbers

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	num1a = ListNode{
		Val:  2,
		Next: &num1b,
	}
	num1b = ListNode{
		Val:  4,
		Next: &num1c,
	}
	num1c = ListNode{
		Val:  3,
		Next: nil,
	}

	num2a = ListNode{
		Val:  5,
		Next: &num2b,
	}
	num2b = ListNode{
		Val:  6,
		Next: &num2c,
	}
	num2c = ListNode{
		Val:  4,
		Next: nil,
	}

	input1  = &num1a
	input2  = &num2a
	expect1 = "342+465=807 7->0->8"
)

func TestAddTwoNumbers(t *testing.T) {
	utils.InitKlog()

	output1 := AddTwoNumbers(input1, input2)

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, toList(output1))
}
