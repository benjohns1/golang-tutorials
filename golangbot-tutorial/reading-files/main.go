package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/gobuffalo/packr"
)

func main() {
	// Files are relative to location of compiled go binary
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	} else {
		fmt.Println("Contents of file (relative):", string(data))
	}

	// Read filename from command line --fpath="path/to/file"
	fptr := flag.String("fpath", "test.txt", "file path to read")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
	data2, err2 := ioutil.ReadFile(*fptr)
	if err2 != nil {
		fmt.Println("File reading error", err2)
	} else {
		fmt.Println("Contents of file (cmd arg):", string(data2))
	}

	// Bundle file with compiled binary
	box := packr.NewBox("../reading-files")
	data3, err3 := box.FindString("test.txt")
	if err3 != nil {
		fmt.Println("File reading error", err3)
	} else {
		fmt.Println("Contents of file (bundled):", data3)
	}
}
