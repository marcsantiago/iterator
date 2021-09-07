package main

import (
	"fmt"
	"strings"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/collections/counter"
	"github.com/marcsantiago/collections/set"
	"github.com/marcsantiago/iterator"
)

// Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that
// occur more than once in the input string. The input string can be assumed to contain only alphabets
// (both uppercase and lowercase) and numeric digits.

// Note this isn't very efficient given all the copies, however I wanted to show how the other packages
// could connect, Implementing IterTrait for counter and set would make this more effective, perhaps for fun in the future
func main() {
	input := "Indivisibilities"
	data := collections.StringValues(
		strings.Split(
			strings.ToLower(input), "")).Data()
	counterMap := counter.Counter(data)

	iter := iterator.IntoIter(data).
		Filter(func(d collections.Data) bool {
			if value, ok := counterMap.Get(d); ok {
				if value.Int() > 1 {
					return false
				}
			}
			return true
		})

	mySet := set.New()
	for d := range iter.Iterate() {
		mySet.Add(d)
	}
	fmt.Println(len(mySet))
}
