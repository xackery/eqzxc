package pfs

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	path := "test/nexus.s3d"
	r, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer r.Close()

	pfs, err := Load(r)
	if err != nil {
		t.Fatalf("load pfs: %v", err)
	}
	if pfs == nil {
		t.Fatalf("nil pfs")
	}
	if 0 == 1 {
		extract(t, pfs)
	}
}

func extract(t *testing.T, pfs *Pfs) {
	dataPath := "test/data/"
	err := os.RemoveAll(dataPath)
	if err != nil {
		t.Fatalf("removeall %s: %v", dataPath, err)
	}
	err = os.MkdirAll(dataPath, 0755)
	if err != nil {
		t.Fatalf("mkdirall %s: %v", dataPath, err)
	}
	for _, entry := range pfs.Files {
		path := fmt.Sprintf("%s%s", dataPath, entry.Name)
		err = ioutil.WriteFile(path, entry.Data, os.ModePerm)
		if err != nil {
			t.Fatalf("write %s: %v", path, err)
		}
	}
}
