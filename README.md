# Finding a kernel regression using `git bisect run`

See https://github.com/cilium/cilium/pull/17394#issuecomment-920902042 and
https://github.com/cilium/cilium/pull/17394#issuecomment-921000409 for context.

```
# assuming linux source tree in $HOME/src/linux
make -C initrd
git bisect start
git bisect good v5.14
git bisect bad v5.15-rc1
git bisect run sh -c bisect.sh
```

This should eventually lead to the following commit being identified as
introducing the regression:

```
75bd50fa841db5434728d238b8b5659498ccf0ab is the first bad commit
commit 75bd50fa841db5434728d238b8b5659498ccf0ab
Author: Tian Tao <tiantao6@hisilicon.com>
Date:   Fri Aug 6 23:02:50 2021 +1200

    drivers/base/node.c: use bin_attribute to break the size limitation of cpumap ABI

    Reading /sys/devices/system/cpu/cpuX/nodeX/ returns cpumap and cpulist.
    However, the size of this file is limited to PAGE_SIZE because of the
    limitation for sysfs attribute.

    This patch moves to use bin_attribute to extend the ABI to be more
    than one page so that cpumap bitmask and list won't be potentially
    trimmed.

    Cc: Greg Kroah-Hartman <gregkh@linuxfoundation.org>
    Cc: "Rafael J. Wysocki" <rafael@kernel.org>
    Reviewed-by: Jonathan Cameron <Jonathan.Cameron@huawei.com>
    Signed-off-by: Tian Tao <tiantao6@hisilicon.com>
    Signed-off-by: Barry Song <song.bao.hua@hisilicon.com>
    Link: https://lore.kernel.org/r/20210806110251.560-5-song.bao.hua@hisilicon.com
    Signed-off-by: Greg Kroah-Hartman <gregkh@linuxfoundation.org>

 drivers/base/node.c | 63 ++++++++++++++++++++++++++++++++++-------------------
 1 file changed, 40 insertions(+), 23 deletions(-)
```

The proposed fix for the regression was [submitted upstream][1], [applied to
the driver-core tree][2] and [was merged][3] to Linus' tree for 5.15-rc4.

[1]: https://lore.kernel.org/lkml/20210916222705.13554-1-tklauser@distanz.ch/
[2]: https://git.kernel.org/pub/scm/linux/kernel/git/gregkh/driver-core.git/commit/?h=driver-core-linus&id=c86a2d9058c5a4a05d20ef89e699b7a6b2c89da6
[3]: https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=84928ce3bb4e20ec7ef0e990630a690855dd44cc

## Reference

The use of `git bisect run` with a simple initrd was based on
https://ldpreload.com/blog/git-bisect-run
