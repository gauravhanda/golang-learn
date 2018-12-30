package main

import "fmt"

func main() {
	if l := 5; l < 15 {
		fmt.Println("x is Less than threshhold")
	}

	fmt.Printf("%s\r\n", switchTest(-11, 11))

	var arraySlice = make([]int, 10, 10)
	iterateIndex(arraySlice)
	iterateSlice(arraySlice)

	var mymap = map[int]string{
		1: "Gaurav",
		2: "Angad",
	}

	iterateMap(mymap)

}

func iterateIndex(arraySlice []int) {
	for x := 1; x <= 10; x = x + 1 {
		arraySlice[x-1] = x
	}
}

func iterateSlice(arraySlice []int) {
	for _, value := range arraySlice {
		fmt.Printf(" value = %d\r\n", value)
	}

}

func iterateMap(mymap map[int]string) {
	for key, value := range mymap {
		fmt.Printf("Key = %d , value = %s\r\n", key, value)
	}
}

func switchTest(x, y int) string {
	switch result := runSwitchAsIfElse(x, y); {
	case result == -1:
		fmt.Printf("rsult = %d", result)
		return "X LESS THAN Y"
	case result == 1:
		fmt.Printf("rsult = %d", result)
		return "X GREATER THAN Y"
	default:
		fmt.Printf("rsult = %d", result)
		return "X EQUAL Y"
	}

	return "UNKNOWN"
}
func runSwitchAsIfElse(x int, y int) int {
	switch {
	case x < y:
		return -1
	case x > y:
		return 1

	}
	return 0
}
