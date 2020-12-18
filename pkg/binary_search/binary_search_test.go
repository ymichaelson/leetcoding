package binary_search

import (
	"leetcoding/utils"
	"testing"

	"github.com/hex-techs/klog"
)

var (
	input1  = []int{-1, 0, 3, 5, 9, 12}
	target1 = 9
	expect1 = 4

	input2  = []int{-1, 0, 3, 5, 9, 12}
	target2 = 2
	expect2 = -1
)

func TestSearch(t *testing.T) {
	utils.InitKlog()

	output1 := Search(input1, target1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}

	output2 := Search(input2, target2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}

	klog.Infof("input: %v; target: %v; expect: %v; output: %v", input1, target1, expect1, output1)
	klog.Infof("input: %v; target: %v; expect: %v; output: %v", input2, target2, expect2, output2)
}
