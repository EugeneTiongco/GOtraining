package main

import "fmt"

func main() {
	var x string
	hajj := map[string]string{"Hajj": "Hajj-e-Akbar", "Umrah": "Hajj-e-Asghar"}
	count := 1
	n, _ := fmt.Scan(&x)

	for n == 1 && x != "*" {
		fmt.Printf("Case %v: %v\n", count, hajj[x])
		count++
		n, _ = fmt.Scan(&x)
	}
}
