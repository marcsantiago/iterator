package iterator

import (
	"reflect"
	"sort"

	"github.com/marcsantiago/collections"
)

type internalMap collections.Map

type MapIter struct {
	ch <-chan collections.Element
	internalMap
}

var _ IterTraitMap = (*MapIter)(nil)

// NewMapIter creates a new instance of MapIter which implements IterTraitMap
func NewMapIter() *MapIter {
	return &MapIter{
		nil,
		collections.NewGenericMap(),
	}
}

// NewMapIterFromElements creates a new instance of MapIter from a collection of []collections.Element
func NewMapIterFromElements(elems []collections.Element) *MapIter {
	m := NewMapIter()
	for _, e := range elems {
		m.Set(e.Key, e.Value)
	}
	return &MapIter{
		nil,
		m,
	}
}

// NewMapIterFromMap creates a new instance of MapIter from a collections.Map
func NewMapIterFromMap(m collections.Map) *MapIter {
	return &MapIter{
		nil,
		m,
	}
}

// All returns true if all the values match the predicate
func (m *MapIter) All(f func(d collections.Data) bool) bool {
	for _, elem := range m.Items() {
		if !f(elem.Value) {
			return false
		}
	}
	return true
}

// Any returns true if any the value matches the predicate
func (m *MapIter) Any(f func(d collections.Data) bool) bool {
	for _, elem := range m.Items() {
		if f(elem.Value) {
			return true
		}
	}
	return false
}

// Count consumes the iterator and returns the count of non nil items
func (m *MapIter) Count() int {
	if m.Len() == 0 {
		return 0
	}
	var count int
	for _, elem := range m.Items() {
		if elem.Value != nil {
			count++
		}
		m.Delete(elem.Key)
	}
	return count
}

// Max returns the max value in the data collection
func (m *MapIter) Max() collections.Data {
	if m.internalMap == nil || m.Len() == 0 {
		return nil
	}

	isFirstElement := true
	var max collections.Data
	for _, elem := range m.Items() {

		if isFirstElement {
			isFirstElement = false
			max = elem.Value
			continue
		}

		if max == nil {
			continue
		}

		if elem.Value.Greater(max) {
			max = elem.Value
		}
	}
	return max
}

// Min returns the min value in the data collection
func (m *MapIter) Min() collections.Data {
	if m.internalMap == nil || m.Len() == 0 {
		return nil
	}

	isFirstElement := true
	var min collections.Data
	for _, elem := range m.Items() {
		if elem.Value == nil {
			continue
		}

		if isFirstElement {
			isFirstElement = false
			min = elem.Value
			continue
		}

		if min == nil {
			continue
		}

		if elem.Value.Less(min) {
			min = elem.Value
		}
	}
	return min
}

// Chain Takes two iterators and creates a new iterator over both in sequence
func (m *MapIter) Chain(other IterTraitMap) IterTraitMap {
	for elem := range other.Iterate() {
		m.Set(elem.Key, elem.Value)
	}
	return m
}

// Collect converts returns the concrete []collections.Element in the iterator
func (m *MapIter) Collect() []collections.Element {
	return m.Items()
}

// CollectKeys returns the concrete keys as []collections.Data
func (m *MapIter) CollectKeys() []collections.Data {
	data := make([]collections.Data, 0, m.Len())
	for _, elem := range m.Items() {
		data = append(data, elem.Key)
	}
	return data
}

// CollectValues returns the concrete values as []collections.Data
func (m *MapIter) CollectValues() []collections.Data {
	data := make([]collections.Data, 0, m.Len())
	for _, elem := range m.Items() {
		data = append(data, elem.Value)
	}
	return data
}

// KeysToIterSlice returns the keys as an IterTraitSlice
func (m *MapIter) KeysToIterSlice() IterTraitSlice {
	data := make([]collections.Data, 0, m.Len())
	for _, elem := range m.Items() {
		data = append(data, elem.Key)
	}
	return IntoIter(data)
}

// ValuesToIterSlice returns the keys as an IterTraitSlice
func (m *MapIter) ValuesToIterSlice() IterTraitSlice {
	data := make([]collections.Data, 0, m.Len())
	for _, elem := range m.Items() {
		data = append(data, elem.Value)
	}
	return IntoIter(data)
}

// Eq determines if this iterator is the same as the other iterator
func (m *MapIter) Eq(other IterTraitMap) bool {
	if m.Len() != other.Len() {
		return false
	}

	mc := m.Collect()
	otherC := other.Collect()
	sort.Sort(elementsByKeyData(mc))
	sort.Sort(elementsByKeyData(otherC))
	return reflect.DeepEqual(mc, otherC)
}

