package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	// map is equal to a object in javascript, in this case is Record<string, string>
	colors := map[string]string {
		"red": "#f00",
		"green": "#0f0",
		"blue": "#00f",
	}

	colors["white"] = "#fff"

	delete(colors, "white")

	// fmt.Println(colors)
	// fmt.Println(colors["red"])
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}