package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type appConfig struct {
	hashType   string
	filename   string
	folder     string
	recursive  bool
	savefile   string
	fileoutput bool
}

var cfg appConfig

// Main Entrypoint for this application
func main() {

	parseCommandline()

	computeSingleHash()

	/*
		inputFile := os.Args[1]

		computeSingleHash(inputFile, "SHA256")
		computeSingleHash(inputFile, "SHA512")
		computeSingleHash(inputFile, "SHA1")
		computeSingleHash(inputFile, "MD5")
	*/
}

// Can use this to redner output to a table
/*

 */

func parseCommandline() {

	// define the flags to start with
	hashPTR := flag.String("hash", "SHA256", "The hash we want to use")
	filenamePTR := flag.String("filename", "", "The file we want to hash")
	folderPTR := flag.String("folder", "", "The fodler we want to iterate through")
	outfilePTR := flag.String("outfile", "", "The file we want to save all the hashs to")
	recursivePTR := flag.Bool("recursive", false, "Do we want to decend the whole folder")
	fileoutputPTR := flag.Bool("fileoutput", false, "Do we want to save to a file or render to the screen")

	// parse the arges and mapp them
	flag.Parse()

	// now assign the values to our struct for later use
	cfg.hashType = *hashPTR
	cfg.filename = *filenamePTR
	cfg.folder = *folderPTR
	cfg.fileoutput = *fileoutputPTR
	cfg.recursive = *recursivePTR
	cfg.savefile = *outfilePTR

	fmt.Println("filename=", *filenamePTR, "hash=", cfg.hashType)

}

// function to compute a single hash and return it
func computeSingleHash() {

	var computedHash string

	if cfg.hashType == "SHA256" {
		h := sha256.New()
		computedHash = computeHash(cfg.filename, h)

	} else if cfg.hashType == "SHA512" {
		h := sha512.New()
		computedHash = computeHash(cfg.filename, h)

	} else if cfg.hashType == "SHA1" {
		h := sha1.New()
		computedHash = computeHash(cfg.filename, h)

	} else if cfg.hashType == "MD5" {
		h := md5.New()
		computedHash = computeHash(cfg.filename, h)
	}

	// now we need to see how the output was specifed and deal with that
	if cfg.fileoutput {
		// save this to a file
	} else {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Filename", "Hash", "Hash Type"})

		t.AppendRows([]table.Row{
			{cfg.filename, computedHash, cfg.hashType},
		})

		t.AppendSeparator()
		t.Render()
	}
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
