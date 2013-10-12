package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func getFileList(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {
	flag.Parse()
	root := "."
	if flag.Arg(0) != "" {
		root = flag.Arg(0)
	}
	getFileList(root)
}
