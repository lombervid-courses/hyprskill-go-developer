package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"runtime"
)

func getCurrentDirectory() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Path:", path)
	fmt.Println()
}

func getFileAttributes() {
	fileInfo, err := os.Stat("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Permission mode:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is directory:", fileInfo.IsDir())

	fmt.Println("Compiler:", runtime.Compiler)
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Printf("GOOS_GOARCH: %s_%s\n\n", runtime.GOOS, runtime.GOARCH)

	osStats(fileInfo) // from `attributes_*.go`
}

func fileModeAndPermissionBits() {
	fileInfo, err := os.Stat("files")
	if err != nil {
		log.Fatal(err)
	}

	mode := fileInfo.Mode()
	fmt.Printf("File perm. bits: %#o\n", mode.Perm())
	fmt.Println("File type bits:", mode.Type())
	fmt.Println("Is regular:", mode.IsRegular())
}

func checkFileExists() {
	fileName := "impostor.png"
	// fileName := "info.txt"
	fileInfo, err := os.Stat(fileName)
	// if os.IsNotExist(err) {
	// 	log.Fatal("The file ", fileName, " does not exist!")
	// }
	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("The file ", fileName, " does not exist!")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Name(), "exists!")
}

func ioFsPackage() {
	fileInfo, err := fs.Stat(os.DirFS("files"), "goland.svg")
	if err != nil {
		log.Fatal(err)
	}

	mode := fileInfo.Mode()
	fmt.Println("File name:", fileInfo.Name())
	fmt.Printf("File perm. bits: %#o\n", mode.Perm())
}

func main() {
	fmt.Println("Getting file attributes")
	// getCurrentDirectory()
	getFileAttributes()
	// fileModeAndPermissionBits()
	// checkFileExists()
	// ioFsPackage()
}
