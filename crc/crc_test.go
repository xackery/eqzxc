package crc

import (
	"hash/crc32"
	"strings"
	"testing"
)

func TestFilenameCrc(t *testing.T) {

	checks := []struct {
		name string
		crc  int
	}{
		{name: "1dirt.bmp", crc: 248793860},
		{name: "arena.wld", crc: 1894535152},
		{name: "claw.bmp", crc: 64377157},
		{name: "lights.wld", crc: 3698793545},
		{name: "mtfloor.bmp", crc: 1379956244},
		{name: "mtinwall.bmp", crc: 2948166389},
		{name: "objects.wld", crc: 1630624230},
		{name: "palette.bmp", crc: 4073721515},
		{name: "rrock.bmp", crc: 1351749524},
		{name: "rrock2drt1.bmp", crc: 3795167928},
		{name: "nexus.wld", crc: 2852726279},
		{name: "lights.wld", crc: 3698793545},
		{name: "objects.wld", crc: 1630624230},
		{name: "palette.bmp", crc: 4073721515},
		{name: "blackgoo0001.dds", crc: 1706639300},
		{name: "blackgoo0002.dds", crc: 1220618316},
		{name: "blackgoo0003.dds", crc: 1407788340},
		{name: "blackgoo0004.dds", crc: 305203036},
		{name: "blackgoo0005.dds", crc: 152635940},
		{name: "blackgoo0006.dds", crc: 610344364},
		{name: "blackgoo0007.dds", crc: 1061753044},
		{name: "blackgoo0008.dds", crc: 2815515004},
		{name: "blackgoo0009.dds", crc: 3170456580},
		{name: "blackgoo0010.dds", crc: 835082925},
		{name: "collide.dds", crc: 628439601},
		{name: "landing.dds", crc: 2452430868},
		{name: "landingb.dds", crc: 2856158063},
		{name: "nexbannstertrim.dds", crc: 3865848291},
		{name: "nexbannstr301.dds", crc: 606761189},
		{name: "nexbanstertrim3.dds", crc: 790310189},
		{name: "nexcaveroof301.dds", crc: 1958875512},
		{name: "nexcaveroof302.dds", crc: 1505363696},
		{name: "nexcavewall301.dds", crc: 1213697353},
		{name: "nexfloor301.dds", crc: 2822020761},
		{name: "nexstepside301.dds", crc: 218663120},
		{name: "nexsteptop.dds", crc: 3615015223},
		{name: "nexsymbol301.dds", crc: 1020028104},
		{name: "nexsymbol302.dds", crc: 297038656},
		{name: "nexwall301.dds", crc: 2791420829},
		{name: "nexwall301b.dds", crc: 563613987},
		{name: "nexwall301c.dds", crc: 984620123},
		{name: "nexwall305.dds", crc: 3401680509},
		{name: "nexwall305a.dds", crc: 897366872},
		{name: "nexwall305b.dds", crc: 402959568},
		{name: "nexwall305c.dds", crc: 53256616},
		{name: "nexwall306.dds", crc: 3887698421},
		{name: "nexwall306b.dds", crc: 3388806371},
		{name: "nexwall306c.dds", crc: 3537181083},
		{name: "nexwall311.dds", crc: 3912686476},
		{name: "nexwall311b.dds", crc: 2050608617},
		{name: "nexwall311c.dds", crc: 1628553361},
		{name: "pillar1antonica.dds", crc: 338383617},
		{name: "pillar1faydwer.dds", crc: 2948685527},
		{name: "pillar1kunark.dds", crc: 4147645966},
		{name: "pillar1odus.dds", crc: 1317575600},
	}
	for _, check := range checks {
		crc := FilenameCRC32(check.name)
		if crc != uint32(check.crc) {
			t.Fatalf("%s wanted 0x%x, got 0x%x (ieee: 0x%x)", strings.ToLower(check.name), check.crc, crc, crc32.ChecksumIEEE([]byte(check.name)))
		}
	}

}
