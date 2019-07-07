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
	err := printSubTree(out, path, printFiles, "")
	if err != nil {
		return err
	}

	return nil
}

func printSubTree(out io.Writer, path string, printFiles bool, prefix string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	amount := len(files)
	lastFileIndex := amount - 1

	for i, f := range files {
		name := f.Name()
		isLastFile := i == lastFileIndex
		if f.IsDir() {
			printLeaf(out, prefix, name, isLastFile)
			newPath := fmt.Sprintf("%s%s%s%s", path, string(os.PathSeparator), name, string(os.PathSeparator))
			var newPrefix string
			if isLastFile {
				newPrefix = fmt.Sprintf("%s%s", prefix, "\t")
			} else {
				newPrefix = fmt.Sprintf("%s|%s", prefix, "\t")
			}
			err := printSubTree(out, newPath, printFiles, newPrefix)
			if err != nil {
				return err
			}
		} else if printFiles {
			printLeaf(out, prefix, name, isLastFile)
		}
	}

	return nil
}

func printLeaf(out io.Writer, prefix string, name string, isLastFile bool) {
	var indent string
	if isLastFile {
		indent = "└───"
	} else {
		indent = "├───"
	}
	fmt.Fprint(out, prefix, indent, name, "\n")
}
