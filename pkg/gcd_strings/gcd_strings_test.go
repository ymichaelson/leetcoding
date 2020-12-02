package gcd_strings

import (
	"leetcoding/utils"
	"testing"

	"github.com/ymichaelson/klog"
)

func TestGcdStrings(t *testing.T) {
	utils.InitKlog()

	var testString1 = "ABCABCABC"
	var testString2 = "ABC"
	var testString3 = "AB"
	var testString4 = "CD"

	klog.Infoln("string1 with string2: ", GcdOfStrings(testString1, testString2))
	klog.Infoln("string1 with string3: ", GcdOfStrings(testString1, testString3))
	klog.Infoln("string1 with string4: ", GcdOfStrings(testString1, testString4))
}
