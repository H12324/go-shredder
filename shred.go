package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

// Shred overwrites the file at the given path with random data and then deletes it.
func Shred(path string) error {	
	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("file doesn't exist: %v", err)
	}
	fileSize := fi.Size()

	for range 3 {
		// Open file for writing (assuming it has write priveleges)
		// If not could either choose to return error or use Chmod function
		// Going to choose to just allow error for now
		f, err := os.OpenFile(path, os.O_WRONLY, 0666)
		if (err != nil) {
			return fmt.Errorf("couldn't open file for writing: %v", err)
		}

		// Write file with new data
		buff := make([]byte, fileSize)
		_, err = rand.Read(buff)
		if (err != nil) {
			f.Close()
			return fmt.Errorf("couldn't fill buffer with random data %v", err)
		}

		_, err = f.Write(buff)
		f.Close()
		if (err != nil) {
			
			return fmt.Errorf("couldn't write to file %v", err)
		}
	}
	
	// Now delete the file
	err = os.Remove(path)
	if (err != nil) {
		return fmt.Errorf("couldn't delete the file %v", err)
	}

	return err
}

func main() {
	// Extract path from arguments otherwise
	path := "temp.txt" // Change to a generate file function
	err := Shred(path)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Succesfully shredded: " + path)
	}
}