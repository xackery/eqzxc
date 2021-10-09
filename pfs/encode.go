package pfs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"sort"

	"github.com/xackery/eqzxc/crc"
)

func (pfs *Pfs) Encode(w io.WriteSeeker) error {
	var err error
	directoryIndex := uint32(12)
	magicNumber := int32(0x20534650)
	versionNumber := int32(0x00020000)

	for _, entry := range pfs.Files {
		entry.chunks, entry.chunksTotalSize, err = deflateChunks(entry.Data)
		if err != nil {
			return fmt.Errorf("deflatechunks %s: %w", entry.Name, err)
		}
		directoryIndex += 4 + 4

		directoryIndex += uint32(len(entry.chunks))

		entry.CRC = crc.FilenameCRC32(entry.Name)
	}
	sort.Sort(ByCRC(pfs.Files))

	directoryBuf := &bytes.Buffer{}
	err = binary.Write(directoryBuf, binary.LittleEndian, uint32(len(pfs.Files)))
	if err != nil {
		return fmt.Errorf("write len pfs.Files: %w", err)
	}
	for _, entry := range pfs.Files {
		err = binary.Write(directoryBuf, binary.LittleEndian, uint32(len(entry.Name)+1))
		if err != nil {
			return fmt.Errorf("write len pfs.File %s: %w", entry.Name, err)
		}
		err = binary.Write(directoryBuf, binary.LittleEndian, []byte(entry.Name))
		if err != nil {
			return fmt.Errorf("write entry name %s: %w", entry.Name, err)
		}
		err = binary.Write(directoryBuf, binary.LittleEndian, []byte{0x00})
		if err != nil {
			return fmt.Errorf("write zero %s: %w", entry.Name, err)
		}
	}
	pfs.directoryChunks, pfs.directoryChunksTotalSize, err = deflateChunks(directoryBuf.Bytes())
	if err != nil {
		return fmt.Errorf("deflateChunks directoryBuf: %w", err)
	}

	for _, entry := range pfs.Files {
		chunkSize := len(entry.chunks)
		if chunkSize < 1 {
			chunkSize = 1
		}
		directoryIndex += uint32(4+4)*uint32(chunkSize) + entry.chunksTotalSize
	}

	chunkSize := len(pfs.directoryChunks)
	if chunkSize < 1 {
		chunkSize = 1
	}

	directoryIndex += uint32(4+4)*uint32(chunkSize) + uint32(pfs.directoryChunksTotalSize)

	////
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

		for i, chunk := range entry.chunks {
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

	filePtr, err := w.Seek(0, io.SeekCurrent)
	if err != nil {
		return fmt.Errorf("seek directory chunk index: %w", err)
	}
	for i, chunk := range pfs.directoryChunks {
		err = binary.Write(w, binary.LittleEndian, chunk.deflatedSize)
		if err != nil {
			return fmt.Errorf("write directory chunk deflated size %d: %w", i, err)
		}
		err = binary.Write(w, binary.LittleEndian, chunk.inflatedSize)
		if err != nil {
			return fmt.Errorf("write directory chunk inflated size %d: %w", i, err)
		}
		err = binary.Write(w, binary.LittleEndian, chunk.data)
		if err != nil {
			return fmt.Errorf("write directory chunk deflated data %d: %w", i, err)
		}
	}

	err = binary.Write(w, binary.LittleEndian, uint32(len(pfs.Files)+1))
	if err != nil {
		return fmt.Errorf("write file list: %w", err)
	}

	for _, entry := range pfs.Files {
		err = binary.Write(w, binary.LittleEndian, entry.CRC)
		if err != nil {
			return fmt.Errorf("write file %s crc: %w", entry.Name, err)
		}
		err = binary.Write(w, binary.LittleEndian, entry.filePointer)
		if err != nil {
			return fmt.Errorf("write file %s filePointer: %w", entry.Name, err)
		}
		err = binary.Write(w, binary.LittleEndian, uint32(len(entry.Data)))
		if err != nil {
			return fmt.Errorf("write file %s uncompressed size: %w", entry.Name, err)
		}
	}

	err = binary.Write(w, binary.LittleEndian, uint32(0xFFFFFFFF))
	if err != nil {
		return fmt.Errorf("write filename directory footer: %w", err)
	}

	err = binary.Write(w, binary.LittleEndian, filePtr)
	if err != nil {
		return fmt.Errorf("write directory filePtr: %w", err)
	}

	directoryTotalSize := uint32(0)
	for _, entry := range pfs.Files {
		directoryTotalSize += uint32(len(entry.Data))
	}
	err = binary.Write(w, binary.LittleEndian, directoryTotalSize)
	if err != nil {
		return fmt.Errorf("write directoryTotalSize: %w", err)
	}

	err = binary.Write(w, binary.LittleEndian, []byte("STEVE"))
	if err != nil {
		return fmt.Errorf("write STEVE footer: %w", err)
	}

	err = binary.Write(w, binary.LittleEndian, uint32(69))
	if err != nil {
		return fmt.Errorf("write datestamp: %w", err)
	}
	return nil
}
