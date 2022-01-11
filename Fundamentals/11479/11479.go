package main

import "fmt"

func main() {
	var T, x, y, z int
	count := 1
	fmt.Scan(&T)
	if T < 20 && T > 0 {
		n, _ := fmt.Scan(&x, &y, &z)
		for n == 3 && count <= T {
			if TIT(x, y, z) == true {
				switch {
				case isScalene(x, y, z) == true:
					fmt.Printf("Case %v: Scalene\n", count)
				case countEqualSides(x, y, z) == 1:
					fmt.Printf("Case %v: Isosceles\n", count)
				case countEqualSides(x, y, z) == 3:
					fmt.Printf("Case %v: Equilateral\n", count)
				}
			} else {
				fmt.Printf("Case %v: Invalid\n", count)
			}
			count++
			if count <= T {
				fmt.Scan(&x, &y, &z)
			}
		}
	} else {
		fmt.Println("Error: Input a positive integer less than 20")
	}
}

// Triangle Inequality Theorem
func TIT(x int, y int, z int) bool {
	if x+y > z && x+z > y && y+z > x {
		return true
	}
	return false
}

func isScalene(x int, y int, z int) bool {
	if x != y && x != z && y != z {
		return true
	}
	return false
}

func countEqualSides(x int, y int, z int) int {
	var identical int
	if x == y {
		identical++
	}
	if y == z {
		identical++
	}
	if x == z {
		identical++
	}
	return identical
}
