package main

import (
	"fmt"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/iterator"
)

func factorial(n int) int {
	iter := iterator.New()
	for i := 1; i <= n; i++ {
		iter.Append(collections.IntValue(i))
	}
	return iter.Product().Int()
}

func main() {
	value := factorial(5)
	fmt.Println(value)
}
