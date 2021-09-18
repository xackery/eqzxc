package pfs

import (
	"fmt"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	path := "test/nexus.s3d"
	r, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}

	pfs, err := Load(r)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	fmt.Println(pfs.ShortName)
}
