package wld

import (
	"fmt"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	path := "test/nexus.wld"
	r, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer r.Close()

	wld, err := Decode(r)
	if err != nil {
		t.Fatalf("load wld: %v", err)
	}
	fmt.Println(wld.ShortName)
}
