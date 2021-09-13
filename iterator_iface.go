package iterator

// https://doc.rust-lang.org/std/iter/trait.Iterator.html
import (
	"github.com/marcsantiago/collections"
)

type BaseIterTrait interface {
	All(f func(d collections.Data) bool) bool
	Any(f func(d collections.Data) bool) bool
	Count() int
	Max() collections.Data
	Min() collections.Data
	Len() int
}

type IterTraitSlice interface {
	BaseIterTrait
	Append(data collections.Data)
	Chain(other IterTraitSlice) IterTraitSlice
	Collect() []collections.Data
	CollectInts() []int
	CollectInt32s() []int32
	CollectInt64s() []int64
	CollectFloat32s() []float32
	CollectFloat64s() []float64
	CollectStrings() []string
	CollectBools() []bool
	Cycle()
	Eq(other IterTraitSlice) bool
	Filter(f func(d collections.Data) bool) IterTraitSlice
	Find(f func(d collections.Data) bool) collections.Data
	Fold(init collections.OperableData, f func(result collections.OperableData, next collections.OperableData) collections.Data) collections.Data
	FoldIntoMap(init IterTraitMap, f func(m IterTraitMap, key collections.Data) IterTraitMap) IterTraitMap
	Ge(other IterTraitSlice) bool
	Gt(other IterTraitSlice) bool
	Iterate() <-chan collections.Data
	Inspect(f func(d collections.Data)) IterTraitSlice
	Last() (index int, data collections.Data)
	Le(other IterTraitSlice) bool
	Lt(other IterTraitSlice) bool
	Map(f func(d collections.Data) collections.Data) IterTraitSlice
	Ne(other IterTraitSlice) bool
	Nth(n int) collections.Data
	Partition(f func(d collections.Data) bool) ([]collections.Data, []collections.Data)
	Peak() collections.Data
	Position(f func(d collections.Data) bool) int
	Product() collections.Data
	Reduce(f func(a, b collections.Data) collections.Data) collections.Data
	Take(n int)
	Sum() collections.Data
	Next() (index int, data collections.Data)
	Zip(other []collections.Data) []collections.Element
	ZipIntoMap(other []collections.Data) IterTraitMap
}

type IterTraitMap interface {
	collections.Map
	BaseIterTrait
	Chain(other IterTraitMap) IterTraitMap
	Collect() []collections.Element
	CollectKeys() []collections.Data
	CollectValues() []collections.Data
	KeysToIterSlice() IterTraitSlice
	ValuesToIterSlice() IterTraitSlice
	Eq(other IterTraitMap) bool
	Filter(f func(d collections.Data) bool) IterTraitMap
	FilterKeys(f func(d collections.Data) bool) IterTraitMap
	Find(f func(d collections.Data) bool) (item collections.Element, ok bool)
	FindByValue(f func(d collections.Data) bool) (item collections.Element, ok bool)
	Ge(other IterTraitMap) bool
	Gt(other IterTraitMap) bool
	Inspect(f func(d collections.Element)) IterTraitMap
	Le(other IterTraitMap) bool
	Lt(other IterTraitMap) bool
	Map(f func(d collections.Data) collections.Data) IterTraitMap
	Ne(other IterTraitMap) bool
	Reduce(f func(a, b collections.Data) collections.Data) collections.Data
	Next() (data collections.Element, ok bool)
}
