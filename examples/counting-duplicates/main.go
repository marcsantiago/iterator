package main

import (
	"fmt"
	"strings"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/collections/counter"
	"github.com/marcsantiago/iterator"
)

// Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that
// occur more than once in the input string. The input string can be assumed to contain only alphabets
// (both uppercase and lowercase) and numeric digits.

func main() {
	input := "Indivisibilities"
	data := collections.StringValues(
		strings.Split(
			strings.ToLower(input), "")).Data()
	counterMap := counter.Counter(data)
	fmt.Println(count(counterMap))
	fmt.Println(countAlt(data))
	fmt.Println(countAlt2(data))

}

func count(m collections.Map) int {
	count := iterator.NewMapIterFromMap(m).
		Filter(func(d collections.Data) bool {
			if d.Int() > 1 {
				return true
			}
			return false
		}).Len()
	return count
}

func countAlt(data []collections.Data) int {
	count := iterator.IntoIter(
		iterator.IntoIter(data).
			FoldIntoMap(iterator.NewMapIter(), func(m iterator.IterTraitMap, key collections.Data) iterator.IterTraitMap {
				if value, ok := m.Get(key); ok {
					m.Set(key, value.(collections.OperableData).Add(collections.IntValue(1)))
					return m
				}
				m.Set(key, collections.IntValue(1))
				return m
			}).CollectValues()).
		Filter(func(d collections.Data) bool {
			if d.Int() > 1 {
				return true
			}
			return false
		}).Len()
	return count
}

func countAlt2(data []collections.Data) int {
	count :=
		iterator.IntoIter(data).
			FoldIntoMap(iterator.NewMapIter(), func(m iterator.IterTraitMap, key collections.Data) iterator.IterTraitMap {
				if value, ok := m.Get(key); ok {
					m.Set(key, value.(collections.OperableData).Add(collections.IntValue(1)))
					return m
				}
				m.Set(key, collections.IntValue(1))
				return m
			}).
			Filter(func(d collections.Data) bool {
				if d.Int() > 1 {
					return true
				}
				return false
			}).Len()
	return count
}
