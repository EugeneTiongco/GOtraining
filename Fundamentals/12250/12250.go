package main

import "fmt"

func main() {
	var x string
	count := 1

	n, _ := fmt.Scan(&x)

	for n == 1 && x != "#" {
		switch {
		case x == "HELLO":
			fmt.Printf("Case %v: ENGLISH\n", count)
		case x == "HOLA":
			fmt.Printf("Case %v: SPANISH\n", count)
		case x == "HALLO":
			fmt.Printf("Case %v: GERMAN\n", count)
		case x == "BONJOUR":
			fmt.Printf("Case %v: FRENCH\n", count)
		case x == "CIAO":
			fmt.Printf("Case %v: ITALIAN\n", count)
		case x == "ZDRAVSTVUJTE":
			fmt.Printf("Case %v: RUSSIAN\n", count)
		default:
			fmt.Printf("Case %v: UNKNOWN\n", count)
		}
		count++
		n, _ = fmt.Scan(&x)
	}
}
