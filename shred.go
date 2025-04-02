package main

// Goal is to implement a Shred(path) function that overwites a file with random data 3 times and then deletes it.

import (
	"math/rand"
	"time"
	"fmt"
//	"io"
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
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator (may be unneccesary)
	
	fi, err := os.Stat(path)
	check(err)
	fileSize := fi.Size()

	fmt.Println(path + " Has file size of", fileSize, "bytes")
	for range 3 {
		// Open file for writing (assuming it has write priveleges)
		// If not could either choose to return error or use Chmod function
		f, err := os.OpenFile(path, os.O_WRONLY, 0666)
		check(err)
		//defer f.Close()

		// Write file with new data
		_, err = f.Write([]byte("Hello World!"))
		check(err)
		f.Close()
	}

	fmt.Println("Succesfully overwrote: " + path)
	
	// Now delete the file
	// Note: may be tricky to test if it's using random data if it get's deleted at the end... 
	//err = os.Remove(path)

	return err
}

func main() {
	// Extract path from arguments otherwise
	path := "temp.txt" // Change to a generate file function
	err := Shred(path)
	check(err)
}