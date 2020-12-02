package gcd_strings

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

const (
	input1 = "ABCABCABC"
	input2 = "ABC"
	input3 = "AB"
	input4 = "CD"

	expect1 = "ABC"
	expect2 = ""
	expect3 = ""
	expect4 = "ABCABCABC"
)

func TestGcdOfStrings(t *testing.T) {
	utils.InitKlog()

	output1 := GcdOfStrings(input1, input2)
	if output1 != expect1 {
		t.Errorf("test failed, expect %s, but got %s", expect1, output1)
	}
	output2 := GcdOfStrings(input1, input3)
	if output2 != expect2 {
		t.Errorf("test failed, expect %s, but got %s", expect2, output2)
	}
	output3 := GcdOfStrings(input1, input4)
	if output3 != expect3 {
		t.Errorf("test failed, expect %s, but got %s", expect3, output3)
	}
	output4 := GcdOfStrings(input1, input1)
	if output4 != expect4 {
		t.Errorf("test failed, expect %s, but got %s", expect4, output4)
	}

	klog.Infof("input1: %s; input2: %s; expect1: %s; output1: %s", input1, input2, expect1, output1)
	klog.Infof("input1: %s; input3: %s; expect2: %s; output2: %s", input1, input3, expect2, output2)
	klog.Infof("input1: %s; input4: %s; expect3: %s; output3: %s", input1, input4, expect3, output3)
	klog.Infof("input1: %s; input1: %s; expect4: %s; output4: %s", input1, input1, expect4, output4)
}
