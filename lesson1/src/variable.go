package main

import "fmt"

var total = 10

func main3() {
	fName, lName := "George", "Michaels"
	var isSexy, name = true, " Trump"

	udpateValue(&total)

	fmt.Printf("Hi %s %s %t %s %d\n", fName, lName, isSexy, name, total)
}

func add(x int, y int) int {
	return x + y
}

func udpateValue(x *int) {
	*x = 60
}
