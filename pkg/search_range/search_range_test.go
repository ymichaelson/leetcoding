package search_range

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	input1  = []int{5, 7, 7, 8, 8, 10}
	expect1 = []int{3, 4}

	input2  = []int{5, 7, 7, 8, 8, 10}
	expect2 = []int{-1, -1}
)

func TestSearchRange(t *testing.T) {
	utils.InitKlog()

	output1 := SearchRange(input1, 8)

	output2 := SearchRange(input2, 6)

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)
}
