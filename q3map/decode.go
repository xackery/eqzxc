package q3map

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Decode(r io.ReadSeeker) (*Q3Map, error) {
	m := &Q3Map{}
	err := m.parse(r)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	return m, nil
}

func (m *Q3Map) parse(r io.ReadSeeker) error {
	buf := bufio.NewScanner(r)

	lastAction := ""
	var lastEntity *Entity
	var lastBrush *Brush
	var lastBrushDef *BrushDef
	var err error

	lineNumber := 0
	for buf.Scan() {
		lineNumber++
		line := strings.TrimSpace(buf.Text())
		if strings.Contains(line, "//") {
			line = line[0:strings.Index(line, "//")]
		}

		if strings.HasPrefix(line, `"origin"`) {
			origin := line[strings.Index(line, `"origin"`)+10 : len(line)-1]
			coords := strings.Split(origin, " ")
			if len(coords) != 3 {
				return fmt.Errorf("line %d expected 3 coordinates, got %d on %s", lineNumber, len(coords), line)
			}
			var val float64
			val, err = strconv.ParseFloat(coords[0], 32)
			if err != nil {
				return fmt.Errorf("line %d parse 0 %s: %w", lineNumber, coords[0], err)
			}
			lastEntity.Origin.X = float32(val)
			val, err = strconv.ParseFloat(coords[1], 32)
			if err != nil {
				return fmt.Errorf("line %d parse 0 %s: %w", lineNumber, coords[1], err)
			}
			lastEntity.Origin.Y = float32(val)
			val, err = strconv.ParseFloat(coords[2], 32)
			if err != nil {
				return fmt.Errorf("line %d parse 0 %s: %w", lineNumber, coords[2], err)
			}
			lastEntity.Origin.Z = float32(val)
			continue
		}
		if strings.HasPrefix(line, `"classname"`) {
			lastEntity.ClassName = line[strings.Index(line, `"classname"`)+13 : len(line)-1]
			continue
		}
		if strings.HasPrefix(line, `"light"`) {
			lastEntity.Light = line[strings.Index(line, `"light"`)+9 : len(line)-1]
			continue
		}
		if strings.HasPrefix(line, `"angle"`) {
			lastEntity.Light = line[strings.Index(line, `"angle"`)+9 : len(line)-1]
			continue
		}
		if strings.HasPrefix(line, "(") && lastAction == "brushDef" {
			lastBrushDef = &BrushDef{}
			defs := strings.Split(line, " ")
			var val float64
			val, err = strconv.ParseFloat(defs[1], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 1 X failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[0].X = float32(val)
			val, err = strconv.ParseFloat(defs[2], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 1 Y failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[0].Y = float32(val)
			val, err = strconv.ParseFloat(defs[3], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 1 Z failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[0].Z = float32(val)

			val, err = strconv.ParseFloat(defs[6], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 2 X failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[1].X = float32(val)
			val, err = strconv.ParseFloat(defs[7], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 2 Y failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[1].Y = float32(val)
			val, err = strconv.ParseFloat(defs[8], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 2 Z failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[1].Z = float32(val)

			val, err = strconv.ParseFloat(defs[11], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 3 X failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[2].X = float32(val)
			val, err = strconv.ParseFloat(defs[12], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 3 Y failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[2].Y = float32(val)
			val, err = strconv.ParseFloat(defs[13], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 3 Z failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[2].Z = float32(val)

			val, err = strconv.ParseFloat(defs[17], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 4 X failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[3].X = float32(val)
			val, err = strconv.ParseFloat(defs[18], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 4 Y failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[3].Y = float32(val)
			val, err = strconv.ParseFloat(defs[19], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 4 Z failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[3].Z = float32(val)

			val, err = strconv.ParseFloat(defs[22], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 5 X failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[4].X = float32(val)
			val, err = strconv.ParseFloat(defs[23], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 5 Y failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[4].Y = float32(val)
			val, err = strconv.ParseFloat(defs[24], 32)
			if err != nil {
				return fmt.Errorf("line %d parse point 5 Z failed: %w", lineNumber, err)
			}
			lastBrushDef.Points[4].Z = float32(val)

			lastBrushDef.Texture = defs[27]
			lastBrushDef.Unk1 = defs[28]
			lastBrushDef.Unk2 = defs[29]
			lastBrushDef.Unk3 = defs[30]

			lastBrush.Defs = append(lastBrush.Defs, lastBrushDef)
			continue
		}
		switch line {
		case "":
			continue
		case "{":
			if lastEntity == nil {
				lastEntity = &Entity{}
				continue
			}
			if lastBrush == nil {
				lastBrush = &Brush{}
				continue
			}
			if lastAction == "brushDef" {
				continue
			}
			return fmt.Errorf("unhandled { at %d", lineNumber)
		case "}":
			if lastBrush != nil {
				lastEntity.Brushes = append(lastEntity.Brushes, lastBrush)
				lastBrush = nil
			}
			if lastEntity != nil {
				m.Entities = append(m.Entities, lastEntity)
				lastEntity = nil
			}
		case "brushDef":
			lastAction = "brushDef"
		default:
			return fmt.Errorf("unhandled '%s' at %d", line, lineNumber)
		}
	}

	return nil
}
