package main

import "fmt"

//PI value in package
const PI float32 = 3.1412

func main1() {
	nameMap := make(map[string]int)
	nameMap["gaurav"] = 37
	nameMap["Angad"] = 10
	nameMap["Angad1"]++
	x := []int{10, 30}
	const (
		pi int = 10
	)

	fmt.Println(x)
	fmt.Println(pi)

	for name, age := range nameMap {
		if age > 10 {
			fmt.Printf("name "+name+" == %d\r\n", age)
		} else {
			fmt.Print("Minor name " + name + " == ")
			fmt.Println(age)
		}
	}
	test_dage()
}

func test_dage() {

}
