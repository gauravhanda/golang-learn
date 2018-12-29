package tempconv

import "fmt"

// Celcius to represent the temperature
type Celcius float64
type Fahrenhite float64

func (c Celcius) String(x string) string {
	return "gaurav" + x
}

func init() {
	fmt.Println("Initialized")
}
