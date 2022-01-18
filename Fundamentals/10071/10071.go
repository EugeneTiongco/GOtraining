package main

import "fmt"

func main() {
	var v, t int

	n, _ := fmt.Scan(&v, &t)

	for n == 2 {
		fmt.Println(v * t * 2)
		n, _ = fmt.Scan(&v, &t)
	}
}
