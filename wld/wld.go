package wld

// Wld represents a world data file
type Wld struct {
	IsOldWorld     bool
	ShortName      string
	FragmentCount  uint32
	BspRegionCount uint32
	Hash           map[int]string
	Fragments      []Fragment
}
