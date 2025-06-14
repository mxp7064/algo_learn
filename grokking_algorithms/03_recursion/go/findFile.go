package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	filePath, err := findFileInDirWhile("/home/panca/Desktop/bla", "panca.txt", 3)
	if err != nil {
		log.Fatal(err)
	}
	if filePath != "" {
		fmt.Println("file found at path: ", filePath)
	} else {
		fmt.Println("file not found")
	}
}

// findFileInDir looks up dir contents and returns the full path of the target if found inside dir
// if not found, it returns empty string
// this is recursive way to find a file in a directory
func findFileInDir(dir string, target string) (string, error) {
	entries, derr := os.ReadDir(dir)
	if derr != nil {
		return "", fmt.Errorf("failed on reading directory %s with: %s", dir, derr)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			filePath, err := findFileInDir(fullPath, target)
			if err != nil {
				return "", fmt.Errorf("failed on findFileInDir for %s with: %s", fullPath, err)
			}
			if filePath != "" {
				return filePath, nil
			}
			continue
		}
		if entry.Name() == target {
			return fullPath, nil
		}
	}
	return "", nil
}

// this is iterative way
func findFileInDirWhile(root string, target string, depthLimit int) (string, error) {
	type dirEntry struct {
		path  string
		depth int
	}

	stack := []dirEntry{{
		path:  root,
		depth: 1,
	}} // root will always be a directory, stack stores full directory paths

	for len(stack) > 0 {
		// pop last element of the stack, with stack we get DFS (depth first search)
		// if we implement is as a queue we would get BFS (breath first search)
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.depth > depthLimit {
			return "", errors.New("depth too much")
		}

		// process current
		entries, err := os.ReadDir(current.path)
		if err != nil {
			return "", fmt.Errorf("failed on os.ReadDir for %s with: %s", current, err)
		}

		for _, entry := range entries { // entries are children of current
			// current is always directory
			fullPath := filepath.Join(current.path, entry.Name())
			if entry.IsDir() {
				stack = append(stack, dirEntry{
					path:  fullPath,
					depth: current.depth + 1,
				})
			} else if entry.Name() == target {
				return fullPath, nil
			}
			// ignore other files which are not named 'target'
		}
	}
	return "", nil
}
