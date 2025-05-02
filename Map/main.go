package main

import "fmt"

func main() {
	//map [keys] values

	/*#1 first way to declare a map:*/
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	//#2nd way:
	// var colors map[string]string

	//#3rd way:
	//colors := make(map[string]string)

	//colors["white"] = "#ffff"
	//delete(colors, "white")
	colors["yellow"] = "#12314"
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}
