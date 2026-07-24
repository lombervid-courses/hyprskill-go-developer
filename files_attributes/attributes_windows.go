//go:build windows

package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func osStats(f os.FileInfo) {
	stat, ok := f.Sys().(*syscall.Win32FileAttributeData)
	if ok {
		fmt.Println(time.Unix(0, stat.CreationTime.Nanoseconds())) // 2022-03-10 11:15:26 -0500 -05
		fmt.Println("Size:", stat.FileSizeLow, "bytes")            // Size: 35 bytes
	}
}
