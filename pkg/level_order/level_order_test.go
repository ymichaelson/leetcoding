package level_order

import (
	"leetcoding/utils"
	"reflect"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	head = TreeNode{
		Val:   3,
		Left:  &head2,
		Right: &head3,
	}
	head2 = TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	head3 = TreeNode{
		Val:   20,
		Left:  &head4,
		Right: &head5,
	}
	head4 = TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	head5 = TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	input1  = &head
	expect1 = [][]int{
		{3},
		{9, 20},
		{15, 7},
	}
)

func TestLevelOrder(t *testing.T) {
	utils.InitKlog()

	output1 := LevelOrder(input1)
	if !reflect.DeepEqual(output1, expect1) {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)

}
