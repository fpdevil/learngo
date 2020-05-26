package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// find empty files and write them to a file
// optimized version for handling large number of files previous
// version allocates a new backing array everytime in  loop which
// is inefficient
//
// There are 2 ways to achieve this:
//
// 1st. Use a Heuristic:
// Multiply the number of files with the reserved average file
// name lengths across common filesystems (it is simple, but
// wastes unnecessary memory)
//
// https://en.wikipedia.org/wiki/Comparison_of_file_systems#Limits
//
// BTRFS   255 bytes
// exFAT   255 UTF-16 characters
// ext2    255 bytes
// ext3    255 bytes
// ext3cow 255 bytes
// ext4    255 bytes
// FAT32   255 bytes
// NTFS    255 characters
// XFS     255 bytes
//
// assuming 255 bytes + 1 byte for newline character
// again on the assumption that a single byte represents
// a single character and 1 byte for '\n'
//
// total := len(files) * (255 + 1)
//
// 2nd. Actually calculate the filename lengths.
// (this is more efficient memory wise but wastes a lot of
// cu cycles)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Give directory as argument")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
	}

	// 1st option
	// total size set
	// total := len(files) * 256
	// fmt.Printf("Total allocated size: %d bytes\n", total)

	// 2nd. option
	// To be exact, find the total size of all the empty files
	var total int
	for _, file := range files {
		if file.Size() == 0 {
			// add an extra 1 (byte) to accommodate for the
			// newline character '\n'
			total += len(file.Name()) + 1
		}
	}
	fmt.Printf("Total allocated size: %d bytes\n", total)

	// moving away from the nil slice now
	// var names []byte
	// now append will insert the data at the beginning of the slice
	// as we have given the length as 0
	names := make([]byte, 0, total)

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()

			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	// for file writing
	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get some details about the output file
	fi, err := os.Stat("out.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("file %s has size %d\n", fi.Name(), fi.Size())
	fmt.Printf("%s", names)
}
