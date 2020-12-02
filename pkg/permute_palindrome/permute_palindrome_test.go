package permute_palindrome

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

const (
	input1  = "asdasd"
	expect1 = true

	input2  = "asdasd"
	expect2 = true

	input3  = "asdasdsm"
	expect3 = false
)

func TestCanPermutePalindrome(t *testing.T) {
	utils.InitKlog()

	output1 := CanPermutePalindrome(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %t, but got %t", expect1, output1)
	}
	output2 := CanPermutePalindrome(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %t, but got %t", expect2, output2)
	}
	output3 := CanPermutePalindrome(input3)
	if output3 != expect3 {
		t.Errorf("test failed, expect %t, but got %t", expect3, output3)
	}

	klog.Infof("input1: %s; expect1: %t; output1: %t", input1, expect1, output1)
	klog.Infof("input2: %s; expect2: %t; output2: %t", input2, expect2, output2)
	klog.Infof("input3: %s; expect3: %t; output3: %t", input3, expect3, output3)
}
