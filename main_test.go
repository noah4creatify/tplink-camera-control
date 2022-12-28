package main

import "testing"

func TestSecurityEncode(t *testing.T) {
	if securityEncode("test_password") != "77Vt4txvyPU0wwK" {
		t.Error("security encode test failed")
	}
}
