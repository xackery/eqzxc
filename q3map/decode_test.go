package q3map

import (
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	path := "test/clz.map"
	r, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer r.Close()

	m, err := Decode(r)
	if err != nil {
		t.Fatalf("decode q3map: %v", err)
	}
	if m == nil {
		t.Fatalf("nil map")
	}

}
