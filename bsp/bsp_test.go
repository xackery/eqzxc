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

	bsp, err := Decode(r)
	if err != nil {
		t.Fatalf("decode bsp: %v", err)
	}
	if bsp == nil {
		t.Fatalf("nil bsp")
	}

	f, err := os.Create("test/out.bsp")
	if err != nil {
		t.Fatalf("create out.bsp: %s", err.Error())
	}
	err = bsp.Encode(f)
	if err != nil {
		t.Fatalf("encode: %s", err.Error())
	}
}
