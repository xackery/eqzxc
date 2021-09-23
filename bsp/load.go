package bsp

import (
	"encoding/binary"
	"fmt"
	"io"
)

func Load(r io.ReadSeeker) (*BSP, error) {
	b := &BSP{}
	err := parse(r, b)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	return b, nil
}

func parse(r io.ReadSeeker, b *BSP) error {
	//var value int32

	bh := &bspHeader{}
	err := binary.Read(r, binary.LittleEndian, bh)
	if err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	if string(bh.Header[:]) != "IBSP" {
		return fmt.Errorf("header got %s, want IBSP", bh.Header[:])
	}

	if bh.Version != 0x2E {
		return fmt.Errorf("header version got 0x%x, want 0x2E", bh.Version)
	}

	dirEntries := make([]entry, 19)
	err = binary.Read(r, binary.LittleEndian, &dirEntries)
	if err != nil {
		return fmt.Errorf("read dirEntries: %w", err)
	}

	b.EntityInfo, err = parseFixedString(r, uint32(dirEntries[dirEntryEntities].Size))
	if err != nil {
		return fmt.Errorf("parse entities: %w", err)
	}

	count := dirEntries[dirEntryTextures].Size / 76
	for i := int32(0); i < count; i++ {
		v := &Texture{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read texture %d: %w", i, err)
		}
		b.Textures = append(b.Textures, v)
	}

	count = dirEntries[dirEntryPlanes].Size / 16
	for i := int32(0); i < count; i++ {
		v := &Plane{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read plane %d: %w", i, err)
		}
		b.Planes = append(b.Planes, v)
	}

	count = dirEntries[dirEntryNodes].Size / 36
	for i := int32(0); i < count; i++ {
		v := &Node{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read node %d: %w", i, err)
		}
		b.Nodes = append(b.Nodes, v)
	}

	count = dirEntries[dirEntryLeafs].Size / 8
	for i := int32(0); i < count; i++ {
		v := &Leaf{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read leaf %d: %w", i, err)
		}
		b.Leaves = append(b.Leaves, v)
	}

	count = dirEntries[dirEntryLeaffaces].Size / 4
	for i := int32(0); i < count; i++ {
		v := &LeafFace{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read leaf face %d: %w", i, err)
		}
		b.LeafFaces = append(b.LeafFaces, v)
	}

	count = dirEntries[dirEntryLeafbrushes].Size / 4
	for i := int32(0); i < count; i++ {
		v := &LeafBrush{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read leaf brush %d: %w", i, err)
		}
		b.LeafBrushes = append(b.LeafBrushes, v)
	}

	count = dirEntries[dirEntryModels].Size / 40
	for i := int32(0); i < count; i++ {
		v := &Model{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read model %d: %w", i, err)
		}
		b.Models = append(b.Models, v)
	}

	count = dirEntries[dirEntryBrushes].Size / 12
	for i := int32(0); i < count; i++ {
		v := &Brush{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read brush %d: %w", i, err)
		}
		b.Brushes = append(b.Brushes, v)
	}

	count = dirEntries[dirEntryBrushsides].Size / 44
	for i := int32(0); i < count; i++ {
		v := &BrushSide{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read brush side %d: %w", i, err)
		}
		b.BrushSides = append(b.BrushSides, v)
	}

	count = dirEntries[dirEntryVertexes].Size / 44
	for i := int32(0); i < count; i++ {
		v := &Vertex{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read vertex %d: %w", i, err)
		}
		b.Vertexes = append(b.Vertexes, v)
	}

	count = dirEntries[dirEntryMeshverts].Size / 4
	for i := int32(0); i < count; i++ {
		v := &MeshVertexOffset{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read mesh vertex offset %d: %w", i, err)
		}
		b.MeshVertexOffsets = append(b.MeshVertexOffsets, v)
	}

	count = dirEntries[dirEntryEffects].Size / 72
	for i := int32(0); i < count; i++ {
		v := &Effect{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read effect %d: %w", i, err)
		}
		b.Effects = append(b.Effects, v)
	}

	count = dirEntries[dirEntryFaces].Size / 108
	for i := int32(0); i < count; i++ {
		v := &Face{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read faces %d: %w", i, err)
		}
		b.Faces = append(b.Faces, v)
	}

	count = dirEntries[dirEntryLightmaps].Size / 49152
	for i := int32(0); i < count; i++ {
		v := &Lightmap{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read lightmap %d: %w", i, err)
		}
		b.Lightmaps = append(b.Lightmaps, v)
	}

	count = dirEntries[dirEntryLightvols].Size / 8
	for i := int32(0); i < count; i++ {
		v := &LightVolume{}
		err = binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			return fmt.Errorf("read lightvolume %d: %w", i, err)
		}
		b.LightVolumes = append(b.LightVolumes, v)
	}

	count = dirEntries[dirEntryVisdata].Size / 8
	for i := int32(0); i < count; i++ {
		v := &VisData{}

		err = binary.Read(r, binary.LittleEndian, &v.VectorCount)
		if err != nil {
			return fmt.Errorf("read visdata count %d: %w", i, err)
		}

		err = binary.Read(r, binary.LittleEndian, &v.VectorSize)
		if err != nil {
			return fmt.Errorf("read visdata size %d: %w", i, err)
		}

		v.Vectors = make([]uint8, v.VectorCount*v.VectorSize)
		for j := range v.Vectors {
			err = binary.Read(r, binary.LittleEndian, &v.Vectors[j])
			if err != nil {
				return fmt.Errorf("read visdata vector %d %d: %w", i, j, err)
			}
		}

		b.VisInfo = append(b.VisInfo, v)
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
