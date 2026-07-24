//go:build unix

/* //go:build !windows */

package main

import (
	"fmt"
	"os"
	"syscall"
)

func osStats(f os.FileInfo) {
	stat, ok := f.Sys().(*syscall.Stat_t)
	if ok {
		fmt.Println("User identifier:", stat.Uid)
		fmt.Println("Group identifier:", stat.Gid)
	}
}
