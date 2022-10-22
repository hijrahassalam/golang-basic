// Main in Golang is equivalent with Index.
// For running an apps we need package main and function main

package main

import (
	// Package terbagi menjadi 2, yaitu standar package (fmt) dan package libary yang kita buat
	// pemanggilan package wajib mencantumkan nama module pada file go.mod
	"example/calculation"
	"fmt"
)

func main() {
	fmt.Println("Hello Word! This is basic example with calculation study case")

	resulta := calculation.Add(9,9);
	results := calculation.Subtract(9,9);
	resultm := calculation.Multiply(9,9);
	resultd := calculation.Divide(9,9);

	fmt.Println("9 add 9 equals ",resulta);
	fmt.Println("9 subtract 9 equals ",results);
	fmt.Println("9 multiply 9 equals ",resultm);
	fmt.Println("9 divide 9 equals ",resultd);
}