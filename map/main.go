package main

import "fmt"

func main(){
	// define a map
	var myMap map[string]int
	// initialize the map
	myMap = map[string]int{}

	// filling the map
	myMap["C#"] = 9
	myMap["Dart"] = 10
	myMap["Vue"] = 20

	fmt.Println(myMap)


	myLang := map[string]string{"Bahasa":"Indonesia", "Language":"English"}
	myLang["Allughoh"] = "Arab"

	fmt.Println(myLang)
	fmt.Println("====================")

	// looping in a map
	for _, value := range myLang{
		fmt.Println(value)
	}

	// find a key in a map
	// returning 2 object, the key and also true/false
	key, isAvailable := myLang["Bahasa"]

	fmt.Println(key, isAvailable)

}