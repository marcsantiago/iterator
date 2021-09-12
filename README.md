

# iterator
`import "/Users/marcsantiago/go/src/github.com/marcsantiago/iterator"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type BaseIterTrait](#BaseIterTrait)
* [type Iter](#Iter)
  * [func IntoIter(data []collections.Data) *Iter](#IntoIter)
  * [func NewIter() *Iter](#NewIter)
  * [func (i *Iter) All(f func(d collections.Data) bool) bool](#Iter.All)
  * [func (i *Iter) Any(f func(d collections.Data) bool) bool](#Iter.Any)
  * [func (i *Iter) Append(data collections.Data)](#Iter.Append)
  * [func (i *Iter) Chain(other IterTraitSlice) IterTraitSlice](#Iter.Chain)
  * [func (i *Iter) Collect() []collections.Data](#Iter.Collect)
  * [func (i *Iter) CollectBools() []bool](#Iter.CollectBools)
  * [func (i *Iter) CollectFloat32s() []float32](#Iter.CollectFloat32s)
  * [func (i *Iter) CollectFloat64s() []float64](#Iter.CollectFloat64s)
  * [func (i *Iter) CollectInt32s() []int32](#Iter.CollectInt32s)
  * [func (i *Iter) CollectInt64s() []int64](#Iter.CollectInt64s)
  * [func (i *Iter) CollectInts() []int](#Iter.CollectInts)
  * [func (i *Iter) CollectStrings() []string](#Iter.CollectStrings)
  * [func (i *Iter) Count() int](#Iter.Count)
  * [func (i *Iter) Cycle()](#Iter.Cycle)
  * [func (i *Iter) Eq(other IterTraitSlice) bool](#Iter.Eq)
  * [func (i *Iter) Filter(f func(d collections.Data) bool) IterTraitSlice](#Iter.Filter)
  * [func (i *Iter) Find(f func(d collections.Data) bool) collections.Data](#Iter.Find)
  * [func (i *Iter) Fold(init collections.OperableData, f func(result collections.OperableData, next collections.OperableData) collections.Data) collections.Data](#Iter.Fold)
  * [func (i *Iter) FoldIntoMap(init IterTraitMap, f func(m IterTraitMap, key collections.Data) IterTraitMap) IterTraitMap](#Iter.FoldIntoMap)
  * [func (i *Iter) Ge(other IterTraitSlice) bool](#Iter.Ge)
  * [func (i *Iter) Gt(other IterTraitSlice) bool](#Iter.Gt)
  * [func (i *Iter) Inspect(f func(d collections.Data)) IterTraitSlice](#Iter.Inspect)
  * [func (i *Iter) Iterate() &lt;-chan collections.Data](#Iter.Iterate)
  * [func (i *Iter) Last() (int, collections.Data)](#Iter.Last)
  * [func (i *Iter) Le(other IterTraitSlice) bool](#Iter.Le)
  * [func (i *Iter) Len() int](#Iter.Len)
  * [func (i *Iter) Lt(other IterTraitSlice) bool](#Iter.Lt)
  * [func (i *Iter) Map(f func(d collections.Data) collections.Data) IterTraitSlice](#Iter.Map)
  * [func (i *Iter) Max() collections.Data](#Iter.Max)
  * [func (i *Iter) Min() collections.Data](#Iter.Min)
  * [func (i *Iter) Ne(other IterTraitSlice) bool](#Iter.Ne)
  * [func (i *Iter) Next() (int, collections.Data)](#Iter.Next)
  * [func (i *Iter) Nth(n int) collections.Data](#Iter.Nth)
  * [func (i *Iter) Partition(f func(d collections.Data) bool) ([]collections.Data, []collections.Data)](#Iter.Partition)
  * [func (i *Iter) Peak() collections.Data](#Iter.Peak)
  * [func (i *Iter) Position(f func(d collections.Data) bool) int](#Iter.Position)
  * [func (i *Iter) Product() collections.Data](#Iter.Product)
  * [func (i *Iter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data](#Iter.Reduce)
  * [func (i *Iter) String() string](#Iter.String)
  * [func (i *Iter) Sum() collections.Data](#Iter.Sum)
  * [func (i *Iter) Take(n int)](#Iter.Take)
  * [func (i *Iter) Zip(other []collections.Data) []collections.Element](#Iter.Zip)
  * [func (i *Iter) ZipIntoMap(other []collections.Data) IterTraitMap](#Iter.ZipIntoMap)
* [type IterTraitMap](#IterTraitMap)
* [type IterTraitSlice](#IterTraitSlice)
* [type MapIter](#MapIter)
  * [func NewMapIter() *MapIter](#NewMapIter)
  * [func NewMapIterFromElements(elems []collections.Element) *MapIter](#NewMapIterFromElements)
  * [func NewMapIterFromMap(m collections.Map) *MapIter](#NewMapIterFromMap)
  * [func (m *MapIter) All(f func(d collections.Data) bool) bool](#MapIter.All)
  * [func (m *MapIter) Any(f func(d collections.Data) bool) bool](#MapIter.Any)
  * [func (m *MapIter) Chain(other IterTraitMap) IterTraitMap](#MapIter.Chain)
  * [func (m *MapIter) Collect() []collections.Element](#MapIter.Collect)
  * [func (m *MapIter) CollectKeys() []collections.Data](#MapIter.CollectKeys)
  * [func (m *MapIter) CollectValues() []collections.Data](#MapIter.CollectValues)
  * [func (m *MapIter) Count() int](#MapIter.Count)
  * [func (m *MapIter) Eq(other IterTraitMap) bool](#MapIter.Eq)
  * [func (m *MapIter) Filter(f func(d collections.Data) bool) IterTraitMap](#MapIter.Filter)
  * [func (m *MapIter) FilterKeys(f func(d collections.Data) bool) IterTraitMap](#MapIter.FilterKeys)
  * [func (m *MapIter) Find(f func(d collections.Data) bool) (item collections.Element, ok bool)](#MapIter.Find)
  * [func (m *MapIter) FindByValue(f func(d collections.Data) bool) (item collections.Element, ok bool)](#MapIter.FindByValue)
  * [func (m *MapIter) Ge(other IterTraitMap) bool](#MapIter.Ge)
  * [func (m *MapIter) Gt(other IterTraitMap) bool](#MapIter.Gt)
  * [func (m *MapIter) Inspect(f func(d collections.Element)) IterTraitMap](#MapIter.Inspect)
  * [func (m *MapIter) KeysToIterSlice() IterTraitSlice](#MapIter.KeysToIterSlice)
  * [func (m *MapIter) Le(other IterTraitMap) bool](#MapIter.Le)
  * [func (m *MapIter) Lt(other IterTraitMap) bool](#MapIter.Lt)
  * [func (m *MapIter) Map(f func(d collections.Data) collections.Data) IterTraitMap](#MapIter.Map)
  * [func (m *MapIter) Max() collections.Data](#MapIter.Max)
  * [func (m *MapIter) Min() collections.Data](#MapIter.Min)
  * [func (m *MapIter) Ne(other IterTraitMap) bool](#MapIter.Ne)
  * [func (m *MapIter) Next() (collections.Element, bool)](#MapIter.Next)
  * [func (m *MapIter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data](#MapIter.Reduce)
  * [func (m *MapIter) ValuesToIterSlice() IterTraitSlice](#MapIter.ValuesToIterSlice)


#### <a name="pkg-files">Package files</a>
[iterator_iface.go](/src/target/iterator_iface.go) [map_iter.go](/src/target/map_iter.go) [slice_iter.go](/src/target/slice_iter.go) 






## <a name="BaseIterTrait">type</a> [BaseIterTrait](/src/target/iterator_iface.go?s=127:317#L8)
``` go
type BaseIterTrait interface {
    All(f func(d collections.Data) bool) bool
    Any(f func(d collections.Data) bool) bool
    Count() int
    Max() collections.Data
    Min() collections.Data
    Len() int
}
```









## <a name="Iter">type</a> [Iter](/src/target/slice_iter.go?s=160:271#L17)
``` go
type Iter struct {
    // contains filtered or unexported fields
}
```






### <a name="IntoIter">func</a> [IntoIter](/src/target/slice_iter.go?s=352:396#L30)
``` go
func IntoIter(data []collections.Data) *Iter
```

### <a name="NewIter">func</a> [NewIter](/src/target/slice_iter.go?s=310:330#L26)
``` go
func NewIter() *Iter
```




### <a name="Iter.All">func</a> (\*Iter) [All](/src/target/slice_iter.go?s=1385:1441#L62)
``` go
func (i *Iter) All(f func(d collections.Data) bool) bool
```
All returns true if all the values match the predicate




### <a name="Iter.Any">func</a> (\*Iter) [Any](/src/target/slice_iter.go?s=1596:1652#L72)
``` go
func (i *Iter) Any(f func(d collections.Data) bool) bool
```
Any returns true if any the value matches the predicate




### <a name="Iter.Append">func</a> (\*Iter) [Append](/src/target/slice_iter.go?s=559:603#L37)
``` go
func (i *Iter) Append(data collections.Data)
```
Append not very Rust like, but adding a convenience to append into internal data without needing to pass in a full copy




### <a name="Iter.Chain">func</a> (\*Iter) [Chain](/src/target/slice_iter.go?s=1825:1882#L82)
``` go
func (i *Iter) Chain(other IterTraitSlice) IterTraitSlice
```
Chain Takes two iterators and creates a new iterator over both in sequence




### <a name="Iter.Collect">func</a> (\*Iter) [Collect](/src/target/slice_iter.go?s=2215:2258#L98)
``` go
func (i *Iter) Collect() []collections.Data
```
Collect converts returns the concrete []collections.Data in the iterator




### <a name="Iter.CollectBools">func</a> (\*Iter) [CollectBools](/src/target/slice_iter.go?s=3814:3850#L157)
``` go
func (i *Iter) CollectBools() []bool
```
CollectBools converts []collections.Data into the concrete type []bools




### <a name="Iter.CollectFloat32s">func</a> (\*Iter) [CollectFloat32s](/src/target/slice_iter.go?s=3069:3111#L130)
``` go
func (i *Iter) CollectFloat32s() []float32
```
CollectFloat32s converts []collections.Data into the concrete type []float32




### <a name="Iter.CollectFloat64s">func</a> (\*Iter) [CollectFloat64s](/src/target/slice_iter.go?s=3321:3363#L139)
``` go
func (i *Iter) CollectFloat64s() []float64
```
CollectFloat64s converts []collections.Data into the concrete type []float64




### <a name="Iter.CollectInt32s">func</a> (\*Iter) [CollectInt32s](/src/target/slice_iter.go?s=2585:2623#L112)
``` go
func (i *Iter) CollectInt32s() []int32
```
CollectInt32s converts []collections.Data into the concrete type []int32




### <a name="Iter.CollectInt64s">func</a> (\*Iter) [CollectInt64s](/src/target/slice_iter.go?s=2825:2863#L121)
``` go
func (i *Iter) CollectInt64s() []int64
```
CollectInt64s converts []collections.Data into the concrete type []int64




### <a name="Iter.CollectInts">func</a> (\*Iter) [CollectInts](/src/target/slice_iter.go?s=2353:2387#L103)
``` go
func (i *Iter) CollectInts() []int
```
CollectInts converts []collections.Data into the concrete type []int




### <a name="Iter.CollectStrings">func</a> (\*Iter) [CollectStrings](/src/target/slice_iter.go?s=3571:3611#L148)
``` go
func (i *Iter) CollectStrings() []string
```
CollectStrings converts []collections.Data into the concrete type []string




### <a name="Iter.Count">func</a> (\*Iter) [Count](/src/target/slice_iter.go?s=4044:4070#L166)
``` go
func (i *Iter) Count() int
```
Count consumes the iterator and returns the count of non nil items




### <a name="Iter.Cycle">func</a> (\*Iter) [Cycle](/src/target/slice_iter.go?s=4349:4371#L181)
``` go
func (i *Iter) Cycle()
```
Cycle enables Next or Last to iterate forever cycling through the items in order seen




### <a name="Iter.Eq">func</a> (\*Iter) [Eq](/src/target/slice_iter.go?s=4467:4511#L186)
``` go
func (i *Iter) Eq(other IterTraitSlice) bool
```
Eq determines if this iterator is the same as the other iterator




### <a name="Iter.Filter">func</a> (\*Iter) [Filter](/src/target/slice_iter.go?s=5102:5171#L221)
``` go
func (i *Iter) Filter(f func(d collections.Data) bool) IterTraitSlice
```
Filter removes all values by which the comparison function returns true




### <a name="Iter.Find">func</a> (\*Iter) [Find](/src/target/slice_iter.go?s=5429:5498#L236)
``` go
func (i *Iter) Find(f func(d collections.Data) bool) collections.Data
```
Find returns the first value that matches the find function, else it returns nil




### <a name="Iter.Fold">func</a> (\*Iter) [Fold](/src/target/slice_iter.go?s=5712:5868#L246)
``` go
func (i *Iter) Fold(init collections.OperableData, f func(result collections.OperableData, next collections.OperableData) collections.Data) collections.Data
```
Fold folds all values based on the fold function that are Operable and returns a single Data value




### <a name="Iter.FoldIntoMap">func</a> (\*Iter) [FoldIntoMap](/src/target/slice_iter.go?s=6145:6262#L257)
``` go
func (i *Iter) FoldIntoMap(init IterTraitMap, f func(m IterTraitMap, key collections.Data) IterTraitMap) IterTraitMap
```
FoldIntoMap folds all the values into a collections.Map




### <a name="Iter.Ge">func</a> (\*Iter) [Ge](/src/target/slice_iter.go?s=6422:6466#L266)
``` go
func (i *Iter) Ge(other IterTraitSlice) bool
```
Ge determines if this iterator is greater than the other iterator




### <a name="Iter.Gt">func</a> (\*Iter) [Gt](/src/target/slice_iter.go?s=6653:6697#L277)
``` go
func (i *Iter) Gt(other IterTraitSlice) bool
```
Gt determines if this iterator is greater than or equal to the other iterator




### <a name="Iter.Inspect">func</a> (\*Iter) [Inspect](/src/target/slice_iter.go?s=7084:7149#L305)
``` go
func (i *Iter) Inspect(f func(d collections.Data)) IterTraitSlice
```
Inspect allows debug lines to be called in-between chained events




### <a name="Iter.Iterate">func</a> (\*Iter) [Iterate](/src/target/slice_iter.go?s=4840:4888#L209)
``` go
func (i *Iter) Iterate() <-chan collections.Data
```
Iterate returns a channel of values that can be ranged over




### <a name="Iter.Last">func</a> (\*Iter) [Last](/src/target/slice_iter.go?s=7466:7511#L314)
``` go
func (i *Iter) Last() (int, collections.Data)
```
Last returns the next element in the collection in the reverse order and its enumerated position and moves the cursor forward
If Last is called Next cannot be called for this copy of Iter, returns -1 as an index when the iterator is fully consumed




### <a name="Iter.Le">func</a> (\*Iter) [Le](/src/target/slice_iter.go?s=8013:8057#L334)
``` go
func (i *Iter) Le(other IterTraitSlice) bool
```
Le determines if this iterator is less than the other iterator




### <a name="Iter.Len">func</a> (\*Iter) [Len](/src/target/slice_iter.go?s=8225:8249#L345)
``` go
func (i *Iter) Len() int
```
Len returns the current length of the underline data slice




### <a name="Iter.Lt">func</a> (\*Iter) [Lt](/src/target/slice_iter.go?s=8355:8399#L350)
``` go
func (i *Iter) Lt(other IterTraitSlice) bool
```
Lt determines if this iterator is less than or equal to the other iterator




### <a name="Iter.Map">func</a> (\*Iter) [Map](/src/target/slice_iter.go?s=8802:8880#L378)
``` go
func (i *Iter) Map(f func(d collections.Data) collections.Data) IterTraitSlice
```
Map takes a closure and creates an iterator which calls that closure on each element




### <a name="Iter.Max">func</a> (\*Iter) [Max](/src/target/slice_iter.go?s=9025:9062#L386)
``` go
func (i *Iter) Max() collections.Data
```
Max returns the max value in the data collection




### <a name="Iter.Min">func</a> (\*Iter) [Min](/src/target/slice_iter.go?s=9296:9333#L400)
``` go
func (i *Iter) Min() collections.Data
```
Min returns the min value in the data collection




### <a name="Iter.Ne">func</a> (\*Iter) [Ne](/src/target/slice_iter.go?s=9583:9627#L414)
``` go
func (i *Iter) Ne(other IterTraitSlice) bool
```
Ne determines if this iterator is different from the other iterator




### <a name="Iter.Next">func</a> (\*Iter) [Next](/src/target/slice_iter.go?s=877:922#L43)
``` go
func (i *Iter) Next() (int, collections.Data)
```
Next returns the next element in the collection and its enumerated position and moves the cursor forward
If Next is called Last cannot be called for this copy of Iter, returns -1 as an index when the iterator is fully consumed




### <a name="Iter.Nth">func</a> (\*Iter) [Nth](/src/target/slice_iter.go?s=9665:9707#L419)
``` go
func (i *Iter) Nth(n int) collections.Data
```
Nth ...




### <a name="Iter.Partition">func</a> (\*Iter) [Partition](/src/target/slice_iter.go?s=9881:9979#L429)
``` go
func (i *Iter) Partition(f func(d collections.Data) bool) ([]collections.Data, []collections.Data)
```
Partition consumes an iterator, creating two collections from it.




### <a name="Iter.Peak">func</a> (\*Iter) [Peak](/src/target/slice_iter.go?s=10281:10319#L443)
``` go
func (i *Iter) Peak() collections.Data
```
Peak returns the next value in the collection without consuming the iterator




### <a name="Iter.Position">func</a> (\*Iter) [Position](/src/target/slice_iter.go?s=10497:10557#L451)
``` go
func (i *Iter) Position(f func(d collections.Data) bool) int
```
Position returns the index value of the first matching element as defined the function




### <a name="Iter.Product">func</a> (\*Iter) [Product](/src/target/slice_iter.go?s=10720:10761#L461)
``` go
func (i *Iter) Product() collections.Data
```
Product iterates over the entire iterator, multiplying all the elements




### <a name="Iter.Reduce">func</a> (\*Iter) [Reduce](/src/target/slice_iter.go?s=11314:11400#L484)
``` go
func (i *Iter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data
```
Reduce reduces all values based on the fold function that are Operable and returns a single Data value




### <a name="Iter.String">func</a> (\*Iter) [String](/src/target/slice_iter.go?s=11796:11826#L502)
``` go
func (i *Iter) String() string
```
String to satisfy the stringer interface




### <a name="Iter.Sum">func</a> (\*Iter) [Sum](/src/target/slice_iter.go?s=11910:11947#L507)
``` go
func (i *Iter) Sum() collections.Data
```
Sum sums the elements of an iterator.




### <a name="Iter.Take">func</a> (\*Iter) [Take](/src/target/slice_iter.go?s=11695:11721#L497)
``` go
func (i *Iter) Take(n int)
```
Take creates an iterator that yields the first n elements, or fewer if the underlying iterator ends sooner.




### <a name="Iter.Zip">func</a> (\*Iter) [Zip](/src/target/slice_iter.go?s=12581:12647#L531)
``` go
func (i *Iter) Zip(other []collections.Data) []collections.Element
```
Zip combines the value for from Iter as the key in Element and other's Data as the value
the Zip is only as long as the min length of collections and returns a []collection.Element




### <a name="Iter.ZipIntoMap">func</a> (\*Iter) [ZipIntoMap](/src/target/slice_iter.go?s=13099:13163#L547)
``` go
func (i *Iter) ZipIntoMap(other []collections.Data) IterTraitMap
```
ZipIntoMap combines the value for from Iter as the key in Element and other's Data as the value
the Zip is only as long as the min length of collections and returns an implementation of IterTraitMap




## <a name="IterTraitMap">type</a> [IterTraitMap](/src/target/iterator_iface.go?s=1873:2825#L57)
``` go
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
```









## <a name="IterTraitSlice">type</a> [IterTraitSlice](/src/target/iterator_iface.go?s=319:1871#L17)
``` go
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
```









## <a name="MapIter">type</a> [MapIter](/src/target/map_iter.go?s=123:190#L12)
``` go
type MapIter struct {
    // contains filtered or unexported fields
}
```






### <a name="NewMapIter">func</a> [NewMapIter](/src/target/map_iter.go?s=308:334#L20)
``` go
func NewMapIter() *MapIter
```
NewMapIter creates a new instance of MapIter which implements IterTraitMap


### <a name="NewMapIterFromElements">func</a> [NewMapIterFromElements](/src/target/map_iter.go?s=502:567#L28)
``` go
func NewMapIterFromElements(elems []collections.Element) *MapIter
```
NewMapIterFromElements creates a new instance of MapIter from a collection of []collections.Element


### <a name="NewMapIterFromMap">func</a> [NewMapIterFromMap](/src/target/map_iter.go?s=757:807#L40)
``` go
func NewMapIterFromMap(m collections.Map) *MapIter
```
NewMapIterFromMap creates a new instance of MapIter from a collections.Map





### <a name="MapIter.All">func</a> (\*MapIter) [All](/src/target/map_iter.go?s=904:963#L48)
``` go
func (m *MapIter) All(f func(d collections.Data) bool) bool
```
All returns true if all the values match the predicate




### <a name="MapIter.Any">func</a> (\*MapIter) [Any](/src/target/map_iter.go?s=1120:1179#L58)
``` go
func (m *MapIter) Any(f func(d collections.Data) bool) bool
```
Any returns true if any the value matches the predicate




### <a name="MapIter.Chain">func</a> (\*MapIter) [Chain](/src/target/map_iter.go?s=2366:2422#L129)
``` go
func (m *MapIter) Chain(other IterTraitMap) IterTraitMap
```
Chain Takes two iterators and creates a new iterator over both in sequence




### <a name="MapIter.Collect">func</a> (\*MapIter) [Collect](/src/target/map_iter.go?s=2587:2636#L137)
``` go
func (m *MapIter) Collect() []collections.Element
```
Collect converts returns the concrete []collections.Element in the iterator




### <a name="MapIter.CollectKeys">func</a> (\*MapIter) [CollectKeys](/src/target/map_iter.go?s=2723:2773#L142)
``` go
func (m *MapIter) CollectKeys() []collections.Data
```
CollectKeys returns the concrete keys as []collections.Data




### <a name="MapIter.CollectValues">func</a> (\*MapIter) [CollectValues](/src/target/map_iter.go?s=2974:3026#L151)
``` go
func (m *MapIter) CollectValues() []collections.Data
```
CollectValues returns the concrete values as []collections.Data




### <a name="MapIter.Count">func</a> (\*MapIter) [Count](/src/target/map_iter.go?s=1346:1375#L68)
``` go
func (m *MapIter) Count() int
```
Count consumes the iterator and returns the count of non nil items




### <a name="MapIter.Eq">func</a> (\*MapIter) [Eq](/src/target/map_iter.go?s=3738:3783#L178)
``` go
func (m *MapIter) Eq(other IterTraitMap) bool
```
Eq determines if this iterator is the same as the other iterator




### <a name="MapIter.Filter">func</a> (\*MapIter) [Filter](/src/target/map_iter.go?s=4090:4160#L191)
``` go
func (m *MapIter) Filter(f func(d collections.Data) bool) IterTraitMap
```
Filter removes all values by which the comparison function returns true for values in the map




### <a name="MapIter.FilterKeys">func</a> (\*MapIter) [FilterKeys](/src/target/map_iter.go?s=4360:4434#L201)
``` go
func (m *MapIter) FilterKeys(f func(d collections.Data) bool) IterTraitMap
```
FilterKeys removes all values by which the comparison function returns true for keys in the map




### <a name="MapIter.Find">func</a> (\*MapIter) [Find](/src/target/map_iter.go?s=4632:4723#L211)
``` go
func (m *MapIter) Find(f func(d collections.Data) bool) (item collections.Element, ok bool)
```
Find returns the first value that matches the find function targeting keys, else it returns nil




### <a name="MapIter.FindByValue">func</a> (\*MapIter) [FindByValue](/src/target/map_iter.go?s=4938:5036#L221)
``` go
func (m *MapIter) FindByValue(f func(d collections.Data) bool) (item collections.Element, ok bool)
```
FindByValue returns the first value that matches the find function targeting values, else it returns nil




### <a name="MapIter.Ge">func</a> (\*MapIter) [Ge](/src/target/map_iter.go?s=5214:5259#L231)
``` go
func (m *MapIter) Ge(other IterTraitMap) bool
```
Ge determines if this iterator is greater than the other iterator




### <a name="MapIter.Gt">func</a> (\*MapIter) [Gt](/src/target/map_iter.go?s=5446:5491#L242)
``` go
func (m *MapIter) Gt(other IterTraitMap) bool
```
Gt determines if this iterator is greater than or equal to the other iterator




### <a name="MapIter.Inspect">func</a> (\*MapIter) [Inspect](/src/target/map_iter.go?s=5975:6044#L268)
``` go
func (m *MapIter) Inspect(f func(d collections.Element)) IterTraitMap
```
Inspect allows debug lines to be called in-between chained events




### <a name="MapIter.KeysToIterSlice">func</a> (\*MapIter) [KeysToIterSlice](/src/target/map_iter.go?s=3219:3269#L160)
``` go
func (m *MapIter) KeysToIterSlice() IterTraitSlice
```
KeysToIterSlice returns the keys as an IterTraitSlice




### <a name="MapIter.Le">func</a> (\*MapIter) [Le](/src/target/map_iter.go?s=6173:6218#L276)
``` go
func (m *MapIter) Le(other IterTraitMap) bool
```
Le determines if this iterator is less than the other iterator




### <a name="MapIter.Lt">func</a> (\*MapIter) [Lt](/src/target/map_iter.go?s=6402:6447#L287)
``` go
func (m *MapIter) Lt(other IterTraitMap) bool
```
Lt determines if this iterator is less than or equal to the other iterator




### <a name="MapIter.Map">func</a> (\*MapIter) [Map](/src/target/map_iter.go?s=6953:7032#L313)
``` go
func (m *MapIter) Map(f func(d collections.Data) collections.Data) IterTraitMap
```
Map takes a closure and creates an iterator which calls that closure on each element




### <a name="MapIter.Max">func</a> (\*MapIter) [Max](/src/target/map_iter.go?s=1593:1633#L83)
``` go
func (m *MapIter) Max() collections.Data
```
Max returns the max value in the data collection




### <a name="MapIter.Min">func</a> (\*MapIter) [Min](/src/target/map_iter.go?s=1968:2008#L106)
``` go
func (m *MapIter) Min() collections.Data
```
Min returns the min value in the data collection




### <a name="MapIter.Ne">func</a> (\*MapIter) [Ne](/src/target/map_iter.go?s=7230:7275#L323)
``` go
func (m *MapIter) Ne(other IterTraitMap) bool
```
Ne determines if this iterator is different from the other iterator




### <a name="MapIter.Next">func</a> (\*MapIter) [Next](/src/target/map_iter.go?s=7432:7484#L328)
``` go
func (m *MapIter) Next() (collections.Element, bool)
```
Next returns the next element in the collection and moves the cursor forward returns false when the iterator is fully consumed




### <a name="MapIter.Reduce">func</a> (\*MapIter) [Reduce](/src/target/map_iter.go?s=7781:7870#L345)
``` go
func (m *MapIter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data
```
Reduce reduces all values based on the fold function that are Operable and returns a single Data value




### <a name="MapIter.ValuesToIterSlice">func</a> (\*MapIter) [ValuesToIterSlice](/src/target/map_iter.go?s=3472:3524#L169)
``` go
func (m *MapIter) ValuesToIterSlice() IterTraitSlice
```
ValuesToIterSlice returns the keys as an IterTraitSlice








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
