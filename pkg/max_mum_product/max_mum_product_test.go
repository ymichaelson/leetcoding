package max_mum_product

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	input1  = []int{1, 2, 3}
	expect1 = 6

	input2  = []int{-1, -2, -3, -4}
	expect2 = -6
)

func TestMaxmumProduct(t *testing.T) {
	utils.InitKlog()

	output1 := MaximumProduct(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}

	output2 := MaximumProduct(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)
}
