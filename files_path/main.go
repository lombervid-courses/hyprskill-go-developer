package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func pathBaseAndDirectory() {
	paths := []string{"/example/files/img/goland.svg", "example/files", "..files//img///", "", "///"}

	for _, path := range paths {
		fmt.Printf("Base:%12s | Dir: %s\n", filepath.Base(path), filepath.Dir(path))
	}
}

func fileNameExtension() {
	fmt.Println(filepath.Ext("main"))
	fmt.Println(filepath.Ext("main.go"))
	fmt.Println(filepath.Ext("../files/index.test.js"))
}

func constructingAPath() {
	fmt.Println(filepath.Join("example", "files", "img"))
	fmt.Println(filepath.Join("example", "", "files/img"))

	fmt.Println(filepath.Join("home/User/example", "../../../../files/"))
	fmt.Println(filepath.Join("", ""))
}

func absolutePaths() {
	paths := []string{"/", ".", "", "./among_us"}

	for _, path := range paths {
		abs, err := filepath.Abs(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(abs)
	}

	fmt.Println()
	fmt.Println(filepath.IsAbs("/home/User/GolandProjects/example")) // true
	fmt.Println(filepath.IsAbs("/.bashrc"))                          // true
	fmt.Println(filepath.IsAbs(".bashrc"))                           // false
	fmt.Println(filepath.IsAbs("/"))                                 // true Linux/macOS | false Windows
	fmt.Println(filepath.IsAbs(""))
}

func walkingTheFileTree() {
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			fmt.Println("Directory:", path, "size:", info.Size(), "bytes")
		} else {
			fmt.Println("File:", path, "size:", info.Size(), "bytes")
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// pathBaseAndDirectory()
	// fileNameExtension()
	// // constructingAPath()
	walkingTheFileTree()
}
