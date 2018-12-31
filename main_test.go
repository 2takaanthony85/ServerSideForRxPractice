package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	if "hello world!!" != getHello() {
		t.Fatal("failed test")
	}
}
