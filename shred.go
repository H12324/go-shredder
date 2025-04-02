package main

import (
	"crypto/rand"
	"fmt"
	"os"
)


// Adapted from https://gobyexample.com/reading-files
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Shred overwrites the file at the given path with random data and then deletes it.
func Shred(path string) error {	
	fi, err := os.Stat(path)
	check(err)
	fileSize := fi.Size()

	for range 3 {
		// Open file for writing (assuming it has write priveleges)
		// If not could either choose to return error or use Chmod function
		// Going to choose to just allow error for now
		f, err := os.OpenFile(path, os.O_WRONLY, 0666)
		check(err)

		// Write file with new data
		buff := make([]byte, fileSize)
		_, err = rand.Read(buff)
		check(err)

		_, err = f.Write(buff)
		check(err)

		err = f.Close()
		check(err)
	}
	
	// Now delete the file
	err = os.Remove(path)
	check(err)
	
	fmt.Println("Succesfully shredded: " + path)

	return err
}

func main() {
	// Extract path from arguments otherwise
	path := "temp.txt" // Change to a generate file function
	err := Shred(path)
	check(err)
}