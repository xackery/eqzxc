package pfs

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"sort"
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
	var directoryIndex uint32
	var magicNumber uint32
	var versionNumber uint32

	var value uint32
	err := binary.Read(r, binary.LittleEndian, &directoryIndex)
	if err != nil {
		return fmt.Errorf("read directory index: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &magicNumber)
	if err != nil {
		return fmt.Errorf("read magic number: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &versionNumber)
	if err != nil {
		return fmt.Errorf("read version number: %w", err)
	}

	_, err = r.Seek(int64(directoryIndex), io.SeekStart)
	if err != nil {
		return fmt.Errorf("seek directory index: %w", err)
	}
	var fileCount uint32
	err = binary.Read(r, binary.LittleEndian, &fileCount)
	if err != nil {
		return fmt.Errorf("read file count: %w", err)
	}
	if fileCount == 0 {
		return fmt.Errorf("empty file")
	}

	filenames := []string{}

	for i := 0; i < int(fileCount); i++ {
		entry := &PfsEntry{}

		err = binary.Read(r, binary.LittleEndian, &entry.CRC)
		if err != nil {
			return fmt.Errorf("read crc %d/%d: %w", i, fileCount, err)
		}

		err = binary.Read(r, binary.LittleEndian, &entry.Offset)
		if err != nil {
			return fmt.Errorf("read offset %d/%d: %w", i, fileCount, err)
		}
		debugInfo := fmt.Sprintf("%d/%d 0x%x", i, fileCount, entry.Offset)
		// size is the uncompressed size of the file
		var size uint32
		err = binary.Read(r, binary.LittleEndian, &size)
		if err != nil {
			return fmt.Errorf("read size %s: %w", debugInfo, err)
		}

		buf := bytes.NewBuffer(nil)

		cachedOffset, err := r.Seek(0, io.SeekCurrent)
		if err != nil {
			return fmt.Errorf("seek cached offset %s: %w", debugInfo, err)
		}

		_, err = r.Seek(int64(entry.Offset), io.SeekStart)
		if err != nil {
			return fmt.Errorf("seek offset %s: %w", debugInfo, err)
		}

		for uint32(buf.Len()) != size {
			var deflatedLength uint32
			var inflatedLength uint32
			err = binary.Read(r, binary.LittleEndian, &deflatedLength)
			if err != nil {
				return fmt.Errorf("read deflated length %s: %w", debugInfo, err)
			}

			err = binary.Read(r, binary.LittleEndian, &inflatedLength)
			if err != nil {
				return fmt.Errorf("read inflated length %s: %w", debugInfo, err)
			}

			//if inflatedLength < deflatedLength {
			//	return fmt.Errorf("inflated < deflated, offset misaligned? %d/%d", i, fileCount)
			//}

			compressedData := make([]byte, deflatedLength)
			err = binary.Read(r, binary.LittleEndian, compressedData)
			if err != nil {
				return fmt.Errorf("read compressed data: %s: %w", debugInfo, err)
			}

			fr, err := zlib.NewReaderDict(bytes.NewReader(compressedData), nil)
			if err != nil {
				return fmt.Errorf("zlib new reader dict %s: %w", debugInfo, err)
			}

			inflatedWritten, err := io.Copy(buf, fr)
			if err != nil {
				return fmt.Errorf("copy: %s: %w", debugInfo, err)
			}

			if inflatedWritten != int64(inflatedLength) {
				return fmt.Errorf("inflate mismatch %s: %w", debugInfo, err)
			}
		}
		if buf.Len() < int(size) {
			return fmt.Errorf("size mismatch %s: %w", debugInfo, err)
		}
		entry.Data = buf.Bytes()

		if entry.CRC == 0x61580AC9 {
			fr := bytes.NewReader(buf.Bytes())
			var filenameCount uint32
			err = binary.Read(fr, binary.LittleEndian, &filenameCount)
			if err != nil {
				return fmt.Errorf("filename count %s: %w", debugInfo, err)
			}

			for j := uint32(0); j < filenameCount; j++ {
				err = binary.Read(fr, binary.LittleEndian, &value)
				if err != nil {
					return fmt.Errorf("filename length %s: %w", debugInfo, err)
				}
				filename, err := parseFixedString(fr, value)
				if err != nil {
					return fmt.Errorf("filename %s: %w", debugInfo, err)
				}
				filenames = append(filenames, filename)
			}

			_, err = r.Seek(cachedOffset, io.SeekStart)
			if err != nil {
				return fmt.Errorf("seek cached offset %s: %w", debugInfo, err)
			}
			continue
		}
		pfs.Files = append(pfs.Files, entry)
		_, err = r.Seek(cachedOffset, io.SeekStart)
		if err != nil {
			return fmt.Errorf("seek cached offset %s: %w", debugInfo, err)
		}
	}

	sort.Sort(ByOffset(pfs.Files))
	for i, entry := range pfs.Files {
		if len(filenames) < i {
			return fmt.Errorf("entry %d has no name", i)
		}
		entry.Name = filenames[i]
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
