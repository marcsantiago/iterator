package main

import (
	"fmt"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/iterator"
)

func main() {
	var data = []collections.Data{
		collections.IntValue(1),
		collections.IntValue(4),
		collections.IntValue(2),
		collections.IntValue(3),
	}
	value := iterator.IntoIter(data).
		Inspect(func(d collections.Data) {
			fmt.Printf("about to filter %v\n", d)
		}).
		Filter(func(d collections.Data) bool {
			if d.Int()%2 == 0 {
				return false
			}
			return true
		}).
		Inspect(func(d collections.Data) {
			fmt.Printf("made it through the filter %v\n", d)
		}).
		Fold(collections.IntValue(0), func(result, next collections.OperableData) collections.Data {
			return result.Add(next)
		})
	fmt.Printf("result %v\n", value)

}
