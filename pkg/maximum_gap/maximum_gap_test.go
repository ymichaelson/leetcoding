package maximum_gap

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	input1  = []int{3, 6, 9, 1}
	expect1 = 3

	input2  = []int{10}
	expect2 = 0
)

func TestMaximumGap(t *testing.T) {
	utils.InitKlog()

	output1 := MaximumGap(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)

	output2 := MaximumGap(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)

}
