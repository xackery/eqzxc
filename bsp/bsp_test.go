package bsp

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	path := "test/box.bsp"
	r, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer r.Close()

	bsp, err := Load(r)
	if err != nil {
		t.Fatalf("load bsp: %v", err)
	}
	if bsp == nil {
		t.Fatalf("nil bsp")
	}
}
