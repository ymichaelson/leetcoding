package stairs

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

const (
	input  = 11
	expect = 144
)

func TestClimbStairs(t *testing.T) {
	utils.InitKlog()

	output := ClimbStairs(input)

	if output != expect {
		t.Errorf("test failed, expect %d, but got %d", expect, output)
	}

	klog.Infof("input: %d; expect: %d; output: %d", input, expect, output)
}
