package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"runtime"
	"syscall"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
)

func getXfrmState(src string, dst string, spi int, key string) *netlink.XfrmState {
	k, _ := hex.DecodeString(key)
	return &netlink.XfrmState{
		Src:   net.ParseIP(src),
		Dst:   net.ParseIP(dst),
		Proto: netlink.XFRM_PROTO_ESP,
		Mode:  netlink.XFRM_MODE_TUNNEL,
		Spi:   spi,
		Ifid:  1,
		Aead: &netlink.XfrmStateAlgo{
			Name:   "rfc4106(gcm(aes))",
			Key:    k,
			ICVLen: 64,
		},
	}
}

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

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ns, err := netns.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create netns: %v\n", err)
		os.Exit(1)
	}
	defer ns.Close()

	err = netlink.XfrmStateAdd(getXfrmState("10.0.0.1", "10.0.0.2", 2, "611d0c8049dd88600ec4f9eded7b1ed540ea607f"))
	if err == nil {
		fmt.Printf("===PIZZA 🍕️\n")
	}

	if os.Getpid() == 1 {
		err = syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
		if err != nil {
			fmt.Fprintf(os.Stderr, "power off failed: %v\n", err)
		}
	}
}
