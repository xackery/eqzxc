package bsp

import (
	"encoding/binary"
	"fmt"
	"io"
)

func (b *BSP) Encode(w io.WriteSeeker) error {

	err := binary.Write(w, binary.LittleEndian, b.header)
	if err != nil {
		return fmt.Errorf("write header: %w", err)
	}

	err = binary.Write(w, binary.LittleEndian, int32(8024))
	if err != nil {
		return fmt.Errorf("entity direntries: %w", err)
	}

	offset := int32(0)
	size := int32(len(b.EntityInfo))
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("entity offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("entity size: %w", err)
	}

	size = int32(len(b.Textures) * dirEntryTexturesSize)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("texture offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("texture size: %w", err)
	}

	size = int32(len(b.Planes) * dirEntryPlanesSize)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("plane offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("plane size: %w", err)
	}

	size = int32(len(b.Leaves) * dirEntryLeafsSize)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("leaf offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("leaf size: %w", err)
	}

	size = int32(len(b.LeafFaces) * dirEntryLeaffacesSize)
	offset += size
	err = binary.Write(w, binary.LittleEndian, offset)
	if err != nil {
		return fmt.Errorf("leafface offset: %w", err)
	}
	err = binary.Write(w, binary.LittleEndian, size)
	if err != nil {
		return fmt.Errorf("leafface size: %w", err)
	}

	size = int32(len(b.LeafBrushes) * dirEntryLeafbrushesSize)
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
