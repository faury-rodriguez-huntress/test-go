package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"test-go/error"
)

func main() {
	fileInfo, err := os.Stat(";32	4;lj	234	5l;kj some-non-existent-file")
	err = error.OsError{WrappedError: err}
	
	fmt.Printf("%v\n", err)
	fmt.Printf("ErrExist: %v\n", errors.Is(err, fs.ErrNotExist))
	fmt.Printf("ErrInvalid: %v\n", errors.Is(err, fs.ErrInvalid))
	fmt.Printf("%v", fileInfo)
}
 