// Filter removes all values by which the comparison function returns true for values in the map
func (m *MapIter) Filter(f func(d collections.Data) bool) IterTraitMap {
	for _, elem := range m.Items() {
		if !f(elem.Value) {
			m.Delete(elem.Key)
		}
	}
	return m
}

// FilterKeys removes all values by which the comparison function returns true for keys in the map
func (m *MapIter) FilterKeys(f func(d collections.Data) bool) IterTraitMap {
	for _, elem := range m.Items() {
		if !f(elem.Key) {
			m.Delete(elem.Key)
		}
	}
	return m
}

// Find returns the first value that matches the find function targeting keys, else it returns nil
func (m *MapIter) Find(f func(d collections.Data) bool) (item collections.Element, ok bool) {
	for _, elem := range m.Items() {
		if f(elem.Key) {
			return elem, true
		}
	}
	return item, false
}

// FindByValue returns the first value that matches the find function targeting values, else it returns nil
func (m *MapIter) FindByValue(f func(d collections.Data) bool) (item collections.Element, ok bool) {
	for _, elem := range m.Items() {
		if f(elem.Value) {
			return elem, true
		}
	}
	return item, false
}

// Ge determines if this iterator is greater than the other iterator
func (m *MapIter) Ge(other IterTraitMap) bool {
	if m.Len() > other.Len() {
		return true
	}
	if m.Eq(other) {
		return true
	}
	return m.Gt(other)
}

// Gt determines if this iterator is greater than or equal to the other iterator
func (m *MapIter) Gt(other IterTraitMap) bool {
	if m.Len() > other.Len() {
		return true
	}

	if m.Len() < other.Len() {
		return false
	}

	if m.Eq(other) {
		return false
	}

	for _, elem := range m.Items() {
		// if the key doesn't exist in the other collection we ignore it, we only
		// want to compare values that are apples to apples
		if data, ok := other.Get(elem.Key); ok {
			if elem.Value.Less(data) {
				return false
			}
		}
	}
	return true
}

// Inspect allows debug lines to be called in-between chained events
func (m *MapIter) Inspect(f func(d collections.Element)) IterTraitMap {
	for _, elem := range m.Items() {
		f(elem)
	}
	return m
}

// Le determines if this iterator is less than the other iterator
func (m *MapIter) Le(other IterTraitMap) bool {
	if m.Len() < other.Len() {
		return true
	}
	if m.Eq(other) {
		return true
	}
	return m.Lt(other)
}

// Lt determines if this iterator is less than or equal to the other iterator
func (m *MapIter) Lt(other IterTraitMap) bool {
	if m.Len() > other.Len() {
		return false
	}
	if m.Len() < other.Len() {
		return true
	}

	if m.Eq(other) {
		return false
	}

	for _, elem := range m.Items() {
		// if the key doesn't exist in the other collection we ignore it, we only
		// want to compare values that are apples to apples
		if data, ok := other.Get(elem.Key); ok {
			if elem.Value.Greater(data) {
				return false
			}
		}

	}
	return true
}

// Map takes a closure and creates an iterator which calls that closure on each element
func (m *MapIter) Map(f func(d collections.Data) collections.Data) IterTraitMap {
	for _, elem := range m.Items() {
		value, _ := m.Get(elem.Key)
		value = f(value)
		m.Set(elem.Key, value)
	}
	return m
}

// Ne determines if this iterator is different from the other iterator
func (m *MapIter) Ne(other IterTraitMap) bool {
	return !m.Eq(other)
}

// Next returns the next element in the collection and moves the cursor forward returns false when the iterator is fully consumed
func (m *MapIter) Next() (collections.Element, bool) {
	var elem collections.Element
	if m.Len() == 0 {
		return elem, false
	}

	if m.ch == nil {
		m.ch = m.Iterate()
	}

	if got, ok := <-m.ch; ok {
		return got, ok
	}
	return elem, false
}

// Reduce reduces all values based on the fold function that are Operable and returns a single Data value
func (m *MapIter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data {
	if m.Len() == 0 {
		return nil
	}

	items := m.Items()
	result := items[0].Value
	for i := 0; i < len(items); i++ {
		result = f(result, items[i].Value)
	}
	return result
}

type elementsByKeyData []collections.Element

func (e elementsByKeyData) Len() int           { return len(e) }
func (e elementsByKeyData) Less(i, j int) bool { return e[i].Key.Less(e[j].Key) }
func (e elementsByKeyData) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
