package main

import "fmt"

func main() {
	title := "Golang the best languanges"

	// basic for loop
	// for _, letter := range title {
	// 	fmt.Println("letter : ", string(letter))
	// }

	// only prints even index
	// for index, letter := range title {
	// 	if index % 2 == 0 {
	// 		fmt.Println("index : ", index,  "letter : ", string(letter))
	// 	}
	// }

	// only prints vocal letter
	for index, letter := range title {
		// using if 
		// if string(letter) == "a" || string(letter) == "i" || string(letter) == "u" || string(letter) == "e" || string(letter) == "o" {
		// 	fmt.Println("index : ", index,  "letter : ", string(letter))
		// }
		
		// using switch
		letterString := string(letter)
		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index : ", index,  "letter : ", string(letter))
		}
	}
}