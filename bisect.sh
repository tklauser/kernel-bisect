#!/bin/bash

set -e
set -o pipefail

make defconfig && \
./scripts/config --disable CONFIG_ACPI_AC && \
./scripts/config --disable CONFIG_ACPI_BATTERY && \
./scripts/config --disable CONFIG_ACPI_BUTTON && \
./scripts/config --disable CONFIG_ACPI_DOCK && \
./scripts/config --disable CONFIG_ACPI_FAN && \
./scripts/config --disable CONFIG_ACPI_THERMAL && \
./scripts/config --disable CONFIG_ACPI_TINY_POWER_BUTTON && \
./scripts/config --disable CONFIG_ACPI_VIDEO && \
./scripts/config --disable CONFIG_DRM && \
./scripts/config --disable CONFIG_EFI && \
./scripts/config --disable CONFIG_ETHTOOL_NETLINK && \
./scripts/config --disable CONFIG_EXT4_FS && \
./scripts/config --disable CONFIG_HID && \
./scripts/config --disable CONFIG_INPUT && \
./scripts/config --disable CONFIG_INPUT_MOUSE && \
./scripts/config --disable CONFIG_IPV6 && \
./scripts/config --disable CONFIG_MACINTOSH_DRIVERS && \
./scripts/config --disable CONFIG_NETFILTER && \
./scripts/config --disable CONFIG_NETWORK_FILESYSTEMS && \
./scripts/config --disable CONFIG_PCMCIA && \
./scripts/config --disable CONFIG_PM && \
./scripts/config --disable CONFIG_QUOTA && \
./scripts/config --disable CONFIG_RFKILL && \
./scripts/config --disable CONFIG_SECURITY_SELINUX && \
./scripts/config --disable CONFIG_SOUND && \
./scripts/config --disable CONFIG_USB && \
./scripts/config --disable CONFIG_WLAN && \
./scripts/config --disable CONFIG_WIRELESS && \
./scripts/config --enable CONFIG_XFRM && \
./scripts/config --enable CONFIG_XFRM_OFFLOAD && \
./scripts/config --enable CONFIG_XFRM_STATISTICS && \
./scripts/config --enable CONFIG_XFRM_ALGO && \
./scripts/config --enable CONFIG_XFRM_USER && \
./scripts/config --enable CONFIG_INET_ESP && \
./scripts/config --disable CONFIG_INET_ESP_OFFLOAD && \
./scripts/config --disable CONFIG_INET_ESPINTCP && \
./scripts/config --enable CONFIG_INET_IPCOMP && \
./scripts/config --enable CONFIG_INET_XFRM_TUNNEL && \
./scripts/config --enable CONFIG_INET_TUNNEL && \
./scripts/config --enable CONFIG_INET6_ESP && \
./scripts/config --enable CONFIG_INET6_IPCOMP && \
./scripts/config --enable CONFIG_INET6_XFRM_TUNNEL && \
./scripts/config --enable CONFIG_INET6_TUNNEL && \
make -j8
if [ $? -ne 0 ] ; then
	exit 125 # build failed, skip current revision.
fi

qemu-system-x86_64 -nographic -append console=ttyS0 -kernel arch/x86/boot/bzImage -initrd ./initrd.gz | grep "^===PIZZA"
