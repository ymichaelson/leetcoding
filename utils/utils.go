package utils

import (
	"flag"

	"github.com/hex-techs/klog"
)

func InitKlog() {
	klog.InitFlags(nil)
	flag.Parse()

	defer klog.Flush()
}
