package main

import "fmt"

// Public utilities
type Utils struct {
}

func (utils Utils) functionOnAType(radius float32) float32 {
	return 3.1412 * radius * radius
}

// Function type defined
type AddPointer func(int, int) int

func main() {
	fmt.Println("Undestanding the functions ")
	withNoParams()
	withSingleParam("Gaurav")
	fmt.Printf("Inside Function with returnType = %d\r\n", withReturnType(10, 3))
	fmt.Printf("Inside Function with withReturnTypeAllParamsSameType = %d\r\n", withReturnTypeAllParamsSameType(10, 3))

	withVariableParams("Gaurav", "Angad")
	withUnderscore(10, "Handa")

	var sum, diff, prod, div = withMultipleReturns(10, 2)
	fmt.Printf("Multiple Returns %d\t%d\t%d\t%d\r\n", sum, diff, prod, div)
	sum, diff, prod, div = withMultipleReturnsWithNames(20, 5)
	fmt.Printf("Multiple Returns With Names %d\t%d\t%d\t%d\r\n", sum, diff, prod, div)
	var utils = Utils{}
	fmt.Printf("Area of circle = %f\r\n", utils.functionOnAType(1))

	var addPointer AddPointer = withReturnType
	fmt.Printf("Function Pointer Used = %d\r\n", addPointer(10, 30))

	// See carefully. Functions are first class citizens
	fmt.Printf("Function Returned from Function = %d\r\n", withFunctionTypeReturn()(10, 10))
	fmt.Printf("Function Passed as parameter = %d\r\n", withFunctionAsParameter(withReturnType, 15, 15))

}

func withNoParams() {
	fmt.Println("Inside Function with no parameters")
}

func withSingleParam(name string) {
	fmt.Println("Inside Function with One parameters = " + name)
}

func withReturnType(x int, y int) int {
	return x + y
}

func withReturnTypeAllParamsSameType(x, y int) int {
	return x + y
}

func withVariableParams(names ...string) {
	fmt.Printf("Called with variable arguments %d\r\n", len(names))
}

func withUnderscore(_ int, names string) {
	fmt.Printf("Method with underscore %s\r\n", names)
}

func withMultipleReturns(x int, y int) (int, int, int, int) {
	return x + y, x - y, x * y, x / y
}

// See the beauty here
// NOTE func withMultipleReturnsWithNames(x int, y int) (add int, diff int, mul int, div int) {
func withMultipleReturnsWithNames(x, y int) (add, diff, mul, div int) {
	add, diff, mul, div = x+y, x-y, x*y, x/y
	return
}

// Returns a function as paramter
func withFunctionTypeReturn() AddPointer {
	return withReturnType
}

// Inputs a function as parameter
func withFunctionAsParameter(addFunction AddPointer, x, y int) int {
	return addFunction(x, y)
}
