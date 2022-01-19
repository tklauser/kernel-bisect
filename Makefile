all:
	CGO_ENABLED=0 go build ./cmd/init
	strip init
	echo init | cpio -H newc -o | gzip > initrd.gz

clean:
	-rm -f init initrd.gz
