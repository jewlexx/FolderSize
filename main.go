package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func PrintError(err error) {
	fmt.Printf("Error: %s\n", err)
	// os.Exit(1)
}

// I may or may not have Ctrl+C Ctrl+V from stackoverflow.. Shhhhh...
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func printSize(name string, size int64) {
	fmt.Printf("%s %d bytes\n", name, size)
}

func getStringPadding(s string) string {
	maxSize := 20

	if (len(s) + 2) > maxSize {
		return s
	} else {
		return fmt.Sprintf("%s%s", s, strings.Repeat(" ", maxSize-len(s)-2))
	}
}

type FileInfo struct {
	Name string
	Size int64
}

func main() {
	fmt.Println("Getting folder size...")

	args := os.Args[1:]

	if stringInSlice("-h", args) || stringInSlice("--help", args) {
		fmt.Printf("Usage: %s\n", os.Args[0])
		fmt.Println("Flags: ")
		fmt.Println("  -h, --help: Show this help")
		fmt.Println("  --no-sort: Do not sort the files by size")
		fmt.Println("  -t, --total: Get the total directory size")
		os.Exit(0)
	}

	cwd, err := os.Getwd()
	if err != nil {
		PrintError(err)
		os.Exit(1)
	}

	if stringInSlice("--total", args) || stringInSlice("-t", args) {
		size, err := DirSize(cwd)
		if err != nil {
			PrintError(err)
		} else {
			printSize(cwd, size)
			os.Exit(0)
		}
	}

	folders, err := os.ReadDir(cwd)
	PrintError(err)

	files := []FileInfo{}

	for _, folder := range folders {
		if folder.IsDir() {
			size, err := DirSize(filepath.Join(cwd, folder.Name()))
			if err != nil {
				PrintError(err)
			} else {
				files = append(files, FileInfo{
					Name: folder.Name(),
					Size: size,
				})
			}
		} else {
			info, err := os.Stat(folder.Name())
			if err != nil {
				PrintError(err)
			} else {
				files = append(files, FileInfo{
					Name: folder.Name(),
					Size: info.Size(),
				})
			}

		}
	}

	if !stringInSlice("--no-sort", args) {
		sort.SliceStable(files, func(i, j int) bool {
			return files[i].Size > files[j].Size
		})
	}

	for _, file := range files {
		printSize(getStringPadding(file.Name), file.Size)
	}
}
