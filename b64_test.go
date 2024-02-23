package b64_test

import (
	"testing"

	"github.com/vxcute/base64/b64"
)

func TestBase64Encode(t *testing.T) {
	original := "AAAAAAAAhhhhhhhhhh912738712dskajdkxnchWWWHKHkashkdhakhkhsk"
	test := "QUFBQUFBQUFoaGhoaGhoaGhoOTEyNzM4NzEyZHNrYWpka3huY2hXV1dIS0hrYXNoa2RoYWtoa2hzaw=="
	encoded := b64.Base64Encode(original)

	if test != encoded {
		t.Fatal("Failed")
	}
}

func TestBase64Decode(t *testing.T) {
	encoded := "QUFBQUFBQUFoaGhoaGhoaGhoOTEyNzM4NzEyZHNrYWpka3huY2hXV1dIS0hrYXNoa2RoYWtoa2hzaw=="
	test := "AAAAAAAAhhhhhhhhhh912738712dskajdkxnchWWWHKHkashkdhakhkhsk"
	decoded := b64.Base64Decode(encoded)

	if test != decoded {
		t.Fatal("Failed")
	}
}
