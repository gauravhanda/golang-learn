package main

import "fmt"

// This remembers me of typedef in C/C++ world
type (
	Int16        = int16
	Integer16    = int16
	BigInteger16 int16
	Employee     = struct{ name string }
)

func main() {
	var (
		k int16        = 1
		x Int16        = 2
		y Integer16    = 3
		z BigInteger16 = 4
	)

	// z = x BigInteger16 type is not equal to int16. Its a new defined type
	var gaurav = Employee{name: "Handa"}
	fmt.Printf("%d\t %d\t %d\t %d\t %s\r\n", x, y, z, k, gaurav.name)
}
