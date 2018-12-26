package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func permute(wordSet map[string]bool, word string, l int, r int, size int) {
	if l == r {
		wordSet[word[0:size]] = true
	} else {
		for i := l; i <= r; i++ {
			word = swap(word, l, i)
			permute(wordSet, word, l+1, r, size)
			word = swap(word, l, i)
		}
	}
}

func swap(word string, i int, j int) string {
	letters := []rune(word)
	temp := letters[i]
	letters[i] = letters[j]
	letters[j] = temp
	return string(letters)
}

func readData(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func readWords() map[string]bool {
	wordSet := make(map[string]bool)
	fileContent := readData("words.txt")
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		wordSet[strings.ToLower(line)] = true
	}
	return wordSet
}

func printMatchingWords(words []string, regexpString string) {
	var matcher = regexp.MustCompile(regexpString)
	var englishWords = readWords()
	var matchingWords []string
	for _, w := range words {
		if regexpString == "" || matcher.MatchString(w) {
			if englishWords[w] {
				matchingWords = append(matchingWords, w)
			}
		}
	}
	sort.Strings(matchingWords)
	for _, w := range matchingWords {
		fmt.Println(w)
	}
}

func findWords(letters string, size int) []string {
	var wordSet = make(map[string]bool)
	permute(wordSet, letters, 0, len(letters)-1, size)
	var words []string
	for key := range wordSet {
		words = append(words, key)
	}
	return words
}

func main() {
	letters := flag.String("letters", "", "Letter to find anagrams for.")
	regexpString := flag.String("regexp", "", "Regexp for matching results")
	size := flag.Int("size", len(*letters), "Turn on more output.")
	flag.Parse()

	words := findWords(*letters, *size)
	printMatchingWords(words, *regexpString)
}
