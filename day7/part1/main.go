package main

/*
  Day 7. Part 1:
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type folder struct {
	name string
	path []string
	size int
}

type filesystem struct {
	folders []folder
	pwd     string
}

func newFolder(name string, fs *filesystem) *folder {
	fs.folders = append(fs.folders, name)
	return &folder{
		name: name,
		path: fs.pwd + "/" + name,
		size: 0,
	}
}

func (fs *filesystem) GetPath(name string) string {
	var result []string

	for _, dir := range fs.folders {
		if dir.name == name {
			return strings.Join(dir.path, "/")
		}
	}
	return nil
}

func (fs *filesystem) CheckFolder(name string) (bool, string) {
	for _, dir := range fs.folders {
		for _, dirname := range dir.path {
			if name == dirname {
				return true, fs.GetPath(name)
			}
		}
	}
	return false, ""
}

func (fld *folder) AddFile(size int) *folder {
	return &folder{
		path: fld.path,
		size: fld.size + size,
	}
}

func (fs *filesystem) Cd(dest string) {
	exists, path := fs.CheckFolder(dest)
	if exists {
		fs.pwd = path
	}
	switch dest {
	case "/":
		fs.pwd = "/"
	case "..":
		full_path := strings.Split(fs.pwd, "/")
		fs.pwd = strings.Join(full_path[:len(full_path)-1], "/")
	default:
		full_path := append(strings.Split(fs.pwd, "/"), dest)
		fs.pwd = strings.Join(full_path, "/")
	}
	fs.pwd = dest
}

func main() {
	inputs, err := openFile("../inputs-day7.txt")
	if err != nil {
		panic(err)
	}

	// input into slice of slices
	var lines [][]string
	for _, line := range inputs {
		command := strings.Split(line, " ")
		lines = append(lines, command)
	}

	// Read and execute commands
	for x := 0; x < len(inputs)-1; x++ {
		stdin := lines[x]

		switch stdin[0] {
		case "$":
			// commands
		case "dir":
			// append subdir to subdir
		}

	}
}

func openFile(file string) ([]string, error) {
	var fileLines []string
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening inputs file")
		return nil, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}
