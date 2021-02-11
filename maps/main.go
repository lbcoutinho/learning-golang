package main

import "fmt"

func main() {
	// 2 ways to create an empty map
	// 1st way
	var emptyMap1 map[string]string
	fmt.Println(emptyMap1)
	// 2nd way
	var emptyMap2 = make(map[string]string)
	fmt.Println(emptyMap2)

	// Create and initialize map
	colors := map[string]string{
		"red":"#ff0000",
		"green":"#00ff00",
		"blue":"#0000ff",
	}
	fmt.Println(colors)

	// Add an element to map
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	// Delete an element from map
	delete(colors, "red")
	printMap(colors)
}

func printMap(m map[string]string)  {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
