package pfs

import (
	"encoding/binary"
	"fmt"
	"io"
	"sort"

	"github.com/xackery/eqzxc/crc"
)

func (pfs *Pfs) Save(w io.WriteSeeker) error {
	var err error
	directoryIndex := uint32(12)
	magicNumber := int32(0x20534650)
	versionNumber := int32(0x00020000)

	for _, entry := range pfs.Files {
		err = entry.compress()
		if err != nil {
			return fmt.Errorf("compress %s: %w", entry.Name, err)
		}
		directoryIndex += 4 + 4

		directoryIndex += uint32(len(entry.deflatedData))

		entry.CRC = crc.FilenameCRC32(entry.Name)
	}
	sort.Sort(ByCRC(pfs.Files))

	err = binary.Write(w, binary.LittleEndian, directoryIndex)
	if err != nil {
		return fmt.Errorf("write directory index: %w", err)
	}

	err = binary.Write(w, binary.LittleEndian, &magicNumber)
	if err != nil {
		return fmt.Errorf("write magic number: %w", err)
	}

	err = binary.Write(w, binary.LittleEndian, &versionNumber)
	if err != nil {
		return fmt.Errorf("write version number: %w", err)
	}

	var ptr int64
	for _, entry := range pfs.Files {
		ptr, err = w.Seek(0, io.SeekCurrent)
		if err != nil {
			return fmt.Errorf("seek file pointer %s: %w", entry.Name, err)
		}
		entry.filePointer = uint32(ptr)

		for i, chunk := range entry.deflatedData {
			err = binary.Write(w, binary.LittleEndian, chunk.deflatedSize)
			if err != nil {
				return fmt.Errorf("write %s deflated size %d: %w", entry.Name, i, err)
			}
			err = binary.Write(w, binary.LittleEndian, chunk.inflatedSize)
			if err != nil {
				return fmt.Errorf("write %s inflated size %d: %w", entry.Name, i, err)
			}
			err = binary.Write(w, binary.LittleEndian, chunk.data)
			if err != nil {
				return fmt.Errorf("write %s deflated data %d: %w", entry.Name, i, err)
			}
		}
	}
	ptr, err = w.Seek(0, io.SeekCurrent)
	if err != nil {
		return fmt.Errorf("seek filename pointer: %w", err)
	}
	pfs.filenamePointer = uint32(ptr)

	return nil
}
