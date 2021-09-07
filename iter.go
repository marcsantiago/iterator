package iterator

import (
	"github.com/marcsantiago/collections"
)

type _direction uint8

const (
	_notSet _direction = iota
	_forward
	_backwards
)

type Iter struct {
	currentIdx  int
	values      []collections.Data
	direction   _direction
	shouldCycle bool
}

var _ IterTrait = (*Iter)(nil)

func New() *Iter {
	return &Iter{}
}

func IntoIter(data []collections.Data) *Iter {
	return &Iter{
		values: data,
	}
}

// Append not very Rust like, but adding a convenience to append into internal data without needing to pass in a full copy
func (i *Iter) Append(data collections.Data) {
	i.values = append(i.values, data)
}

// Next returns the next element in the collection and its enumerated position and moves the cursor forward
// If Next is called Last cannot be called for this copy of Iter, returns -1 as an index when the iterator is fully consumed
func (i *Iter) Next() (int, collections.Data) {
	if i.direction == _notSet {
		i.direction = _forward
	} else if i.direction == _backwards {
		panic("Last was called for this iterator, therefore Next cannot be called, create a new iterator to use Next")
	}

	if !i.shouldCycle {
		if i.currentIdx >= len(i.values) {
			return -1, nil
		}
	}

	curIdx, data := i.currentIdx, i.values[i.currentIdx%len(i.values)]
	i.currentIdx++
	return curIdx, data
}

// All returns true if all the values match the predicate
func (i *Iter) All(f func(d collections.Data) bool) bool {
	for idx := range i.values {
		if !f(i.values[idx]) {
			return false
		}
	}
	return true
}

// Any returns true if any the value matches the predicate
func (i *Iter) Any(f func(d collections.Data) bool) bool {
	for idx := range i.values {
		if f(i.values[idx]) {
			return true
		}
	}
	return false
}

// Chain Takes two iterators and creates a new iterator over both in sequence
func (i *Iter) Chain(other IterTrait) IterTrait {
	values := make([]collections.Data, 0, len(i.values)+other.Len())
	for _, value := range i.values {
		values = append(values, value)
	}

	for {
		pos, value := other.Next()
		if pos == -1 {
			break
		}
		values = append(values, value)
	}

	return &Iter{
		values: values,
	}
}

// CollectInts converts []collections.Data into the concrete type []int
func (i *Iter) CollectInts() []int {
	a := make([]int, len(i.values))
	for ii := 0; ii < len(i.values); ii++ {
		a[ii] = i.values[ii].Int()
	}
	return a
}

// CollectInt32s converts []collections.Data into the concrete type []int32
func (i *Iter) CollectInt32s() []int32 {
	a := make([]int32, len(i.values))
	for ii := 0; ii < len(i.values); ii++ {
		a[ii] = i.values[ii].Int32()
	}
	return a
}

// CollectInt64s converts []collections.Data into the concrete type []int64
func (i *Iter) CollectInt64s() []int64 {
	a := make([]int64, len(i.values))
	for ii := 0; ii < len(i.values); ii++ {
		a[ii] = i.values[ii].Int64()
	}
	return a
}

// CollectFloat32s converts []collections.Data into the concrete type []float32
func (i *Iter) CollectFloat32s() []float32 {
	a := make([]float32, len(i.values))
	for ii := 0; ii < len(i.values); ii++ {
		a[ii] = i.values[ii].Float32()
	}
	return a
}

// CollectFloat64s converts []collections.Data into the concrete type []float64
func (i *Iter) CollectFloat64s() []float64 {
	a := make([]float64, len(i.values))
	for ii := 0; ii < len(i.values); ii++ {
		a[ii] = i.values[ii].Float64()
	}
	return a
}

// CollectStrings converts []collections.Data into the concrete type []string
func (i *Iter) CollectStrings() []string {
	a := make([]string, len(i.values))
	for ii := 0; ii < len(i.values); ii++ {
		a[ii] = i.values[ii].String()
	}
	return a
}

// CollectBools converts []collections.Data into the concrete type []bools
func (i *Iter) CollectBools() []bool {
	a := make([]bool, len(i.values))
	for ii := 0; ii < len(i.values); ii++ {
		a[ii] = i.values[ii].Bool()
	}
	return a
}

