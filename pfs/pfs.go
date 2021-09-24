package pfs

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

// Pfs is a compression/zip format for everquest
type Pfs struct {
	ShortName       string
	Files           []*PfsEntry
	filenamePointer uint32
}

type ByOffset []*PfsEntry

func (s ByOffset) Len() int {
	return len(s)
}

func (s ByOffset) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByOffset) Less(i, j int) bool {
	return s[i].Offset < s[j].Offset
}

type ByCRC []*PfsEntry

func (s ByCRC) Len() int {
	return len(s)
}

func (s ByCRC) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByCRC) Less(i, j int) bool {
	return s[i].CRC < s[j].CRC
}

type PfsEntry struct {
	Name         string
	Data         []byte
	CRC          uint32
	Offset       uint32
	deflatedData []*PfsDeflatedEntry
	filePointer  uint32
}

type PfsDeflatedEntry struct {
	deflatedSize int32
	inflatedSize int32
	data         []byte
}

func (e *PfsEntry) compress() error {
	buf := &bytes.Buffer{}

	dataSize := len(e.Data)
	i := 0
	for i < dataSize {
		ce := &PfsDeflatedEntry{}
		blockSize := 8193
		if dataSize-i < 8192 {
			blockSize = dataSize - i
		}
		_, err := zlib.NewWriter(buf).Write(e.Data[i : i+blockSize])
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
		ce.data = buf.Bytes()
		ce.inflatedSize = int32(blockSize)
		ce.deflatedSize = int32(len(ce.data))
		e.deflatedData = append(e.deflatedData, ce)
		i += blockSize
	}
	return nil
}
