package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {
	var k int
	var division, p point

	fmt.Scan(&k)

	for k != 0 {
		fmt.Scan(&division.x, &division.y)

		for i := 0; i < k; i++ {
			fmt.Scan(&p.x, &p.y)

			switch {
			case p.x == division.x || p.y == division.y:
				fmt.Println("divisa")
			case p.x > division.x:
				if p.y > division.y {
					fmt.Println("NE")
				} else {
					fmt.Println("SE")
				}
			case p.x < division.x:
				if p.y > division.y {
					fmt.Println("NO")
				} else {
					fmt.Println("SO")
				}
			}
		}
		fmt.Scan(&k)
	}
}
