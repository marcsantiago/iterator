

# iterator
`import "/Users/marcsantiago/go/src/github.com/marcsantiago/iterator"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Iter](#Iter)
  * [func IntoIter(data []collections.Data) *Iter](#IntoIter)
  * [func New() *Iter](#New)
  * [func (i *Iter) All(f func(d collections.Data) bool) bool](#Iter.All)
  * [func (i *Iter) Any(f func(d collections.Data) bool) bool](#Iter.Any)
  * [func (i *Iter) Append(data collections.Data)](#Iter.Append)
  * [func (i *Iter) Chain(other IterTrait) IterTrait](#Iter.Chain)
  * [func (i *Iter) CollectBools() []bool](#Iter.CollectBools)
  * [func (i *Iter) CollectFloat32s() []float32](#Iter.CollectFloat32s)
  * [func (i *Iter) CollectFloat64s() []float64](#Iter.CollectFloat64s)
  * [func (i *Iter) CollectInt32s() []int32](#Iter.CollectInt32s)
  * [func (i *Iter) CollectInt64s() []int64](#Iter.CollectInt64s)
  * [func (i *Iter) CollectInts() []int](#Iter.CollectInts)
  * [func (i *Iter) CollectStrings() []string](#Iter.CollectStrings)
  * [func (i *Iter) Count() int](#Iter.Count)
  * [func (i *Iter) Cycle()](#Iter.Cycle)
  * [func (i *Iter) Eq(other IterTrait) bool](#Iter.Eq)
  * [func (i *Iter) Filter(f func(d collections.Data) bool) IterTrait](#Iter.Filter)
  * [func (i *Iter) Find(f func(d collections.Data) bool) collections.Data](#Iter.Find)
  * [func (i *Iter) Flatten() IterTrait](#Iter.Flatten)
  * [func (i *Iter) Fold(init collections.OperableData, f func(result collections.OperableData, next collections.OperableData) collections.Data) collections.Data](#Iter.Fold)
  * [func (i *Iter) Ge(other IterTrait) bool](#Iter.Ge)
  * [func (i *Iter) Gt(other IterTrait) bool](#Iter.Gt)
  * [func (i *Iter) Inspect(f func(d collections.Data)) IterTrait](#Iter.Inspect)
  * [func (i *Iter) Iterate() &lt;-chan collections.Data](#Iter.Iterate)
  * [func (i *Iter) Last() (int, collections.Data)](#Iter.Last)
  * [func (i *Iter) Le(other IterTrait) bool](#Iter.Le)
  * [func (i *Iter) Len() int](#Iter.Len)
  * [func (i *Iter) Lt(other IterTrait) bool](#Iter.Lt)
  * [func (i *Iter) Map(f func(d collections.Data) collections.Data) IterTrait](#Iter.Map)
  * [func (i *Iter) Max() collections.Data](#Iter.Max)
  * [func (i *Iter) Min() collections.Data](#Iter.Min)
  * [func (i *Iter) Ne(other IterTrait) bool](#Iter.Ne)
  * [func (i *Iter) Next() (int, collections.Data)](#Iter.Next)
  * [func (i *Iter) Nth(n int) collections.Data](#Iter.Nth)
  * [func (i *Iter) Partition(f func(d collections.Data) bool) ([]collections.Data, []collections.Data)](#Iter.Partition)
  * [func (i *Iter) Peak() collections.Data](#Iter.Peak)
  * [func (i *Iter) Position(f func(d collections.Data) bool) int](#Iter.Position)
  * [func (i *Iter) Product() collections.Data](#Iter.Product)
  * [func (i *Iter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data](#Iter.Reduce)
  * [func (i *Iter) Sum() collections.Data](#Iter.Sum)
  * [func (i *Iter) Take(n int)](#Iter.Take)
* [type IterTrait](#IterTrait)


#### <a name="pkg-files">Package files</a>
[iter.go](/src/target/iter.go) [iterator_iface.go](/src/target/iterator_iface.go) 






## <a name="Iter">type</a> [Iter](/src/target/iter.go?s=152:263#L15)
``` go
type Iter struct {
    // contains filtered or unexported fields
}
```






### <a name="IntoIter">func</a> [IntoIter](/src/target/iter.go?s=335:379#L28)
``` go
func IntoIter(data []collections.Data) *Iter
```

### <a name="New">func</a> [New](/src/target/iter.go?s=297:313#L24)
``` go
func New() *Iter
```




### <a name="Iter.All">func</a> (\*Iter) [All](/src/target/iter.go?s=1368:1424#L60)
``` go
func (i *Iter) All(f func(d collections.Data) bool) bool
```
All returns true if all the values match the predicate




### <a name="Iter.Any">func</a> (\*Iter) [Any](/src/target/iter.go?s=1579:1635#L70)
``` go
func (i *Iter) Any(f func(d collections.Data) bool) bool
```
Any returns true if any the value matches the predicate




### <a name="Iter.Append">func</a> (\*Iter) [Append](/src/target/iter.go?s=542:586#L35)
``` go
func (i *Iter) Append(data collections.Data)
```
Append not very Rust like, but adding a convenience to append into internal data without needing to pass in a full copy




### <a name="Iter.Chain">func</a> (\*Iter) [Chain](/src/target/iter.go?s=1808:1855#L80)
``` go
func (i *Iter) Chain(other IterTrait) IterTrait
```
Chain Takes two iterators and creates a new iterator over both in sequence




### <a name="Iter.CollectBools">func</a> (\*Iter) [CollectBools](/src/target/iter.go?s=3670:3706#L154)
``` go
func (i *Iter) CollectBools() []bool
```
CollectBools converts []collections.Data into the concrete type []bools




### <a name="Iter.CollectFloat32s">func</a> (\*Iter) [CollectFloat32s](/src/target/iter.go?s=2925:2967#L127)
``` go
func (i *Iter) CollectFloat32s() []float32
```
CollectFloat32s converts []collections.Data into the concrete type []float32




### <a name="Iter.CollectFloat64s">func</a> (\*Iter) [CollectFloat64s](/src/target/iter.go?s=3177:3219#L136)
``` go
func (i *Iter) CollectFloat64s() []float64
```
CollectFloat64s converts []collections.Data into the concrete type []float64




### <a name="Iter.CollectInt32s">func</a> (\*Iter) [CollectInt32s](/src/target/iter.go?s=2441:2479#L109)
``` go
func (i *Iter) CollectInt32s() []int32
```
CollectInt32s converts []collections.Data into the concrete type []int32




### <a name="Iter.CollectInt64s">func</a> (\*Iter) [CollectInt64s](/src/target/iter.go?s=2681:2719#L118)
``` go
func (i *Iter) CollectInt64s() []int64
```
CollectInt64s converts []collections.Data into the concrete type []int64




### <a name="Iter.CollectInts">func</a> (\*Iter) [CollectInts](/src/target/iter.go?s=2209:2243#L100)
``` go
func (i *Iter) CollectInts() []int
```
CollectInts converts []collections.Data into the concrete type []int




### <a name="Iter.CollectStrings">func</a> (\*Iter) [CollectStrings](/src/target/iter.go?s=3427:3467#L145)
``` go
func (i *Iter) CollectStrings() []string
```
CollectStrings converts []collections.Data into the concrete type []string




### <a name="Iter.Count">func</a> (\*Iter) [Count](/src/target/iter.go?s=3900:3926#L163)
``` go
func (i *Iter) Count() int
```
Count consumes the iterator and returns the count of non nil items




### <a name="Iter.Cycle">func</a> (\*Iter) [Cycle](/src/target/iter.go?s=4175:4197#L177)
``` go
func (i *Iter) Cycle()
```
Cycle enables Next or Last to iterate forever cycling through the items in order seen




### <a name="Iter.Eq">func</a> (\*Iter) [Eq](/src/target/iter.go?s=4293:4332#L182)
``` go
func (i *Iter) Eq(other IterTrait) bool
```
Eq determines if this iterator is the same as the other iterator




### <a name="Iter.Filter">func</a> (\*Iter) [Filter](/src/target/iter.go?s=4923:4987#L217)
``` go
func (i *Iter) Filter(f func(d collections.Data) bool) IterTrait
```
Filter removes all values by which the comparison function returns true




### <a name="Iter.Find">func</a> (\*Iter) [Find](/src/target/iter.go?s=5246:5315#L232)
``` go
func (i *Iter) Find(f func(d collections.Data) bool) collections.Data
```
Find returns the first value that matches the find function, else it returns nil




### <a name="Iter.Flatten">func</a> (\*Iter) [Flatten](/src/target/iter.go?s=5501:5535#L242)
``` go
func (i *Iter) Flatten() IterTrait
```
Flatten is currently a stub as Iter only supports 1 dimensional slices




### <a name="Iter.Fold">func</a> (\*Iter) [Fold](/src/target/iter.go?s=5653:5809#L247)
``` go
func (i *Iter) Fold(init collections.OperableData, f func(result collections.OperableData, next collections.OperableData) collections.Data) collections.Data
```
Fold folds all values based on the fold function that are Operable and returns a single Data value




### <a name="Iter.Ge">func</a> (\*Iter) [Ge](/src/target/iter.go?s=6096:6135#L258)
``` go
func (i *Iter) Ge(other IterTrait) bool
```
Ge determines if this iterator is greater than the other iterator




### <a name="Iter.Gt">func</a> (\*Iter) [Gt](/src/target/iter.go?s=6322:6361#L269)
``` go
func (i *Iter) Gt(other IterTrait) bool
```
Gt determines if this iterator is greater than or equal to the other iterator




### <a name="Iter.Inspect">func</a> (\*Iter) [Inspect](/src/target/iter.go?s=6748:6808#L297)
``` go
func (i *Iter) Inspect(f func(d collections.Data)) IterTrait
```
Inspect allows debug lines to be called in-between chained events




### <a name="Iter.Iterate">func</a> (\*Iter) [Iterate](/src/target/iter.go?s=4661:4709#L205)
``` go
func (i *Iter) Iterate() <-chan collections.Data
```
Iterate returns a channel of values that can be ranged over




### <a name="Iter.Last">func</a> (\*Iter) [Last](/src/target/iter.go?s=7125:7170#L306)
``` go
func (i *Iter) Last() (int, collections.Data)
```
Last returns the next element in the collection in the reverse order and its enumerated position and moves the cursor forward
If Last is called Next cannot be called for this copy of Iter, returns -1 as an index when the iterator is fully consumed




### <a name="Iter.Le">func</a> (\*Iter) [Le](/src/target/iter.go?s=7672:7711#L326)
``` go
func (i *Iter) Le(other IterTrait) bool
```
Le determines if this iterator is less than the other iterator




### <a name="Iter.Len">func</a> (\*Iter) [Len](/src/target/iter.go?s=7879:7903#L337)
``` go
func (i *Iter) Len() int
```
Len returns the current length of the underline data slice




### <a name="Iter.Lt">func</a> (\*Iter) [Lt](/src/target/iter.go?s=8009:8048#L342)
``` go
func (i *Iter) Lt(other IterTrait) bool
```
Lt determines if this iterator is less than or equal to the other iterator




### <a name="Iter.Map">func</a> (\*Iter) [Map](/src/target/iter.go?s=8451:8524#L370)
``` go
func (i *Iter) Map(f func(d collections.Data) collections.Data) IterTrait
```
Map takes a closure and creates an iterator which calls that closure on each element




### <a name="Iter.Max">func</a> (\*Iter) [Max](/src/target/iter.go?s=8669:8706#L378)
``` go
func (i *Iter) Max() collections.Data
```
Max returns the max value in the data collection




### <a name="Iter.Min">func</a> (\*Iter) [Min](/src/target/iter.go?s=8940:8977#L392)
``` go
func (i *Iter) Min() collections.Data
```
Min returns the min value in the data collection




### <a name="Iter.Ne">func</a> (\*Iter) [Ne](/src/target/iter.go?s=9227:9266#L406)
``` go
func (i *Iter) Ne(other IterTrait) bool
```
Ne determines if this iterator is different from the other iterator




### <a name="Iter.Next">func</a> (\*Iter) [Next](/src/target/iter.go?s=860:905#L41)
``` go
func (i *Iter) Next() (int, collections.Data)
```
Next returns the next element in the collection and its enumerated position and moves the cursor forward
If Next is called Last cannot be called for this copy of Iter, returns -1 as an index when the iterator is fully consumed




### <a name="Iter.Nth">func</a> (\*Iter) [Nth](/src/target/iter.go?s=9293:9335#L410)
``` go
func (i *Iter) Nth(n int) collections.Data
```



### <a name="Iter.Partition">func</a> (\*Iter) [Partition](/src/target/iter.go?s=9509:9607#L420)
``` go
func (i *Iter) Partition(f func(d collections.Data) bool) ([]collections.Data, []collections.Data)
```
Partition consumes an iterator, creating two collections from it.




### <a name="Iter.Peak">func</a> (\*Iter) [Peak](/src/target/iter.go?s=9909:9947#L434)
``` go
func (i *Iter) Peak() collections.Data
```
Peak returns the next value in the collection without consuming the iterator




### <a name="Iter.Position">func</a> (\*Iter) [Position](/src/target/iter.go?s=10125:10185#L442)
``` go
func (i *Iter) Position(f func(d collections.Data) bool) int
```
Position returns the index value of the first matching element as defined the function




### <a name="Iter.Product">func</a> (\*Iter) [Product](/src/target/iter.go?s=10348:10389#L452)
``` go
func (i *Iter) Product() collections.Data
```
Product iterates over the entire iterator, multiplying all the elements




### <a name="Iter.Reduce">func</a> (\*Iter) [Reduce](/src/target/iter.go?s=10942:11028#L475)
``` go
func (i *Iter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data
```
Reduce reduces all values based on the fold function that are Operable and returns a single Data value




### <a name="Iter.Sum">func</a> (\*Iter) [Sum](/src/target/iter.go?s=11379:11416#L489)
``` go
func (i *Iter) Sum() collections.Data
```
Sum sums the elements of an iterator.




### <a name="Iter.Take">func</a> (\*Iter) [Take](/src/target/iter.go?s=11281:11307#L484)
``` go
func (i *Iter) Take(n int)
```
Take creates an iterator that yields the first n elements, or fewer if the underlying iterator ends sooner.




## <a name="IterTrait">type</a> [IterTrait](/src/target/iterator_iface.go?s=127:1545#L8)
``` go
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
```













- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
