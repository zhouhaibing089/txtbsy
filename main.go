package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "help" {
		fmt.Println("This program is used to verify whether a process can rewrite itself.")
		os.Exit(0)
	}

	// print pid
	log.Printf("pid: %d", os.Getpid())

	path, err := os.Executable()
	if err != nil {
		log.Fatal("failed to find executable")
	}

	showStat(path)

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open %s: %s", path, err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	file.Close()

	// trying to rewrite it
	newfile, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed to create %s: %s", path, err)
	}
	defer newfile.Close()
	_, err = io.Copy(newfile, bytes.NewReader(data))
	if err != nil {
		log.Fatalf("failed to copy file: %s", err)
	}

	showStat(path)
}

func showStat(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatalf("failed to stat %s: %s", path, err)
	}

	stat, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalf("failed to get underlying stat for %s", path)
	}

	log.Printf("path: %s, dev: %d, inode %d", path, stat.Dev, stat.Ino)
}
