package wld

import "github.com/xackery/eqzxc/wld/fragment"

// Wld represents a world data file
type Wld struct {
	IsOldWorld     bool
	ShortName      string
	FragmentCount  uint32
	BspRegionCount uint32
	Hash           map[int]string
	Fragments      []fragment.Fragment
}
