//Main package creates a program that counts instances of words from files.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

type WordCounter struct {
	mu    sync.Mutex
	words map[string]int
}

func main() {
	wc := WordCounter{words: make(map[string]int)}

	if len(os.Args) < 2 {
		log.Fatalf("Expected at least one filename")
	}

	filenames := os.Args[1:]
	ch := make(chan []string)
	//V2 version concurrently runs GetWords and CountWords at the same time vs V1 where GetWords
	//is run concurrently, finished, then CountWords is then run concurrently
	for _, filename := range filenames {
		go GetWords(filename, ch)
		wordList := make([]string, 0)
		wordList = append(wordList, <-ch...)
		go wc.CountWords(wordList)
	}

	time.Sleep(time.Second)
	sortedWords := wc.GetSortedWords()

	for i := range sortedWords {
		fmt.Printf("%v %v \n", sortedWords[i], wc.GetWordAmount(sortedWords[i]))
	}

}

//GetWords reads the words from the files concurrently using a channel.
func GetWords(filename string, ch chan []string) {
	wordSlice := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open a file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	defer file.Close()

	removeSpecial := regexp.MustCompile(`(?m)[^a-z0-9]`)
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		word = removeSpecial.ReplaceAllString(word, "")
		wordSlice = append(wordSlice, word)
	}
	time.Sleep(3 * time.Second)
	ch <- wordSlice

	//defer close(ch)
}

//CountWords takes each word from the list to be prepared for counting.
func (wc *WordCounter) CountWords(wordList []string) {
	for _, word := range wordList {
		wc.Inc(word)
	}
}

//Inc increments the word instances in a map.
func (wc *WordCounter) Inc(key string) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	if key != "" {
		wc.words[key]++
	}
}

//GetWordAmount returns the amount of instances per word
func (wc *WordCounter) GetWordAmount(key string) int {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.words[key]
}

//GetSortedWords returns an array of strings containing the words arranged in alphabetical order.
func (wc *WordCounter) GetSortedWords() []string {
	keys := make([]string, 0, len(wc.words))
	for key := range wc.words {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}
