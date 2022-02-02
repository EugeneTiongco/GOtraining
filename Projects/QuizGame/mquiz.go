// Package main provides a quiz tool
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {

	filename := flag.String("csv", "problems.csv", "text or input file")
	itemNum := flag.Int("n", 10, "amount of questions")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	if filepath.Ext(*filename) != ".csv" {
		log.Fatalf("Incorrect database format for %v. It should be in .csv", *filename)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse csv file: %s", err)
	}
	if len(lines) < *itemNum {
		log.Fatalf("Insufficient questions found in database. It needs at least %v questions", *itemNum)
	}

	lines = ShuffleQuestions(lines)

	score := TakeTest(*itemNum, lines)
	fmt.Printf("You answered %v out of %v questions correctly.\n", score, *itemNum)

}

//TakeTest calculates and returns the user score.
func TakeTest(itemNum int, lines [][]string) int {
	score := 0

	for i := 0; i < itemNum; i++ {
		var a string
		fmt.Printf("%v = ", lines[i][0])
		fmt.Scan(&a)

		if a == lines[i][1] {
			score++
		}
	}
	return score
}

//ShuffleQuestions randomizes the questions.
func ShuffleQuestions(lines [][]string) [][]string {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})

	return lines
}
