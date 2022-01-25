package main

import (
	"fmt"
	"math"
)

const (
	b = iota
	c
	g
)

var (
	binColor = [][3]int{{b, c, g}, {b, g, c}, {c, b, g}, {c, g, b}, {g, b, c}, {g, c, b}}
	binCode  = []byte{'B', 'C', 'G'}
)

func main() {
	var bin [3][3]int
	var idx int
	total := 0
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d%d%d", &bin[i][b], &bin[i][g], &bin[i][c])
		total += bin[i][b] + bin[i][g] + bin[i][c]
	}
	for total != 0 {
		fmt.Printf("STARTING TO COUNT: %v\n", total)
		minMove := math.MaxInt32
		for i, v := range binColor {
			if move := total - bin[0][v[0]] - bin[1][v[1]] - bin[2][v[2]]; move < minMove {
				minMove = move
				idx = i
			}
		}
		fmt.Printf("%c%c%c %d\n", binCode[binColor[idx][0]], binCode[binColor[idx][1]], binCode[binColor[idx][2]], minMove)
		total = 0
		for i := 0; i < 3; i++ {
			_, err := fmt.Scanf("%d%d%d", &bin[i][b], &bin[i][g], &bin[i][c])
			total += bin[i][b] + bin[i][g] + bin[i][c]
			if err != nil {
				total = 0
			}
		}
		fmt.Println(total)
	}
}
