package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"syscall"
)

func main() {
	if os.Getpid() == 1 {
		if err := os.Mkdir("/proc", 0755); err != nil {
			fmt.Fprintf(os.Stderr, "mkdir: %v\n", err)
			os.Exit(1)
		}
		if err := os.Mkdir("/sys", 0755); err != nil {
			fmt.Fprintf(os.Stderr, "mkdir: %v\n", err)
			os.Exit(1)
		}

		if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
			fmt.Fprintf(os.Stderr, "mount: %v\n", err)
			os.Exit(1)
		}

		if err := syscall.Mount("sysfs", "/sys", "sysfs", 0, ""); err != nil {
			fmt.Fprintf(os.Stderr, "mount: %v\n", err)
			os.Exit(1)
		}
	}

	cpulist, err := os.ReadFile("/sys/devices/system/node/node0/cpulist")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file: %v\n", err)
		os.Exit(1)
	} else {
		hd := hex.Dump(cpulist)
		if c := cpulist[len(cpulist)-1]; c == '\n' {
			fmt.Printf("===PIZZA 🍕️ - %s\n", hd)
		} else {
			fmt.Printf("===PINEAPPLE 🍍️ - %s\n", hd)
		}
	}

	if os.Getpid() == 1 {
		err = syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
		if err != nil {
			fmt.Fprintf(os.Stderr, "power off failed: %v\n", err)
		}
	}
}
