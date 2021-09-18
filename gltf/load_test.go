package gltf

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	path := "test/Cube.gltf"
	g, err := LoadFile(path)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	fmt.Println(g)
}
