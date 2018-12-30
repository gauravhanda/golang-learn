package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumbers("X")
	go printNumbers("Y")
	printAfter10SecondsAsync()
	defer waitForFinish()
	fmt.Println("Execution and defer started....")
}

func waitForFinish() {
	fmt.Println("defered execited")
	fmt.Scanln()
}

func printNumbers(flag string) {
	for i := 1; i <= 100; i++ {
		fmt.Printf("type = %s value = %d\r\n", flag, i)
		time.Sleep(250 * time.Millisecond)
	}
}

func printAfter10SecondsAsync() {
	go func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("1 second has passed")
		defer func() { fmt.Println("Defered function called") }()
	}()
}
