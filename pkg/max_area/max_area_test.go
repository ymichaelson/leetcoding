package max_area

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	input1  = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expect1 = 49

	input2  = []int{4, 3, 2, 1, 4}
	expect2 = 16
)

func TestMaxArea(t *testing.T) {
	utils.InitKlog()

	output1 := MaxArea(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}

	output2 := MaxArea(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)
}
