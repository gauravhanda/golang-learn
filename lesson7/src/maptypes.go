package main

import (
	"fmt"
)

/*
  Refer to https://golang.org/ref/spec#Map_types

  A map is an unordered group of elements of one type, called the element type,
  indexed by a set of unique keys of another type, called the key type.  The
  value of an uninitialized map is nil.

  The comparison operators == and != must be fully defined for operands of the
  key type; thus the key type must not be a function, map, or slice. If the key
  type is an interface type, these comparison operators must be defined for the
  dynamic key values; failure will cause a run-time panic.

*/

type String2StringMap map[string]string
type String2IntMap map[string]int
type Id2EmployeeMap map[int]EmpDetails

type EmpDetails interface {
	fullName() string
}

type Employee struct {
	firstName string
	lastName  string
	EmpDetails
}

type UKEmployee struct {
	Employee
}

func main() {
	runStringToStringMap()
	runStringToIntMap()
	runIntToEmployeeMap()
}

// String to string map
func runStringToStringMap() {
	var string2stringMap = make(String2StringMap)
	string2stringMap["USA"] = "Chicago"
	string2stringMap["USA"] = "Washington DC"
	string2stringMap["China"] = "Beijing"
	fmt.Printf("Map says %s \r\n", string2stringMap["USA"])

}

// String to string map
func runStringToIntMap() {
	var string2IntMap = make(String2IntMap)
	string2IntMap["USA"] = 100
	string2IntMap["China"] = 200
	fmt.Printf("Map says %d \r\n", len(string2IntMap))
	delete(string2IntMap, "USA")
	fmt.Printf("Map says %d \r\n", len(string2IntMap))
	fmt.Printf("Map says %d \r\n", string2IntMap["USA"])
	fmt.Printf("Map says %d \r\n", len(string2IntMap))
}

func (t Employee) fullName() string { return t.lastName + ", " + t.firstName }

// Overrides the fullName for UK employees with first name as start
func (t UKEmployee) fullName() string { return t.firstName + ", " + t.lastName }

func runIntToEmployeeMap() {
	var intToEmployeeMap = make(Id2EmployeeMap)
	// See the inheritance in work
	var simpleEmployee EmpDetails = Employee{firstName: "Gaurav", lastName: "Handa"}
	intToEmployeeMap[100] = simpleEmployee
	var ukEmployee EmpDetails = UKEmployee{Employee{firstName: "Gaurav", lastName: "Handa"}}
	intToEmployeeMap[200] = ukEmployee

	fmt.Printf("Map says %s \r\n", intToEmployeeMap[100].fullName())
	fmt.Printf("Map says %s \r\n", intToEmployeeMap[200].fullName())

}
