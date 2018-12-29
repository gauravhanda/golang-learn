package main

import (
	"fmt"
	"os"
)

func main2() {
	fName, lName := "angad", "Gauirav"
	test := "gauirav"

	fmt.Println("Hello world " + fName + lName + test)

	for index, args := range os.Args[0:] {
		fmt.Print(index)
		fmt.Println(" = " + args)
	}
	fmt.Printf("%f", PI)

}
