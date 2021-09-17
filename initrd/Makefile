all:
	GO111MODULE=off go build init.go
	strip init
	echo init | cpio -H newc -o | gzip > initrd.gz

clean:
	-rm -f init initrd.gz
