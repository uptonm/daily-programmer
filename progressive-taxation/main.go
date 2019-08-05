package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(tax(1234567))
}

func tax(income int) float64 {
	if income > 100000 {
		return (0.1 * float64(20000)) + (0.25 * float64(70000)) + (0.4 * math.Abs(float64(100000-income)))
	} else if income > 30000 {
		return (0.1 * float64(20000)) + (0.25 * math.Abs(float64(30000-income)))
	} else if income > 10000 {
		return 0.1 * math.Abs(float64(10000-income))
	} else {
		return 0
	}
}
