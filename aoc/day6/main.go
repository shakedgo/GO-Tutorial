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
	// content := string(byteCont)

	for i, _ := range byteCont {
		if i > 3 {
			if isUnique([]byte{byteCont[i], byteCont[i+1], byteCont[i+2], byteCont[i+3]}) {
				fmt.Println(string(byteCont[i]), string(byteCont[i+1]), string(byteCont[i+2]), string(byteCont[i+3]))
				fmt.Println(i + 4)
				break
			}
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
