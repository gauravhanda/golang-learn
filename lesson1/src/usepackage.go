package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	var tempInCelcius tempconv.Celcius = 60.0

	fmt.Println(tempconv.ToFah())
	fmt.Println(tempInCelcius)

}
