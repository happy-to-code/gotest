package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var s = `Ly/lkIjnuqbnpLrkvosKY29udHJhY3QgVGVzdEV4YW1wbGUyIHsKCS8v5Yid5aeL5YyW5LiA5Liq6LSm5oi3CglwdWJsaWMgc3RyaW5nIHNldChzdHJpbmcgYWNjb3VudCxpbnQgYW1vdW50KXsKCQlwcmludChhY2NvdW50KQoJCXByaW50KGFtb3VudCkKCQlkYl9zZXQoYWNjb3VudCxhbW91bnQpCiAgICAgICAgcmV0dXJuICJzdWNjZXNzIgoJfQoJCglwdWJsaWMgc3RyaW5nIGluaXQoKXsKCQlyZXR1cm4gInN1Y2Nlc3MiCgl9CgkKCS8v6L2s6LSm5pON5L2cCglwdWJsaWMgc3RyaW5nIHRyYW5zZmVyKHN0cmluZyBmcm9tLCBzdHJpbmcgdG8sIGludCBhbW91bnQpIHsKCQlpbnQgYmFsQSA9IGRiX2dldDxpbnQ+KGZyb20pCgkJaW50IGJhbEIgPSBkYl9nZXQ8aW50Pih0bykKCQliYWxBID0gYmFsQS1hbW91bnQKCQlpZiAoYmFsQT4wKXsKCQkJYmFsQiA9IGJhbEIrYW1vdW50CgkJCWRiX3NldChmcm9tLCBiYWxBKQoJCQlkYl9zZXQodG8sIGJhbEIpCgkJfWVsc2V7CiAgICAgICAgICAgIHJldHVybiAiZmFpbGVkIgogICAgICAgIH0KCQlyZXR1cm4gInN1Y2Nlc3MiCgl9CgoJLy/mn6Xor6LotKbmiLfkvZnpop0KCXB1YmxpYyBpbnQgZ2V0QmxhbmNlKHN0cmluZyBhY2NvdW50KXsKCQlwcmludChhY2NvdW50KQoJCWludCBhID0gZGJfZ2V0PGludD4oYWNjb3VudCkKCQlyZXR1cm4gYQoJfQoJCn0=`

	bytes, _ := base64.StdEncoding.DecodeString(s)
	fmt.Printf("bytes:%+v\n", bytes)

	fmt.Printf("str::%s\n", bytes)
}
