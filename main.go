package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/xackery/eqzxc/pfs"
)

func main() {
	start := time.Now()
	err := run()
	fmt.Printf("finished in %0.2f seconds", time.Since(start).Seconds())
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
}

func run() error {
	dirs, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("readdir: %w", err)
	}
	for _, entry := range dirs {
		if entry.Type().IsDir() {
			continue
		}
		err = extract(entry.Name())
		if err != nil {
			return fmt.Errorf("extract %s: %w", entry.Name(), err)
		}
	}
	//pfs.Load
	return nil
}

func extract(path string) error {
	if filepath.Ext(path) != ".s3d" {
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	content, err := pfs.Load(f)
	if err != nil {
		return fmt.Errorf("load: %w", err)
	}
	outpath := fmt.Sprintf("_%s/", path)
	err = os.MkdirAll(outpath, 0755)
	if err != nil {
		return fmt.Errorf("mkdirall %s: %w", path, err)
	}
	fmt.Println(outpath)
	for _, entry := range content.Files {
		fPath := fmt.Sprintf("%s%s", outpath, entry.Name)
		fmt.Println(fPath)
		err = ioutil.WriteFile(fPath, entry.Data, os.ModePerm)
		if err != nil {
			return fmt.Errorf("write %s: %w", fPath, err)
		}
	}
	return nil
}
