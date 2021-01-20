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

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

type appConfig struct {
	hashType string
	filename string
	compare  string
}

var cfg appConfig

// Main Entrypoint for this application
func main() {

	// quit if this fails
	if parseCommandline() > 0 {
		return
	}

	// now determine how to proceed
	if len(cfg.compare) == 0 {
		computeSingleHash()
	} else {
		compareHashStrings()
	}

}

// Can use this to redner output to a table
/*

 */

func parseCommandline() int {

	// define the flags to start with
	hashPTR := flag.String("hash", "SHA256", "The hash we want to use")
	filenamePTR := flag.String("filename", "", "The file we want to hash")
	comparePTR := flag.String("compare", "", "The hash we want to compare")

	// parse the arges and mapp them
	flag.Parse()

	// now assign the values to our struct for later use
	if len(*hashPTR) == 0 {
		cfg.hashType = "SHA256"
	} else {
		cfg.hashType = *hashPTR
	}

	if len(*filenamePTR) == 0 {
		printUsage()
		return (1)
	} else {
		cfg.filename = *filenamePTR
	}

	if len(*comparePTR) > 0 {
		cfg.compare = *comparePTR
	}

	return (0)
}

// prints the usage syntax
func printUsage() {

	fmt.Println("USAGE:  fileHackChecker -filename <OPTIONS>")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("\t-filename\tThe name of the file we want to hash")
	fmt.Println("\t-hash\t\tThe hash to use - SHA1, SHA256, SHA512, MD5  (Default SHA256)")
	fmt.Println("\t-compare\tThe string of the hash you want to compare against")
	fmt.Print("\n\nEXAMPLE:\n\n")
	fmt.Println("To hash a file using SHA256 and print the compated Hash to the screen")
	fmt.Print("\tfileHashChecker -filename=\"myfile.exe\" -hash=\"SHA256\"\n\n\n")
	fmt.Println("To compute the hash of a selected file and compare it to a specified hash")
	fmt.Print("\tfileHashChecker -filename=\"myfile.exe\" -hash=\"SHA256\" -compare=\"FGHJKWETYU345FGH67\"\n\n\n")

}

func compareHashStrings() {
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

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Filename", "Computed Hash", "Result"})

	if computedHash == cfg.compare {

		result := color.New(color.FgHiGreen).SprintfFunc()
		t.AppendRows([]table.Row{
			{cfg.filename, computedHash, result("OK")},
		})

	} else {

		result := color.New(color.FgHiRed).SprintfFunc()
		t.AppendRows([]table.Row{
			{cfg.filename, computedHash, result("ERROR")},
		})

	}

	t.AppendSeparator()
	t.Render()
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

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Filename", "Hash", "Hash Type"})

	t.AppendRows([]table.Row{
		{cfg.filename, computedHash, cfg.hashType},
	})

	t.AppendSeparator()
	t.Render()

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
