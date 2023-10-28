package main

import (
	"fmt"
	"strings"
)

// declare a map type in order to give it methods, similar to the extension methods in Dart.
type WordCountMap map[string]int

func (w *WordCountMap) addWords(words []string) map[string]int {

	// required to allocate and initialize as the custom map is initially nil and can not be used.
	(*w) = make(WordCountMap)

	// enumerate over an array using short variable declaration and range keyword, analogous to enumerate(list) in Python
	for _, word := range words {
		(*w)[word] += 1
	}

	// type assertion required to return a map
	return map[string]int(*w)
}

func main() {

	var m WordCountMap

	var word_slice []string

	var bible_verse string = "The steps of a good man are ordered by the LORD, And He delights in his way. Though he fall, he shall not be utterly cast down; For the LORD upholds him with His hand. I have been young, and now am old; Yet I have not seen the righteous forsaken, Nor his descendants begging bread."

	word_slice = strings.Fields(bible_verse)

	m.addWords(word_slice)

	fmt.Printf("\nWord Count hash map:\n\n%v", m)

}
