package main

import (
	"strings"
	"testing"

	"github.com/marcsantiago/collections"
	"github.com/marcsantiago/collections/counter"
)

//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkCount-12    	 4875249	       246.8 ns/op	     120 B/op	       4 allocs/op
//BenchmarkCount-12    	 4857900	       243.7 ns/op	     120 B/op	       4 allocs/op
//BenchmarkCount-12    	 5023674	       237.5 ns/op	     120 B/op	       4 allocs/op
func BenchmarkCount(b *testing.B) {
	input := "Indivisibilities"
	data := collections.StringValues(
		strings.Split(
			strings.ToLower(input), "")).Data()
	counterMap := counter.Counter(data)
	for i := 0; i < b.N; i++ {
		count(counterMap)
	}
}

//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkCountAlt-12    	  492591	      2402 ns/op	    1417 B/op	       7 allocs/op
//BenchmarkCountAlt-12    	  462390	      2413 ns/op	    1417 B/op	       7 allocs/op
//BenchmarkCountAlt-12    	  467856	      2460 ns/op	    1417 B/op	       7 allocs/op
func BenchmarkCountAlt(b *testing.B) {
	input := "Indivisibilities"
	data := collections.StringValues(
		strings.Split(
			strings.ToLower(input), "")).Data()
	for i := 0; i < b.N; i++ {
		countAlt(data)
	}
}

//cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
//BenchmarkCountAlt2-12    	  441657	      2622 ns/op	    1249 B/op	       6 allocs/op
//BenchmarkCountAlt2-12    	  425917	      2635 ns/op	    1249 B/op	       6 allocs/op
//BenchmarkCountAlt2-12    	  431150	      2661 ns/op	    1249 B/op	       6 allocs/op
func BenchmarkCountAlt2(b *testing.B) {
	input := "Indivisibilities"
	data := collections.StringValues(
		strings.Split(
			strings.ToLower(input), "")).Data()
	for i := 0; i < b.N; i++ {
		countAlt2(data)
	}
}
