package main

import "fmt"

func main() {
	// colors := make(map[string]string)
	// colors["black"] = "#000000"
	// delete(colors, "black")
	colors := map[string]string{
		"white": "#FFFFFF",
		"black": "#000000",
		"red":   "#FF0000",
	}
	delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k + " in hex is: " + v )
	}
}