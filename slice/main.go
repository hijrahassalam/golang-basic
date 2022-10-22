package main

import "fmt"

func main() {
	// This is a slice
	var gamingConsole []string

	gamingConsole = append(gamingConsole, "PS4")
	gamingConsole = append(gamingConsole, "PS6")
	gamingConsole = append(gamingConsole, "PS5")


	for _, console := range gamingConsole{
		fmt.Println(console)
	}
}