// Count consumes the iterator and returns the count of non nil items
func (i *Iter) Count() int {
	if len(i.values) == 0 {
		return 0
	}
	var count int
	for ii := 0; ii < len(i.values); ii++ {
		if i.values[ii] != nil {
			count++
		}
	}
	return count
}

// Cycle enables Next or Last to iterate forever cycling through the items in order seen
func (i *Iter) Cycle() {
	i.shouldCycle = true
}

// Eq determines if this iterator is the same as the other iterator
func (i *Iter) Eq(other IterTrait) bool {
	if i.Len() != other.Len() {
		return false
	}

	selfCh := i.Iterate()
	otherCh := other.Iterate()
	for {
		sV, selfOK := <-selfCh
		otherV, otherOK := <-otherCh
		if sV != otherV {
			return false
		}

		if !selfOK && !otherOK {
			break
		}
	}

	return true
}

// Iterate returns a channel of values that can be ranged over
func (i *Iter) Iterate() <-chan collections.Data {
	ch := make(chan collections.Data)
	go func() {
		for idx := range i.values {
			ch <- i.values[idx]
		}
		close(ch)
	}()
	return ch
}

// Filter removes all values by which the comparison function returns true
func (i *Iter) Filter(f func(d collections.Data) bool) IterTrait {
	var k int
	for j := range i.values {
		if !f(i.values[j]) {
			i.values[k] = i.values[j]
			k++
			continue
		}
	}
	i.currentIdx = 0
	i.values = i.values[:k]
	return i
}

// Find returns the first value that matches the find function, else it returns nil
func (i *Iter) Find(f func(d collections.Data) bool) collections.Data {
	for ii := 0; ii < len(i.values); ii++ {
		if f(i.values[ii]) {
			return i.values[ii]
		}
	}
	return nil
}

// Flatten is currently a stub as Iter only supports 1 dimensional slices
func (i *Iter) Flatten() IterTrait {
	return i
}

// Fold folds all values based on the fold function that are Operable and returns a single Data value
func (i *Iter) Fold(init collections.OperableData, f func(result collections.OperableData, next collections.OperableData) collections.Data) collections.Data {
	result := init
	for ii := 0; ii < len(i.values); ii++ {
		if op, ok := i.values[ii].(collections.OperableData); ok {
			result = f(result, op).(collections.OperableData)
		}
	}
	return result.(collections.Data)
}

// Ge determines if this iterator is greater than the other iterator
func (i *Iter) Ge(other IterTrait) bool {
	if i.Len() > other.Len() {
		return true
	}
	if i.Eq(other) {
		return true
	}
	return i.Gt(other)
}

// Gt determines if this iterator is greater than or equal to the other iterator
func (i *Iter) Gt(other IterTrait) bool {
	if i.Len() > other.Len() {
		return true
	}

	if i.Len() < other.Len() {
		return false
	}

	selfCh := i.Iterate()
	otherCh := other.Iterate()
	for {
		sV, selfOK := <-selfCh
		otherV, otherOK := <-otherCh

		if !selfOK && !otherOK {
			break
		}

		if sV.Greater(otherV) {
			return true
		}

	}
	return false
}

// Inspect allows debug lines to be called in-between chained events
func (i *Iter) Inspect(f func(d collections.Data)) IterTrait {
	for j := range i.values {
		f(i.values[j])
	}
	return i
}

// Last returns the next element in the collection in the reverse order and its enumerated position and moves the cursor forward
// If Last is called Next cannot be called for this copy of Iter, returns -1 as an index when the iterator is fully consumed
func (i *Iter) Last() (int, collections.Data) {
	if i.direction == _notSet {
		i.direction = _backwards
		i.values = reverse(i.values)
	} else if i.direction == _forward {
		panic("Next was called for this iterator, therefore Last cannot be called, create a new iterator to use Last")
	}

	if !i.shouldCycle {
		if i.currentIdx >= len(i.values) {
			return -1, nil
		}
	}

	curIdx, data := i.currentIdx, i.values[i.currentIdx%len(i.values)]
	i.currentIdx++
	return curIdx, data
}

// Le determines if this iterator is less than the other iterator
func (i *Iter) Le(other IterTrait) bool {
	if i.Len() < other.Len() {
		return true
	}
	if i.Eq(other) {
		return true
	}
	return i.Lt(other)
}

// Len returns the current length of the underline data slice
func (i *Iter) Len() int {
	return len(i.values)
}

