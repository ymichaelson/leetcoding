package lemonade_change

import (
	"leetcoding/utils"
	"testing"

	"github.com/hex-techs/klog"
)

var (
	input1  = []int{5, 5, 5, 10, 20}
	expect1 = true

	input2  = []int{5, 5, 10}
	expect2 = true

	input3  = []int{5, 5, 10, 10, 20}
	expect3 = false
)

func TestLemonadeChange(t *testing.T) {
	utils.InitKlog()

	output1 := LemonadeChange(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)

	output2 := LemonadeChange(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)

	output3 := LemonadeChange(input3)
	if output3 != expect3 {
		t.Errorf("test failed, expect %v, but got %v", expect3, output3)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input3, expect3, output3)

}
