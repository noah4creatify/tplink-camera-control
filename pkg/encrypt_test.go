package pkg

import (
	"fmt"
	"testing"

	"github.com/adrg/xdg"
)

func TestSecurityEncode(t *testing.T) {
	if securityEncode("test_password") != "77Vt4txvyPU0wwK" {
		t.Error("security encode test failed")
	}
}

func TestN(t *testing.T) {
	fmt.Println(xdg.ConfigHome)
}
