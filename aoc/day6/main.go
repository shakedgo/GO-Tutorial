package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteCont, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for i := 3; i < len(byteCont); i++ {
		if isUnique(byteCont[i-3 : i+1]) {
			fmt.Println("Part 1:", i+1)
			break
		}
	}
	for i := 13; i < len(byteCont); i++ {
		if isUnique(byteCont[i-13 : i+1]) {
			fmt.Println("Part 2:", i+1)
			break
		}
	}
}

func isUnique(check []byte) bool {
	set := make(map[byte]bool)

	for _, ns := range check {
		if set[ns] {
			return false
		}
		set[ns] = true
	}

	return true
}
