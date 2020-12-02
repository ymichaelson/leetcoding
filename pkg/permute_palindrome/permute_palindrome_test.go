package permute_palindrome

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

func TestPermutePalindrome(t *testing.T) {
	utils.InitKlog()

	klog.Infoln("permute palindrome: ", CanPermutePalindrome("asdasd"))
	klog.Infoln("permute palindrome: ", CanPermutePalindrome("asdasd"))
	klog.Infoln("permute palindrome: ", CanPermutePalindrome("asdasdsm"))
}
