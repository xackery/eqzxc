package fragment

import (
	"encoding/binary"
	"fmt"
	"io"
)

// BspRegion information
type BspRegion struct {
	hashIndex   uint32
	HasPolygons bool
	Reference   uint32
	RegionType  uint32
}

func LoadBspRegion(r io.ReadSeeker) (*BspRegion, error) {
	v := &BspRegion{}
	err := parseBspRegion(r, v)
	if err != nil {
		return nil, fmt.Errorf("parse bsp region: %w", err)
	}
	return v, nil
}

func parseBspRegion(r io.ReadSeeker, v *BspRegion) error {
	if v == nil {
		return fmt.Errorf("bsp region is nil")
	}
	var value uint32
	err := binary.Read(r, binary.LittleEndian, &v.hashIndex)
	if err != nil {
		return fmt.Errorf("read hash index: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read flags: %w", err)
	}
	// Flags
	// 0x181 - Regions with polygons
	// 0x81 - Regions without
	// Bit 5 - PVS is WORDS
	// Bit 7 - PVS is bytes
	if value == 0x181 {
		v.HasPolygons = true
	}

	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read unknown1: %w", err)
	}

	var data1Size uint32
	err = binary.Read(r, binary.LittleEndian, &data1Size)
	if err != nil {
		return fmt.Errorf("read data1size: %w", err)
	}

	var data2Size uint32
	err = binary.Read(r, binary.LittleEndian, &data2Size)
	if err != nil {
		return fmt.Errorf("read data2size: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read unknown2: %w", err)
	}

	var data3Size uint32
	err = binary.Read(r, binary.LittleEndian, &data3Size)
	if err != nil {
		return fmt.Errorf("read data3size: %w", err)
	}

	var data4Size uint32
	err = binary.Read(r, binary.LittleEndian, &data4Size)
	if err != nil {
		return fmt.Errorf("read data4size: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read unknown3: %w", err)
	}

	var data5Size uint32
	err = binary.Read(r, binary.LittleEndian, &data5Size)
	if err != nil {
		return fmt.Errorf("read data1size: %w", err)
	}

	var data6Size uint32
	err = binary.Read(r, binary.LittleEndian, &data6Size)
	if err != nil {
		return fmt.Errorf("read data2size: %w", err)
	}

	_, err = r.Seek(int64(12*data1Size+12*data2Size), io.SeekCurrent)
	if err != nil {
		return fmt.Errorf("seek past data1size: %w", err)
	}

	for i := 0; i < int(data3Size); i++ {
		err = binary.Read(r, binary.LittleEndian, &value)
		if err != nil {
			return fmt.Errorf("read data3flags (%d): %w", i, err)
		}

		err = binary.Read(r, binary.LittleEndian, &value)
		if err != nil {
			return fmt.Errorf("read data3flags seek (%d): %w", i, err)
		}

		_, err = r.Seek(int64(value*4), io.SeekCurrent)
		if err != nil {
			return fmt.Errorf("seek past data3flags (%d): %w", i, err)
		}
	}

	//TODO: move past data 4? skipped?

	for i := 0; i < int(data5Size); i++ {
		_, err = r.Seek(int64(7*4), io.SeekCurrent)
		if err != nil {
			return fmt.Errorf("seek past data5size: %w", err)
		}
	}

	var pvsSize uint16
	err = binary.Read(r, binary.LittleEndian, &pvsSize)
	if err != nil {
		return fmt.Errorf("read pvsSize: %w", err)
	}

	_, err = r.Seek(int64(pvsSize), io.SeekCurrent)
	if err != nil {
		return fmt.Errorf("seek pvsSize: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &value)
	if err != nil {
		return fmt.Errorf("read bytes: %w", err)
	}

	_, err = r.Seek(16, io.SeekCurrent)
	if err != nil {
		return fmt.Errorf("seek unknown: %w", err)
	}

	if v.HasPolygons {
		err = binary.Read(r, binary.LittleEndian, &v.Reference)
		if err != nil {
			return fmt.Errorf("read mesh reference: %w", err)
		}
	}

	return nil
}

func (v *BspRegion) FragmentType() string {
	return "BSP Region"
}
