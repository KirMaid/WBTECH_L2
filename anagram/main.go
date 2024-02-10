package anagram

import (
	"fmt"
	"sort"
	"strings"
	_ "unicode"
)

func FindAnagrams(words *[]string) *map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range *words {
		lower := strings.ToLower(word)
		runes := []rune(lower)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		sorted := string(runes)
		anagrams[sorted] = append(anagrams[sorted], word)
	}

	for key, set := range anagrams {
		if len(set) > 1 {
			sort.Strings(set)
			anagrams[key] = set
		} else {
			delete(anagrams, key)
		}
	}

	return &anagrams
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := FindAnagrams(&words)
	for key, set := range *anagrams {
		fmt.Printf("%s: %v\n", key, set)
	}
}
