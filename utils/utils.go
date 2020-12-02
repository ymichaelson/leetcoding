package utils

import (
	"flag"

	"github.com/ymichaelson/klog"
)

func InitKlog() {
	klog.InitFlags(nil)
	flag.Parse()

	defer klog.Flush()
}
