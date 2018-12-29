package main

import "fmt"

// Inline declaration for Array
var ageArrayWithInline = [5]byte{1, 2, 3, 4, 5}
var age2DArrayWithInline = [2][2]byte{{1, 2},
	{3, 4}}

const N = 1

var ageDynamicArray = [3 * N]int{1, 2, 3}
var arrayDeclaration [5]int

// Slice Golden Rule:
// 1. A slice can represent a pre existing array. existinArray[1:3]
// 2. A slice can also be created with hidden arrary with potential to grow dynamically. make([]int, length, capacity)
// 3. append method with always add elements in end to increase capacity.
// 4. In case append is to be used, set length to 0
// 5. More than one slice can point to same array sharing the memory location

// Slicing array from exsiting array
var sliceDemo = ageArrayWithInline[0:3]
var sliceInSlice = sliceDemo[0:2]

// Creating new slice
var newSlice = make([]int, 0, 2)

func main() {
	// Print the inline declared arrays
	fmt.Printf("One dimensional Array %d\r\n", ageArrayWithInline[0])

	// Populating and printing the declared array
	arrayDeclaration[0] = 1
	arrayDeclaration[1] = 2
	arrayDeclaration[2] = 3
	arrayDeclaration[3] = 4
	arrayDeclaration[4] = 5
	fmt.Printf("Declared dimensional Array %d\r\n", arrayDeclaration[3])

	// Populating a new slice. The slice has length 0 but capacity 5. It can grow.
	for i := 0; i < 5; i++ {
		newSlice = append(newSlice, i+1)
		fmt.Printf("Slice Info Length = %d Capacity = %d\r\n", len(newSlice), cap(newSlice))
	}

	fmt.Printf("One dimensional Array Before slice update [0] = %d\t[1] = %d\r\n", ageArrayWithInline[0], ageArrayWithInline[1])
	// Slice representing the pre-existing array.
	sliceDemo[0] = 100
	// Printing arrayt to prove that slice actually represents same array.
	fmt.Printf("One dimensional Array Post slice update %d\r\n", ageArrayWithInline[0])

	sliceInSlice[1] = 10
	// Printing arrayt to prove that slice actually represents same array.
	fmt.Printf("One dimensional Array Post slice in slice update %d\r\n", ageArrayWithInline[1])

	fmt.Printf("Two dimensional Array [%d,%d]\t[%d,%d]\r\n", age2DArrayWithInline[0][0],
		age2DArrayWithInline[0][1], age2DArrayWithInline[1][0], age2DArrayWithInline[1][1])
}
