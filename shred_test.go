/* Test cases that I chose not to implement but should be in a fuller test suite.
   Mostly didn't implement them since they would just be the already implemented tests with minor modifications
*- Path is a folder instead of a file don't shred
*- File is already open somewhere else so don't shred
*- Very large file should be shreddable
*- Practical file test (i.e. Exe/PDFs/Images) where information is valuable
*- Requires super user priveleges to shred
*/
package main

import (
	"os"
	"testing"
	"crypto/rand"
	"errors"
)

func createTestFile(path string, size int64 ) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	buff := make([]byte, size) // Maybe make it less random
	_, err = rand.Read(buff)
	if err != nil {return err}

	_, err = f.Write(buff)
	return err
}

// Shred a Regular small file to see if it functions correctly
func TestShredRegular(t *testing.T) {
	path := "testRegular.txt"
	err:= createTestFile(path, 526)
	if (err != nil) {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Check file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Failed to find file to be shred")
	}

	// Shred
	err = Shred(path)
	if (err != nil) {
		t.Errorf("Shred function returned the error: %v", err)
	}

	// Check if file is gone after
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("File was not shredded correctly")
	}
}

// File does not exist so should throw an error when shredding
func TestShredNotExist(t *testing.T) {
	path := "testDoesNotExist.txt"

	// Checks file doesn't already exist
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("File does exist")
	}

	// Shred
	if (Shred(path) == nil) {
		t.Errorf("Shred function did not return an error for non-existant file")
	}

	// Check if file is gone after
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("File that shouldn't exist, exists")
	}
}

// Empty file should still be deleted
func TestShredEmptyFile(t *testing.T) {
	path := "testEmpty.txt"
	err:= createTestFile(path, 0)
	if (err != nil) {
		t.Fatalf("Failed to create file: %v", err)
	}
	// Check file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Failed to find file to be shred")
	}
	// Shred
	err = Shred(path)
	if (err != nil) {
		t.Errorf("Shred function returned the error: %v", err)
	}
	// Check if file is gone after
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("File was not shredded correctly")
	}
}

// File does not have writing permissions so don't shred it until user changes that
func TestShredUnwritable(t *testing.T) {
	path := "testRegular.txt"
	err:= createTestFile(path, 526)
	if (err != nil) {
		t.Fatalf("Failed to create file: %v", err)
	}

	// Check file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Failed to find file to be shred")
	}

	if err:= os.Chmod(path, 0444); err != nil {
		t.Fatalf("Failed to make file read-only: %v", err)
	}

	// Shred
	err = Shred(path)
	if (err == nil) {
		// Would be fairly trivial to remove it's priveleges and shred anyways but I'll assume that's not the intent for the function
		t.Errorf("Shred function did not return error when trying to write to read-only file")
	}

	// Check if file was shredded
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Errorf("File was shredded anyways")
	}	

	err = os.Remove(path) // So it doesn't interfere again
	if (err != nil) {
		t.Fatalf("Failed to remove test file")
	}
}