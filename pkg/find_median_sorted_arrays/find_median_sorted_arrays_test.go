package find_median_sorted_arrays

import (
	"leetcoding/utils"
	"testing"

	"github.com/hex-techs/klog"
)

var (
	input1  = []int{-2, -1}
	input1b = []int{3}
	expect1 = -1.00000

	input2  = []int{1, 2}
	input2b = []int{3, 4}
	expect2 = 2.50000

	input3  = []int{0, 0}
	input3b = []int{0, 0}
	expect3 = 0.00000

	input4  = []int{}
	input4b = []int{1}
	expect4 = 1.00000

	input5  = []int{2}
	input5b = []int{}
	expect5 = 2.00000
)

func TestFindMedianSortedArrays(t *testing.T) {
	utils.InitKlog()

	output1 := FindMedianSortedArrays(input1, input1b)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}
	klog.Infof("input: %v %v; expect: %v; output: %v", input1, input1b, expect1, output1)

	output2 := FindMedianSortedArrays(input2, input2b)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}
	klog.Infof("input: %v %v; expect: %v; output: %v", input2, input2b, expect2, output2)

	output3 := FindMedianSortedArrays(input3, input3b)
	if output3 != expect3 {
		t.Errorf("test failed, expect %v, but got %v", expect3, output3)
		return
	}
	klog.Infof("input: %v %v; expect: %v; output: %v", input3, input3b, expect3, output3)

	output4 := FindMedianSortedArrays(input4, input4b)
	if output4 != expect4 {
		t.Errorf("test failed, expect %v, but got %v", expect4, output4)
		return
	}
	klog.Infof("input: %v %v; expect: %v; output: %v", input4, input4b, expect4, output4)

	output5 := FindMedianSortedArrays(input5, input5b)
	if output5 != expect5 {
		t.Errorf("test failed, expect %v, but got %v", expect5, output5)
		return
	}
	klog.Infof("input: %v %v; expect: %v; output: %v", input5, input5b, expect5, output5)
}
