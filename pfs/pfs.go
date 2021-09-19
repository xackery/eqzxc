package pfs

// Pfs is a compression/zip format for everquest
type Pfs struct {
	ShortName string
	Files     []*PfsEntry
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

type PfsEntry struct {
	Name   string
	Data   []byte
	CRC    uint32
	Offset uint32
}
