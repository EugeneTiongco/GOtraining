package main

import (
	"fmt"
	"os"
)

func main() {

	// x := 15
	// y := 12.6784
	// z := 0.0
	//fmt.Printf("x = %5v %3.3f %3.2f\n", x, y, z)

	//Scanf allows capturing values with specific pattern/formatting
	var x, y int
	n, err := fmt.Scanf("data: %v %v", &x, &y)
	if err != nil {
		fmt.Printf("Expecting 2 values but got %v\n", n)
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Printf("x = %v y = %v\n", x, y)
	/*
		var s string
		fmt.Scanf("%v", &s)
		fmt.Printf("s = %v\n", s)
	*/
}
