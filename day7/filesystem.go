package day7

import (
	"strconv"
	"strings"
)

type Directory struct {
	directories []*Directory
	files       []int
	name        string
	parent      *Directory
	size        int
}

func (d *Directory) calcSize() int {
	var sz int = 0
	for _, file := range d.files {
		sz += file
	}

	for _, directory := range d.directories {
		sz += directory.calcSize()
	}

	d.size = sz
	return sz
}

func buildFileSystem(terminalOutput []string) *Directory {
	var root *Directory = nil
	var currDirectory *Directory
	for _, termLine := range terminalOutput {

		if termLine[0] == '$' {
			// command
			lineSplit := strings.Split(termLine, " ")
			cmd := lineSplit[1]
			var arg string

			if len(lineSplit) == 3 {
				arg = lineSplit[2]
			}

			if cmd == "cd" {
				if arg == ".." {
					// traverse "up" a directory
					if currDirectory.parent != nil {
						currDirectory = currDirectory.parent
					}
				} else if root == nil {
					// initialize root and set current directory to root
					root = &Directory{
						name:   arg,
						parent: nil,
					}
					currDirectory = root
				} else {
					// change to new directory
					for _, directory := range currDirectory.directories {
						if directory.name == arg {
							currDirectory = directory
						}
					}
				}
			}
		} else if strings.HasPrefix(termLine, "dir") {
			// directory
			lineSplit := strings.Split(termLine, " ")
			newDirName := lineSplit[1]
			newDir := &Directory{
				name:   newDirName,
				parent: currDirectory,
			}
			currDirectory.directories = append(currDirectory.directories, newDir)
		} else {
			// file
			lineSplit := strings.Split(termLine, " ")
			newFileSize, _ := strconv.Atoi(lineSplit[0])
			currDirectory.files = append(currDirectory.files, newFileSize)
		}
	}
	return root
}
