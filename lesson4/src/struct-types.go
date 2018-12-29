package main

import "fmt"

type Salary struct {
	fName  string
	amount float32
}

//Employee First class to be exported
type Employee struct {
	firstName string
	lastName  string
	reportees []string
	Salary    // -> Field namne will be same as the struct name
	fullName  func() string
}

func main() {
	var data = make([]string, 2, 5)
	data[0] = "Gaurav"

	var x = Employee{"Gaurav", "Handa", data, Salary{"Angad", 100}, func() string { return "Hello" }}
	x.fName = "Angad H" // Promoted Field
	fmt.Printf("First Name = %s, Last Name = %s,  Salary = %f,  Salary.Name = %s, Func = %s\r\n", x.firstName,
		x.lastName, x.Salary.amount, x.fName, x.fullName())
}
