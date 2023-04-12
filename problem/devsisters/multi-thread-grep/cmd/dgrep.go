package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	keyword         string
	relativePath    string
	targetFilePaths []string

	maxGoroutine = runtime.NumCPU()

	wg      sync.WaitGroup
	workIDX atomic.Int32
)

const (
	whiteSpace = ' '
)

func parseCommandLine() error {
	if len(os.Args) != 3 {
		return errors.New("invalid command line args")
	}
	keyword = os.Args[1]
	relativePath = os.Args[2]
	return nil
}

func isDir() (bool, error) {
	info, err := os.Stat(relativePath)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func findTarget() error {
	return filepath.WalkDir(relativePath, walkFn)
}

func walkFn(path string, d fs.DirEntry, err error) error {
	if !d.IsDir() {
		targetFilePaths = append(targetFilePaths, path)
	}
	return nil
}

func dgrep(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineNum int
	for scanner.Scan() {
		lineNum++
		if strings.Contains(scanner.Text(), keyword) {
			fmt.Sprint(file.Name(), whiteSpace, lineNum)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
