package q3map

import "github.com/g3n/engine/math32"

// https://www.gamers.org/dEngine/quake/QDP/qmapspec.html

type Q3Map struct {
	Entities []*Entity
}

type Entity struct {
	ClassName string
	Origin    math32.Vector3
	Light     string
	Brushes   []*Brush
}

type Brush struct {
	Defs []*BrushDef
}

type BrushDef struct {
	Points  [5]math32.Vector3
	Texture string
	Unk1    string
	Unk2    string
	Unk3    string
}
