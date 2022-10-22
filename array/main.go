package main

import "fmt"

func main() {
	// //define an array
	// var bahasa [3]string
	// bahasa[0] = "Arab"
	// bahasa[1] = "China"
	// bahasa[2] = "Jepang"

	// fmt.Println(bahasa)
	// length := len(bahasa)
	// fmt.Println("Panjang array adalah ",length)

	// // define and fill an array 
	// // use [...] if the length is not already counted
	// languages := [...]string{
	// 	"Dart",
	// 	"C#",
	// 	"Ruby",
	// 	"Javascript",
	// }
	// // looping an array
	// for index, lang := range languages{
	// 	fmt.Println("Index ke :", index, lang)
	// }

	fmt.Println("==========")
	// calculating averages in array
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	avgArray(scores)
	fmt.Println("==========")

	// counting score greater than a defined value
	var goodScore []int 
	for _, score := range scores {
		if score >= 90 {
			goodScore = append(goodScore, score)
		}
	}
	fmt.Println(goodScore)
}

func avgArray(array [8]int){
	sum := 0
	length := len(array)
	result := 0
	for _, score := range array{
		sum += score
	}
	result = sum/length
	fmt.Println(result)
}

