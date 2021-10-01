package pfs

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

// Pfs is a compression/zip format for everquest
type Pfs struct {
	ShortName                string
	Files                    []*PfsEntry
	directoryChunks          []*ChunkEntry
	directoryChunksTotalSize uint32
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
	Name            string
	Data            []byte
	CRC             uint32
	Offset          uint32
	chunks          []*ChunkEntry
	chunksTotalSize uint32
	filePointer     uint32
}

type ChunkEntry struct {
	deflatedSize int32
	inflatedSize int32
	data         []byte
}

func deflateChunks(data []byte) ([]*ChunkEntry, uint32, error) {

	chunks := []*ChunkEntry{}
	dataSize := len(data)
	chunksTotalSize := uint32(0)
	i := 0
	for i < dataSize {
		ce := &ChunkEntry{}
		blockSize := 8192
		if dataSize < 8192 {
			blockSize = dataSize
		}
		buf := &bytes.Buffer{}
		rawData := data[i : i+blockSize]
		if rawData == nil {
			fmt.Println("test")
		}

		zw := zlib.NewWriter(buf)
		_, err := zw.Write(data[i : i+blockSize])
		if err != nil {
			return nil, 0, fmt.Errorf("write: %w", err)
		}
		err = zw.Flush()
		if err != nil {
			return nil, 0, fmt.Errorf("flush: %w", err)
		}
		ce.data = buf.Bytes()
		ce.inflatedSize = int32(blockSize)
		ce.deflatedSize = int32(len(ce.data))
		chunks = append(chunks, ce)
		chunksTotalSize += uint32(ce.deflatedSize)
		i += blockSize
		dataSize = len(data) - i
	}
	return chunks, chunksTotalSize, nil
}
