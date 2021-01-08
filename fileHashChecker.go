package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// Main Entrypoint for this application
func main() {

	inputFile := os.Args[1]
	inputHash := os.Args[2]

	computedHash := computeSHA256(inputFile)

	if inputHash == computedHash {
		fmt.Println("Hashes Match")
	} else {
		fmt.Println(fmt.Sprintf("Hashes do not match.  \n\tInput = %s\n\tComputed = %s", inputHash, computedHash))
	}

	return
}

// function to compute a hash based on the specified file
// takes a filename
// returns the computed checksum and and error object
func computeSHA256(filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 1024*1024)
	h := sha256.New()

	for {
		bytesRead, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		h.Write(buf[:bytesRead])
	}

	// return the computed string
	return (hex.EncodeToString(h.Sum(nil)))
}

// function to compute a hash based on the specified file
// takes a filename
// returns the computed checksum and and error object
func computeSHA512(filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 1024*1024)
	h := sha512.New()

	for {
		bytesRead, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		h.Write(buf[:bytesRead])
	}

	// return the computed string
	return (hex.EncodeToString(h.Sum(nil)))
}

// function to compute a hash based on the specified file
// takes a filename
// returns the computed checksum and and error object
func computeSHA1(filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 1024*1024)
	h := sha1.New()

	for {
		bytesRead, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		h.Write(buf[:bytesRead])
	}

	// return the computed string
	return (hex.EncodeToString(h.Sum(nil)))
}

// function to compute a hash based on the specified file
// takes a filename
// returns the computed checksum and and error object
func computeMD5(filename string) string {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 1024*1024)
	h := md5.New()

	for {
		bytesRead, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}

			break
		}

		h.Write(buf[:bytesRead])
	}

	// return the computed string
	return (hex.EncodeToString(h.Sum(nil)))
}
