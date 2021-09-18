package fragment

import (
	"encoding/binary"
	"fmt"
	"io"
)

// ParticleSpriteReference information
type ParticleSpriteReference struct {
	hashIndex uint32
	Reference uint32
}

func LoadParticleSpriteReference(r io.ReadSeeker) (*ParticleSpriteReference, error) {
	v := &ParticleSpriteReference{}
	err := parseParticleSpriteReference(r, v)
	if err != nil {
		return nil, fmt.Errorf("parse particle sprite reference: %w", err)
	}
	return v, nil
}

func parseParticleSpriteReference(r io.ReadSeeker, v *ParticleSpriteReference) error {
	if v == nil {
		return fmt.Errorf("particle sprite reference is nil")
	}
	var value uint32
	err := binary.Read(r, binary.LittleEndian, &v.hashIndex)
	if err != nil {
		return fmt.Errorf("read hash index: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &v.Reference)
	if err != nil {
		return fmt.Errorf("read reference: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read post reference: %w", err)
	}
	if value != 8 {
		return fmt.Errorf("post reference got %d, wanted %d", value, 8)
	}

	return nil
}

func (v *ParticleSpriteReference) FragmentType() string {
	return "Particle Sprite Reference"
}
