package stairs

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

func TestClimbStairs(t *testing.T) {
	utils.InitKlog()

	klog.Infoln("climb stairs: ", ClimbStairs(5))
}
