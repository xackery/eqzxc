package gltf

import (
	"fmt"
	"io"

	"github.com/qmuntal/gltf"
)

func Save(w io.WriteSeeker, g *GLTF) error {
	enc := gltf.NewEncoder(w)
	enc.AsBinary = false
	err := enc.Encode(g.Document)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	return nil
}
