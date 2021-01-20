package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

func Top10(sourceTxt string) []string {
	hashTable := make(map[string]int)

	type row struct {
		word    string
		counter int
	}

	var splitedTxt []string = strings.Fields(sourceTxt)

	for _, w := range splitedTxt {
		hashTable[w]++
	}

	wordsSlice := make([]row, len(hashTable))

	var index int = 0
	for word, counter := range hashTable {
		wordsSlice[index].word = word
		wordsSlice[index].counter = counter
		index++
	}

	sort.SliceStable(wordsSlice, func(i, j int) bool {
		return wordsSlice[i].counter < wordsSlice[j].counter
	})

	const resCount int = 10 // output limit
	var breaker int = resCount

	result := make([]string, 0)
	for x := range wordsSlice {
		result = append(result, wordsSlice[len(wordsSlice)-1-x].word)

		breaker--
		if breaker == 0 { // more safe than use counter in for
			break
		}
	}

	return result
}
