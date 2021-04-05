package main

import (
	"fmt"
	"io/fs"
	"os"
	"testing/fstest"
)

type osFS struct{}

func (o *osFS) Open(name string) (fs.File, error) {
	if name == "/." || name == "./." {
		return nil, fmt.Errorf("invalid name: %s", name)
	}

	return os.Open(name)
}

func main() {
	_, err := os.Create("./file.txt")
	if err != nil {
		panic(err)
	}

	defer os.Remove("./file.txt")

	_, err = os.Create("/file.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove("/file.txt")

	myFS := new(osFS)

	if err = fstest.TestFS(myFS, "file.txt"); err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
}
