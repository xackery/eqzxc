package bsp

import (
	"encoding/binary"
	"fmt"
	"io"
)

func (b *BSP) Encode(w io.WriteSeeker) error {
	err := binary.Write(w, binary.LittleEndian, []byte("IBSP"))
	if err != nil {
		return fmt.Errorf("write header: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, []byte{0x2E, 0x00})
	if err != nil {
		return fmt.Errorf("write header version: %w", err)
	}

	offset := int16(8 + 119)
	size := int16(len(b.EntityInfo))
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("entity offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("entity size: %w", err)
	}

	size = int16(len(b.Textures) * 72)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("texture offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("texture size: %w", err)
	}

	size = int16(len(b.Planes) * 16)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("plane offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("plane size: %w", err)
	}

	size = int16(len(b.Leaves) * 48)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("leaf offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("leaf size: %w", err)
	}

	size = int16(len(b.LeafFaces) * 4)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("leafface offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("leafface size: %w", err)
	}

	size = int16(len(b.LeafBrushes) * 4)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("leafbrush offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("leafbrush size: %w", err)
	}

	return nil
}
