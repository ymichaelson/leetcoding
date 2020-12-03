package single_numbers

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	input1  = []int{4, 1, 4, 6}
	expect1 = []int{1, 6}

	input2  = []int{1, 2, 10, 4, 1, 4, 3, 3}
	expect2 = []int{2, 10}
)

func TestSingleNumbers(t *testing.T) {
	utils.InitKlog()

	output1 := SingleNumbers(input1)
	output2 := SingleNumbers2(input2)

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)
}
