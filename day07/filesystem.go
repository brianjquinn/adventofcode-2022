package day07

import (
	"math"
	"strconv"
	"strings"
)

type Filesystem struct {
	root *Directory
}

func (f Filesystem) totalDiskSpace() int {
	return 70000000
}

func (f Filesystem) unusedSpace() int {
	return f.totalDiskSpace() - f.root.size
}

func (f *Filesystem) calculateAllDirectorySizes() {
	f.root.calculateSize()
}

func (f Filesystem) sumOfSizeOfDirsWithSizeLessThan(size int) int {
	return sumHelper(size, f.root, 0)
}

func sumHelper(size int, directory *Directory, sum int) int {
	if directory.size <= size {
		sum += directory.size
	}

	for _, dir := range directory.directories {
		sum = sumHelper(size, dir, sum)
	}
	return sum
}

func (f Filesystem) sizeOfSmallestDirectoryToDelete(minSize int) int {
	return sizeHelper(minSize, f.root, math.MaxInt)
}

func sizeHelper(minSize int, directory *Directory, smallestSize int) int {
	if directory.size >= minSize && directory.size < smallestSize {
		smallestSize = directory.size
	}

	for _, dir := range directory.directories {
		smallestSize = sizeHelper(minSize, dir, smallestSize)
	}

	return smallestSize
}

type Directory struct {
	directories []*Directory
	files       []int
	name        string
	parent      *Directory
	size        int
}

func (d *Directory) calculateSize() {
	for _, file := range d.files {
		d.size += file
	}

	for _, directory := range d.directories {
		directory.calculateSize()
		d.size += directory.size
	}
}

func buildFileSystem(terminalOutput []string) *Filesystem {
	var filesystem *Filesystem = nil
	var currDir *Directory

	for _, termLine := range terminalOutput {
		// command
		if termLine[0] == '$' {
			lineSplit := strings.Split(termLine, " ")
			cmd := lineSplit[1]
			var arg string

			if len(lineSplit) == 3 {
				arg = lineSplit[2]
			}

			if cmd == "cd" {
				// initialize the file system since this is the initial "cd /" command
				if arg == "/" && filesystem == nil {
					rootDir := &Directory{
						name:   arg,
						parent: nil,
					}

					filesystem = &Filesystem{
						root: rootDir,
					}
					currDir = rootDir
				} else if arg == ".." {
					// traverse "up" a directory
					if currDir.parent != nil {
						currDir = currDir.parent
					}
				} else {
					// change to a new directory specified by arg
					for _, directory := range currDir.directories {
						if directory.name == arg {
							currDir = directory
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
				parent: currDir,
			}
			currDir.directories = append(currDir.directories, newDir)
		} else {
			// file
			lineSplit := strings.Split(termLine, " ")
			newFileSize, _ := strconv.Atoi(lineSplit[0])
			currDir.files = append(currDir.files, newFileSize)
		}
	}
	return filesystem
}
