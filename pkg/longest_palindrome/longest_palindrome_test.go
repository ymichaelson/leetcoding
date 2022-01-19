package longest_palindrome

import (
	"leetcoding/utils"
	"testing"

	"github.com/hex-techs/klog"
)

var (
	input1  = "babad"
	expect1 = "aba" // "bab"

	input2  = "cbbd"
	expect2 = "bb"
)

func TestLongestPalindrome(t *testing.T) {
	utils.InitKlog()

	output1 := LongestPalindrome(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)

	output2 := LongestPalindrome(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)

}
