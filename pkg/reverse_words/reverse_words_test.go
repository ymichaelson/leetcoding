package reverse_words

import (
	"leetcoding/utils"
	"testing"

	"github.com/hex-techs/klog"
)

var (
	input1  = "the sky is blue"
	expect1 = "blue is sky the"

	input2  = "  hello world!  "
	expect2 = "world! hello"

	input3  = "a good   example"
	expect3 = "example good a"
)

func TestReverseWords(t *testing.T) {
	utils.InitKlog()

	output1 := ReverseWords(input1)
	if output1 != expect1 {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}

	output2 := ReverseWords(input2)
	if output2 != expect2 {
		t.Errorf("test failed, expect %v, but got %v", expect2, output2)
		return
	}

	output3 := ReverseWords(input3)
	if output3 != expect3 {
		t.Errorf("test failed, expect %v, but got %v", expect3, output3)
		return
	}

	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)
	klog.Infof("input: %v; expect: %v; output: %v", input2, expect2, output2)
	klog.Infof("input: %v; expect: %v; output: %v", input3, expect3, output3)
}
