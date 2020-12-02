package max_sub_array

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	input1  = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect1 = 6
)

func TestMaxSubArray(t *testing.T) {
	utils.InitKlog()

	output1 := MaxSubArray(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %d, but got %d", expect1, output1)
		return
	}

	klog.Infof("input: %v; expect: %d; output: %d", input1, expect1, output1)
}
