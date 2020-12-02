package main

import (
	"flag"
	"leetcoding/pkg/gcd_strings"
	"leetcoding/pkg/permute_palindrome"
	"leetcoding/pkg/stairs"

	"github.com/ymichaelson/klog"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	defer klog.Flush()

	klog.Infoln("climb stairs: ", stairs.ClimbStairs(5))

	klog.Infoln("gcd strings: ", gcd_strings.GcdOfStrings("ABDABDABD", "ABD"))

	klog.Infoln("permute palindrome: ", permute_palindrome.CanPermutePalindrome("asdasd"))
}
