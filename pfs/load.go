package pfs

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Load will load a pfs file
func Load(r io.ReadSeeker) (*Pfs, error) {
	pfs := &Pfs{}
	err := parse(r, pfs)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	return pfs, nil
}

func parse(r io.ReadSeeker, pfs *Pfs) error {
	var value uint32
	err := binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read directory offset: %w", err)
	}
	_, err = r.Seek(int64(value), io.SeekCurrent)
	if err != nil {
		return fmt.Errorf("seek directory offset: %w", err)
	}
	var fileCount uint32
	err = binary.Read(r, binary.LittleEndian, &fileCount)
	if err != nil {
		return fmt.Errorf("read file count: %w", err)
	}

	//files := []string{}

	for i := 0; i < int(fileCount); i++ {
		var crc uint32
		err = binary.Read(r, binary.LittleEndian, &crc)
		if err != nil {
			return fmt.Errorf("read crc %d/%d: %w", i, fileCount, err)
		}
		var offset uint32
		err = binary.Read(r, binary.LittleEndian, &offset)
		if err != nil {
			return fmt.Errorf("read offset %d/%d: %w", i, fileCount, err)
		}
		var size uint32
		err = binary.Read(r, binary.LittleEndian, &size)
		if err != nil {
			return fmt.Errorf("read size %d/%d: %w", i, fileCount, err)
		}

		//cachedOffset, err := r.Seek(0, io.SeekCurrent)
		//if err != nil {
		//	return fmt.Errorf("seek cached offset %d/%d: %w", i, fileCount, err)
		//}

	}
	return nil
}
