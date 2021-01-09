package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

// Main Entrypoint for this application
func main() {

	inputFile := os.Args[1]

	computeSingleHash(inputFile, "SHA256")
	computeSingleHash(inputFile, "SHA512")
	computeSingleHash(inputFile, "SHA1")
	computeSingleHash(inputFile, "MD5")

	return
}

// function to compute a single hash and return it
func computeSingleHash(inputFile string, hashType string) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Filename", "Hash", "Hash Type"})

	if hashType == "SHA256" {
		h := sha256.New()
		computedHash := computeHash(inputFile, h)

		t.AppendRows([]table.Row{
			{inputFile, computedHash, "SHA256"},
		})

	} else if hashType == "SHA512" {
		h := sha512.New()
		computedHash := computeHash(inputFile, h)

		t.AppendRows([]table.Row{
			{inputFile, computedHash, "SHA512"},
		})
	} else if hashType == "SHA1" {
		h := sha1.New()
		computedHash := computeHash(inputFile, h)

		t.AppendRows([]table.Row{
			{inputFile, computedHash, "SHA1"},
		})
	} else if hashType == "MD5" {
		h := md5.New()
		computedHash := computeHash(inputFile, h)

		t.AppendRows([]table.Row{
			{inputFile, computedHash, "MD5"},
		})
	}

	t.AppendSeparator()
	t.Render()

	return
}

// function to compute a hash based on the specified file
// takes a filename and a hash.Hash from the specific crypto class
// returns the computed checksum and and error object
func computeHash(filename string, h hash.Hash) string {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// defer closing the file until we exit this method
	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 1024*1024)
	// here

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
