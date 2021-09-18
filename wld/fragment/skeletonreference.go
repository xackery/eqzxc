package fragment

import (
	"encoding/binary"
	"fmt"
	"io"
)

// SkeletonReference information
type SkeletonReference struct {
	hashIndex uint32
	Reference uint32
	FrameMs   uint32
}

func LoadSkeletonReference(r io.ReadSeeker) (*SkeletonReference, error) {
	v := &SkeletonReference{}
	err := parseSkeletonReference(r, v)
	if err != nil {
		return nil, fmt.Errorf("parse skeleton reference: %w", err)
	}
	return v, nil
}

func parseSkeletonReference(r io.ReadSeeker, v *SkeletonReference) error {
	if v == nil {
		return fmt.Errorf("skeleton reference is nil")
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
		return fmt.Errorf("read flag: %w", err)
	}

	//TODO: review
	// Either 4 or 5 - maybe something to look into
	// Bits are set 0, or 2. 0 has the extra field for delay.
	// 2 doesn't have any additional fields.
	if value&1 == 1 {
		err = binary.Read(r, binary.LittleEndian, &v.FrameMs)
		if err != nil {
			return fmt.Errorf("read frame ms: %w", err)
		}
	}

	return nil
}

func (v *SkeletonReference) FragmentType() string {
	return "Skeleton Reference"
}
