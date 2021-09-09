package iterator

// https://doc.rust-lang.org/std/iter/trait.Iterator.html
import (
	"github.com/marcsantiago/collections"
)

type IterTrait interface {
	Append(data collections.Data)
	All(f func(d collections.Data) bool) bool
	Any(f func(d collections.Data) bool) bool
	Chain(other IterTrait) IterTrait
	CollectInts() []int
	CollectInt32s() []int32
	CollectInt64s() []int64
	CollectFloat32s() []float32
	CollectFloat64s() []float64
	CollectStrings() []string
	CollectBools() []bool
	Count() int
	Cycle()
	Eq(other IterTrait) bool
	Iterate() <-chan collections.Data
	Filter(f func(d collections.Data) bool) IterTrait
	Find(f func(d collections.Data) bool) collections.Data
	Flatten() IterTrait
	Fold(init collections.OperableData, f func(result collections.OperableData, next collections.OperableData) collections.Data) collections.Data
	Ge(other IterTrait) bool
	Gt(other IterTrait) bool
	Inspect(f func(d collections.Data)) IterTrait
	Last() (index int, data collections.Data)
	Le(other IterTrait) bool
	Lt(other IterTrait) bool
	Map(f func(d collections.Data) collections.Data) IterTrait
	Max() collections.Data
	Min() collections.Data
	Ne(other IterTrait) bool
	Nth(n int) collections.Data
	Partition(f func(d collections.Data) bool) ([]collections.Data, []collections.Data)
	Peak() collections.Data
	Position(f func(d collections.Data) bool) int
	Product() collections.Data
	Reduce(f func(a, b collections.Data) collections.Data) collections.Data
	Take(n int)
	Sum() collections.Data
	Next() (index int, data collections.Data)
	Len() int
}
