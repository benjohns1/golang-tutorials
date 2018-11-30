package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	// Asserting an error type, and getting more info from struct fields
	f, err := os.Open("/test.txt")
	// 'err' implements Go's error interface, fmt.Println() calls interface's Error() method internally
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
	} else {
		fmt.Println(f.Name(), "opened successfully")
	}

	// Asserting an error type, and using methods to get more info
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
	} else {
		fmt.Println(addr)
	}

	// Direct error comparison
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
	} else {
		fmt.Println("matched files", files)
	}
}
