package gltf

import (
	"fmt"

	"github.com/qmuntal/gltf"
)

func LoadFile(path string) (*GLTF, error) {
	var err error
	g := &GLTF{}
	g.Document, err = gltf.Open(path)
	if err != nil {
		return nil, fmt.Errorf("gltf open: %w", err)
	}

	return g, nil
}
