package main

import (
	"fmt"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/iterator"
)

func main() {
	var data = []collections.Data{
		collections.IntValue(10),
		collections.IntValue(20),
		collections.IntValue(5),
		collections.IntValue(-23),
		collections.IntValue(0),
	}
	value := iterator.IntoIter(data).
		Reduce(func(a, b collections.Data) collections.Data {
			if a.Equal(b) || a.Greater(b) {
				return a
			}
			return b
		})
	fmt.Printf("result %v\n", value)
}
