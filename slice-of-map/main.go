package main

import "fmt"

func main() {
	// Slice can be contained by map (key : value) data
	students := []map[string]string{
		{"name":"Hijrah", "score":"AAA"},
		{"name":"Ahmad", "score":"BBB"},
		{"name":"Dali", "score":"CCC"},
	}

	for _, student := range students{
		fmt.Println(student["name"], " - ", student["score"])
	}
}