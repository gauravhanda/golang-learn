package main

import "fmt"

func main() {
	// Octal and hexa decimal
	var i, j int = 021, 0x21

	// Floating point
	PI := 311412.e-5
	const salute = "Mr"

	var rawString = `Gaurav Handa\n\r\ris good`
	var specialString = "Gaurav Handa\t is \n \x41 \u0041 \U00000041 \101 good\u65e5本\U00008a9e\t\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"

	var hexRune2, hexRune4, unicodeRune, octalRune = '\x41', '\u0041', '\U00000041', '\101'

	// Calling the method in same package but different file
	var x = add(10, 30)

	// Finally printing the variables
	fmt.Printf("x = %d \t i=%d\t j = %d\t product = %d\t PI = %f\r\n", x, i, j, product(3, 2), PI)
	fmt.Printf("Unicode hexRune2 = %3d\t hexRune4 = %3d\t unicodeRune = %3d\t octalRune = %3d\r\n",
		hexRune2, hexRune4, unicodeRune, octalRune)
	fmt.Println("Name = " + rawString + " Escaped String = " + specialString)

}

// Method to caluclate the product
func product(x int, y int) int { return x * y }
