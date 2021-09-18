package wld

import (
	"fmt"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	path := "test/nexus.wld"
	r, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer r.Close()

	wld, err := Load(r)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	fmt.Println(wld.ShortName)
}
