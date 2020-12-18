package erase_overlap_intervals

import (
	"leetcoding/utils"
	"testing"

	"github.com/hex-techs/klog"
)

var (
	input1  = [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	expect1 = 2

	input2  = [][]int{{1, 2}, {1, 2}, {1, 2}}
	expect2 = 2

	input3  = [][]int{{1, 2}, {2, 3}}
	expect3 = 0
)

func TestEraseOverlapIntervals(t *testing.T) {
	utils.InitKlog()

	output1 := EraseOverlapIntervals(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}

	output2 := EraseOverlapIntervals(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}

	output3 := EraseOverlapIntervals(input3)
	if output3 != expect3 {
		t.Errorf("test failed, expect %v, but got %v", expect3, output3)
		return
	}

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)
	klog.Infof("input: %v; expect: %v; output: %v", input3, expect3, output3)
}
