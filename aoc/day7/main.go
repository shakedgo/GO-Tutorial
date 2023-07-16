package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	readDirs := make(map[string]bool)
	var currDir []string

	for _, line := range lines {
		words := strings.Split(line, " ")
		if words[0] == "$" {
			if words[1] == "cd" {
				if words[2] == "/" {
					currDir = currDir[:0]
				} else if words[2] == ".." {
					currDir = currDir[:len(currDir)-1]
				} else {
					currDir = append(currDir, words[2])
				}
			}
			// else if words[1] == "ls" {
			// 	currDir = append(currDir, words[2])
			// }
		}
		if words[0] == "dir" {
			if !readDirs[words[1]] {
				readDirs[words[1]] = true
			}
		} else {
			// add number to directory
		}
	}

	// Check if any error occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

func isUnique(s []string) bool {
	set := make(map[string]bool)

	for _, ns := range s {
		if set[ns] {
			return false
		}
		set[ns] = true
	}

	return true
}
