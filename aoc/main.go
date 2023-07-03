package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	bs,err := buildPatternFile("pattern.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(bs)
	}
	pt, err := ioutil.ReadFile("pattern.txt")
	

	rows := strings.Split(string(pt), "\n")
	columns := make([]string, 0)

	// Find the maximum row length
	maxLength := 0
	for _, row := range rows {
		if len(row) > maxLength {
			maxLength = len(row)
		}
	}

	// Transpose the pattern
	for i := 0; i < maxLength; i++ {
		column := ""
		for _, row := range rows {
			if len(row) > i {
				char := row[i]
				if char != ' ' && char != '[' && char != ']' {
					column += string(char)
				}
			}
		}
		if (len(column) > 0) {
			columns = append(columns, column)
		}
	}
	fmt.Println(columns)

	// values := make([]string, 0)
	// for _, column := range columns {
	// 	values = append(values, column)
	// }

	// fmt.Println(values)

	// s1 := stack{}
	// s1.Push("s")
	// fmt.Println(s1.getTop())
}

func buildPatternFile(filename string) (string, error) {
	_, err := ioutil.ReadFile(filename)
	if err == nil {
		return "file already exists", err
		
	} else {
		file, err := os.Open("Actions.txt")
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return "",err
		}
		defer file.Close()

		// Create a new scanner to read the file line by line
		scanner := bufio.NewScanner(file)

		// Read each line until the end of the file
		for scanner.Scan() {
			line := scanner.Text()
			
			bs, err := ioutil.ReadFile(filename)
			if err != nil {
				// Check if the error is due to the file not existing
				if os.IsNotExist(err) {
					// Create the file with an empty line
					err = ioutil.WriteFile(filename, []byte("\n"), 0666)
					if err != nil {
						fmt.Println("Error creating file:", err)
					}
				} else {
					fmt.Println("Error reading file:", err)
					return "",err
				}
			}
			if len(line) != 0 {
				ioutil.WriteFile(filename, []byte(string(bs)+line+string("\n")), 0666)
			} else {
				return "done",nil
			}
		}
		// Check for any errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the file:", err)
		}
		return "",err
	}
}