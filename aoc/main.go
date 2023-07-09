package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	err := buildPatternFile("pattern.txt")
	if err != nil {
		fmt.Println(err)
	}
	pt, err := ioutil.ReadFile("pattern.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cols := getCols(string(pt));

	stacks := []stack{} // slice of stacks
	for _, column := range cols {
		tempStack := stack{}
		for i := len(column)-1; i >= 0; i-- {
			if (i != len(column)-1) {
				tempStack.Push(string(column[i]))
			}
		}
		stacks = append(stacks, tempStack)
	}
	
	stacks2 := make([]stack, len(stacks))
	copy(stacks2, stacks)

	acts, err := getActions()
	if err != nil {
		os.Exit(1)
	}
	
	applyActions(stacks, stacks2, acts)
	
	tops := getAllTopValues(stacks)
	tops2 := getAllTopValues(stacks2)

	fmt.Println("Part one :", tops)
	fmt.Println("Part Two :", tops2)
}

func buildPatternFile(filename string) (error) {
	_, err := ioutil.ReadFile(filename)
	if err == nil {
		fmt.Println("file already exists")
		return err
	} else {
		file, err := os.Open("Actions.txt")
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return err
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
					return err
				}
			}
			if len(line) != 0 {
				ioutil.WriteFile(filename, []byte(string(bs)+line+string("\n")), 0666)
			} else {
				return nil
			}
		}
		// Check for any errors during scanning
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the file:", err)
		}
		return err
	}
}

func getCols(pt string) []string {
	rows := strings.Split(pt, "\n")
	columns := make([]string, 0)

	// Find the maximum row length
	maxLength := 0
	for _,row := range rows {
		if len(row) > maxLength {
			maxLength = len(row)
		}
	}

	// Transpose the pattern
	for i := 0; i < maxLength; i++ {
		column := ""
		for _,row := range rows {
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
	return columns
}

type action struct {
	qty, from, to int
}

func getActions() ([]action,error) {
	actions := []action{}

	file, err := os.Open("Actions.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return actions, err
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Read each line until the end of the file
	for scanner.Scan() {
		line := scanner.Text()
		act := action{}
		fmt.Sscanf(line, "move %d from %d to %d", &act.qty, &act.from, &act.to)
		// fmt.Println(act)
		empty := action{0,0,0}
		if act != empty {
			actions = append(actions, act)
		}
	}
	// fmt.Println(actions)
	return actions,nil
}

func applyActions(stacks, stacks2 []stack, acts []action) {
	for _, act := range acts {
		tempStack2 := stack{}
		for i := 0; i < act.qty; i++ {
			si, bol := stacks[act.from-1].Pop()
			si2, _ := stacks2[act.from-1].Pop()
			if bol {
				stacks[act.to-1].Push(si)
				tempStack2.Push(si2)
			}
		}
		for i := len(tempStack2) - 1; i >= 0; i-- {
			stacks2[act.to-1].Push(tempStack2[i])
		}
	}
}

func getAllTopValues(stacks []stack) []string {
	var tops []string
	for _, s := range stacks {
		tops = append(tops, s.getTop())
	}
	return tops
}