package main

import "fmt"

func main() {
	//Arrays, hard assign
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	//Declare AND Assign
	fruitArr := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[2:3])
}
