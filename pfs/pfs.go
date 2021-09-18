package pfs

// Pfs is a compression/zip format for everquest
type Pfs struct {
	ShortName string
	Files     []*PfsEntry
}

type PfsEntry struct {
	Name   string
	Data   []byte
	CRC    uint32
	Offset uint32
}