// Lt determines if this iterator is less than or equal to the other iterator
func (i *Iter) Lt(other IterTrait) bool {
	if i.Len() < other.Len() {
		return true
	}

	if i.Len() > other.Len() {
		return false
	}

	selfCh := i.Iterate()
	otherCh := other.Iterate()
	for {
		sV, selfOK := <-selfCh
		otherV, otherOK := <-otherCh

		if !selfOK && !otherOK {
			break
		}

		if sV.Less(otherV) {
			return true
		}

	}
	return false
}

// Map takes a closure and creates an iterator which calls that closure on each element
func (i *Iter) Map(f func(d collections.Data) collections.Data) IterTrait {
	for ii := 0; ii < len(i.values); ii++ {
		i.values[ii] = f(i.values[ii])
	}
	return i
}

// Max returns the max value in the data collection
func (i *Iter) Max() collections.Data {
	if len(i.values) == 0 {
		return nil
	}
	max := i.values[0]
	for ii := 1; ii < len(i.values); ii++ {
		if i.values[ii].Greater(max) {
			max = i.values[ii]
		}
	}
	return max
}

// Min returns the min value in the data collection
func (i *Iter) Min() collections.Data {
	if len(i.values) == 0 {
		return nil
	}
	min := i.values[0]
	for ii := 1; ii < len(i.values); ii++ {
		if i.values[ii].Less(min) {
			min = i.values[ii]
		}
	}
	return min
}

// Ne determines if this iterator is different from the other iterator
func (i *Iter) Ne(other IterTrait) bool {
	return !i.Eq(other)
}

func (i *Iter) Nth(n int) collections.Data {
	if n > len(i.values) {
		return nil
	}
	data := i.values[n]
	i.values = i.values[n:]
	return data
}

// Partition consumes an iterator, creating two collections from it.
func (i *Iter) Partition(f func(d collections.Data) bool) ([]collections.Data, []collections.Data) {
	var a, b []collections.Data
	for ii := 0; ii < len(i.values); ii++ {
		if f(i.values[ii]) {
			a = append(a, i.values[ii])
			continue
		}
		b = append(b, i.values[ii])
	}
	i.currentIdx = len(i.values)
	return a, b
}

// Peak returns the next value in the collection without consuming the iterator
func (i *Iter) Peak() collections.Data {
	if i.currentIdx < len(i.values) {
		return i.values[i.currentIdx]
	}
	return nil
}

// Position returns the index value of the first matching element as defined the function
func (i *Iter) Position(f func(d collections.Data) bool) int {
	for pos, data := range i.values {
		if f(data) {
			return pos
		}
	}
	return -1
}

// Product iterates over the entire iterator, multiplying all the elements
func (i *Iter) Product() collections.Data {
	if len(i.values) == 0 {
		return collections.IntValue(0)
	}

	var result collections.Data
	if op, ok := i.values[0].(collections.OperableData); ok {
		result = op.(collections.Data)
	}

	if result == nil {
		return collections.IntValue(0)
	}

	for ii := 1; ii < len(i.values); ii++ {
		if op, ok := i.values[ii].(collections.OperableData); ok {
			result = result.(collections.OperableData).Mul(op).(collections.Data)
		}
	}
	return result
}

// Reduce reduces all values based on the fold function that are Operable and returns a single Data value
func (i *Iter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data {
	result := i.values[0].(collections.Data)
	for ii := 0; ii < len(i.values); ii++ {
		result = f(result, i.values[ii])
	}
	return result
}

// Take creates an iterator that yields the first n elements, or fewer if the underlying iterator ends sooner.
func (i *Iter) Take(n int) {
	i.values = i.values[:n]
}

// Sum sums the elements of an iterator.
func (i *Iter) Sum() collections.Data {
	if len(i.values) == 0 {
		return collections.IntValue(0)
	}

	var result collections.Data
	if op, ok := i.values[0].(collections.OperableData); ok {
		result = op.(collections.Data)
	}

	if result == nil {
		return collections.IntValue(0)
	}

	for ii := 1; ii < len(i.values); ii++ {
		if op, ok := i.values[ii].(collections.OperableData); ok {
			result = result.(collections.OperableData).Add(op).(collections.Data)
		}
	}
	return result
}

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func reverse(s []collections.Data) []collections.Data {
	a := make([]collections.Data, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}
