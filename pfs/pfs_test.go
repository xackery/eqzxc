package pfs

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/xackery/eqzxc/wld"
	"github.com/xackery/wd"
)

func TestDecodeAndEncode(t *testing.T) {
	filename := "arena"
	path := fmt.Sprintf("test/%s.s3d", filename)
	r, err := os.Open(path)
	if err != nil {
		t.Fatalf("open %s: %v", path, err)
	}
	defer r.Close()

	pfs, err := Decode(r)
	if err != nil {
		t.Fatalf("load pfs: %v", err)
	}
	if pfs == nil {
		t.Fatalf("nil pfs")
	}
	if 1 == 1 {
		extract(t, fmt.Sprintf("test/_%s.s3d/", filename), pfs)
	}

	r.Seek(0, io.SeekStart)

	f, err := os.Create(fmt.Sprintf("test/%s_out.s3d", filename))
	if err != nil {
		t.Fatalf("create %s_out.s3d: %s", filename, err.Error())
	}
	defer f.Close()

	err = pfs.Encode(&wd.WriteDebug{Input: r})
	if err != nil {
		t.Fatalf("save: %s", err.Error())
	}
}

func extract(t *testing.T, path string, pfs *Pfs) {
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatalf("removeall %s: %v", path, err)
	}
	err = os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("mkdirall %s: %v", path, err)
	}
	for _, entry := range pfs.Files {
		fPath := fmt.Sprintf("%s%s", path, entry.Name)
		err = ioutil.WriteFile(fPath, entry.Data, os.ModePerm)
		if err != nil {
			t.Fatalf("write %s: %v", fPath, err)
		}

		if filepath.Ext(entry.Name) == ".wld" {
			wld, err := wld.Decode(bytes.NewReader(entry.Data))
			if err != nil {
				t.Fatalf("wld load %s: %v", fPath, err)
			}
			mPath := fmt.Sprintf("%s.toml", fPath)
			data := fmt.Sprintf(`
shortname = "%s"`, wld.ShortName)
			err = ioutil.WriteFile(mPath, []byte(data), os.ModePerm)
			if err != nil {
				t.Fatalf("write %s: %v", mPath, err)
			}
		}
	}
}
