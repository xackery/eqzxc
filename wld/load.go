package wld

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// Load will load a world file
func Load(r io.ReadSeeker) (*Wld, error) {
	wld := &Wld{}
	err := parse(r, wld)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	return wld, nil
}

func parse(r io.ReadSeeker, wld *Wld) error {
	if wld == nil {
		return fmt.Errorf("wld nil")
	}
	var value uint32

	err := binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read wld header: %w", err)
	}
	if value != 0x54503D02 {
		return fmt.Errorf("unknown wld header: wanted 0x%x, got 0x%x", 0x54503D02, value)
	}

	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read identifier: %w", err)
	}
	switch value {
	case 0x00015500:
		wld.IsOldWorld = true
	case 0x1000C800:
		wld.IsOldWorld = false
	default:
		return fmt.Errorf("unknown wld identifier %d", value)
	}

	err = binary.Read(r, binary.LittleEndian, &wld.FragmentCount)
	if err != nil {
		return fmt.Errorf("read fragment count: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &wld.BspRegionCount)
	if err != nil {
		return fmt.Errorf("read bsp region count: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read after bsp region offset: %w", err)
	}
	//if value != 0x000680D4 {
	//	return fmt.Errorf("after bsp region offset wanted 0x%x, got 0x%x", 0x000680D4, value)
	//}//

	var hashSize uint32
	err = binary.Read(r, binary.LittleEndian, &hashSize)
	if err != nil {
		return fmt.Errorf("read hash size: %w", err)
	}
	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read after hash size offset: %w", err)
	}

	hashRaw, err := parseFixedString(r, hashSize)
	if err != nil {
		return fmt.Errorf("read hash: %w", err)
	}

	hashSplit := strings.Split(hashRaw, "\x00")
	wld.Hash = make(map[int]string)
	offset := 0
	for _, h := range hashSplit {
		wld.Hash[offset] = h
		offset += len(h) + 1
	}
	fmt.Println(wld.Hash)

	for i := 0; i < int(wld.FragmentCount); i++ {
		var fragSize uint32
		var fragIndex int32
		err = binary.Read(r, binary.LittleEndian, &fragSize)
		if err != nil {
			return fmt.Errorf("read fragment size %d/%d: %w", i, wld.FragmentCount, err)
		}
		err = binary.Read(r, binary.LittleEndian, &fragIndex)
		if err != nil {
			return fmt.Errorf("read fragment index %d/%d: %w", i, wld.FragmentCount, err)
		}
		switch fragIndex {
		case 0x1B:
			l, err := loadLightSource(r)
			if err != nil {
				return fmt.Errorf("parse light source %d/%d: %w", i, wld.FragmentCount, err)
			}
			wld.Fragments = append(wld.Fragments, l)
		case 0x30:
			m, err := loadMaterial(r)
			if err != nil {
				return fmt.Errorf("parse material %d/%d: %w", i, wld.FragmentCount, err)
			}
			wld.Fragments = append(wld.Fragments, m)
		case 0x32:
			v, err := loadVertexColor(r)
			if err != nil {
				return fmt.Errorf("parse vertex color %d/%d: %w", i, wld.FragmentCount, err)
			}
			wld.Fragments = append(wld.Fragments, v)
		}
	}
	return nil
}

func parseFixedString(r io.ReadSeeker, size uint32) (string, error) {
	in := make([]byte, size)
	_, err := r.Read(in)
	if err != nil {
		return "", fmt.Errorf("read: %w", err)
	}
	final := ""
	for _, char := range in {
		if char == 0x0 {
			continue
		}
		final += string(char)
	}
	return final, nil
}
