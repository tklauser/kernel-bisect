# Finding a kernel regression using `git bisect run`

See https://github.com/cilium/cilium/issues/18500 for context. The test failure
suspiciously looked like a kernel issue (i.e. a netlink command that previously
succeeded started returning `EINVAL` after updating to latest `net-next`
kernel).

```
# assuming linux source tree in $HOME/src/linux
K=$HOME/src/linux
make
cp bisect.sh initrd.gz $K
cd $K
git bisect start
git bisect good v5.16
git bisect bad 0c947b893d69231a9add855939da7c66237ab44f # master as of 2022-01-17
git bisect run sh -c ./bisect.sh
```

This should eventually lead to the following commit being identified as
introducing the regression:

```
68ac0f3810e76a853b5f7b90601a05c3048b8b54 is the first bad commit
commit 68ac0f3810e76a853b5f7b90601a05c3048b8b54
Author: Antony Antony <antony.antony@secunet.com>
Date:   Sun Dec 12 11:35:00 2021 +0100

    xfrm: state and policy should fail if XFRMA_IF_ID 0

    xfrm ineterface does not allow xfrm if_id = 0
    fail to create or update xfrm state and policy.

    With this commit:
     ip xfrm policy add src 192.0.2.1 dst 192.0.2.2 dir out if_id 0
     RTNETLINK answers: Invalid argument

     ip xfrm state add src 192.0.2.1 dst 192.0.2.2 proto esp spi 1 \
                reqid 1 mode tunnel aead 'rfc4106(gcm(aes))' \
                0x1111111111111111111111111111111111111111 96 if_id 0
     RTNETLINK answers: Invalid argument

    v1->v2 change:
     - add Fixes: tag

    Fixes: 9f8550e4bd9d ("xfrm: fix disable_xfrm sysctl when used on xfrm interfaces")
    Signed-off-by: Antony Antony <antony.antony@secunet.com>
    Signed-off-by: Steffen Klassert <steffen.klassert@secunet.com>

 net/xfrm/xfrm_user.c | 21 ++++++++++++++++++---
 1 file changed, 18 insertions(+), 3 deletions(-)
```

The fix needs to be made in the netlink library to only set the `XFRMA_IF_ID`
netlink attribute in case `if_id` is `!= 0`, see
https://github.com/vishvananda/netlink/pull/727

## Reference

The use of `git bisect run` with a simple initrd was based on
https://ldpreload.com/blog/git-bisect-run
