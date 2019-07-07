package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	outNames := make([]string, 0, len(files))
	for _, f := range files {
		if f.IsDir() || printFiles {
			outNames = append(outNames, f.Name())
		}
	}

	// sort.Strings(outNames)
	fmt.Println("Strings:", outNames)

	return nil
}
