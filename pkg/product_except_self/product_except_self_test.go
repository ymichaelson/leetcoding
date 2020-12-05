package product_except_self

import (
	"leetcoding/utils"
	"reflect"
	"testing"

	"github.com/ymichaelson/klog"
)

var (
	input1  = []int{1, 2, 3, 4}
	expect1 = []int{24, 12, 8, 6}
)

func TestProductExceptSelf(t *testing.T) {
	utils.InitKlog()

	output1 := ProductExceptSelf(input1)
	if !reflect.DeepEqual(output1, expect1) {
		t.Errorf("test failed, expect %v, but got %v", expect1, output1)
		return
	}
	klog.Infof("input: %v; expect: %v; output: %v", input1, expect1, output1)

}
