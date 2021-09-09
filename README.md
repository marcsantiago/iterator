

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




### <a name="Iter.Collect">func</a> (\*Iter) [Collect](/src/target/iter.go?s=2213:2256#L100)
``` go
func (i *Iter) Collect() []collections.Data
```
Collect converts returns the concrete []collections.Data in the iterator




### <a name="Iter.CollectBools">func</a> (\*Iter) [CollectBools](/src/target/iter.go?s=3812:3848#L159)
``` go
func (i *Iter) CollectBools() []bool
```
CollectBools converts []collections.Data into the concrete type []bools




### <a name="Iter.CollectFloat32s">func</a> (\*Iter) [CollectFloat32s](/src/target/iter.go?s=3067:3109#L132)
``` go
func (i *Iter) CollectFloat32s() []float32
```
CollectFloat32s converts []collections.Data into the concrete type []float32




### <a name="Iter.CollectFloat64s">func</a> (\*Iter) [CollectFloat64s](/src/target/iter.go?s=3319:3361#L141)
``` go
func (i *Iter) CollectFloat64s() []float64
```
CollectFloat64s converts []collections.Data into the concrete type []float64




### <a name="Iter.CollectInt32s">func</a> (\*Iter) [CollectInt32s](/src/target/iter.go?s=2583:2621#L114)
``` go
func (i *Iter) CollectInt32s() []int32
```
CollectInt32s converts []collections.Data into the concrete type []int32




### <a name="Iter.CollectInt64s">func</a> (\*Iter) [CollectInt64s](/src/target/iter.go?s=2823:2861#L123)
``` go
func (i *Iter) CollectInt64s() []int64
```
CollectInt64s converts []collections.Data into the concrete type []int64




### <a name="Iter.CollectInts">func</a> (\*Iter) [CollectInts](/src/target/iter.go?s=2351:2385#L105)
``` go
func (i *Iter) CollectInts() []int
```
CollectInts converts []collections.Data into the concrete type []int




### <a name="Iter.CollectStrings">func</a> (\*Iter) [CollectStrings](/src/target/iter.go?s=3569:3609#L150)
``` go
func (i *Iter) CollectStrings() []string
```
CollectStrings converts []collections.Data into the concrete type []string




### <a name="Iter.Count">func</a> (\*Iter) [Count](/src/target/iter.go?s=4042:4068#L168)
``` go
func (i *Iter) Count() int
```
Count consumes the iterator and returns the count of non nil items




### <a name="Iter.Cycle">func</a> (\*Iter) [Cycle](/src/target/iter.go?s=4317:4339#L182)
``` go
func (i *Iter) Cycle()
```
Cycle enables Next or Last to iterate forever cycling through the items in order seen




### <a name="Iter.Eq">func</a> (\*Iter) [Eq](/src/target/iter.go?s=4435:4474#L187)
``` go
func (i *Iter) Eq(other IterTrait) bool
```
Eq determines if this iterator is the same as the other iterator




### <a name="Iter.Filter">func</a> (\*Iter) [Filter](/src/target/iter.go?s=5065:5129#L222)
``` go
func (i *Iter) Filter(f func(d collections.Data) bool) IterTrait
```
Filter removes all values by which the comparison function returns true




### <a name="Iter.Find">func</a> (\*Iter) [Find](/src/target/iter.go?s=5388:5457#L237)
``` go
func (i *Iter) Find(f func(d collections.Data) bool) collections.Data
```
Find returns the first value that matches the find function, else it returns nil




### <a name="Iter.Flatten">func</a> (\*Iter) [Flatten](/src/target/iter.go?s=5643:5677#L247)
``` go
func (i *Iter) Flatten() IterTrait
```
Flatten is currently a stub as Iter only supports 1 dimensional slices




### <a name="Iter.Fold">func</a> (\*Iter) [Fold](/src/target/iter.go?s=5795:5951#L252)
``` go
func (i *Iter) Fold(init collections.OperableData, f func(result collections.OperableData, next collections.OperableData) collections.Data) collections.Data
```
Fold folds all values based on the fold function that are Operable and returns a single Data value




### <a name="Iter.Ge">func</a> (\*Iter) [Ge](/src/target/iter.go?s=6238:6277#L263)
``` go
func (i *Iter) Ge(other IterTrait) bool
```
Ge determines if this iterator is greater than the other iterator




### <a name="Iter.Gt">func</a> (\*Iter) [Gt](/src/target/iter.go?s=6464:6503#L274)
``` go
func (i *Iter) Gt(other IterTrait) bool
```
Gt determines if this iterator is greater than or equal to the other iterator




### <a name="Iter.Inspect">func</a> (\*Iter) [Inspect](/src/target/iter.go?s=6890:6950#L302)
``` go
func (i *Iter) Inspect(f func(d collections.Data)) IterTrait
```
Inspect allows debug lines to be called in-between chained events




### <a name="Iter.Iterate">func</a> (\*Iter) [Iterate](/src/target/iter.go?s=4803:4851#L210)
``` go
func (i *Iter) Iterate() <-chan collections.Data
```
Iterate returns a channel of values that can be ranged over




### <a name="Iter.Last">func</a> (\*Iter) [Last](/src/target/iter.go?s=7267:7312#L311)
``` go
func (i *Iter) Last() (int, collections.Data)
```
Last returns the next element in the collection in the reverse order and its enumerated position and moves the cursor forward
If Last is called Next cannot be called for this copy of Iter, returns -1 as an index when the iterator is fully consumed




### <a name="Iter.Le">func</a> (\*Iter) [Le](/src/target/iter.go?s=7814:7853#L331)
``` go
func (i *Iter) Le(other IterTrait) bool
```
Le determines if this iterator is less than the other iterator




### <a name="Iter.Len">func</a> (\*Iter) [Len](/src/target/iter.go?s=8021:8045#L342)
``` go
func (i *Iter) Len() int
```
Len returns the current length of the underline data slice




### <a name="Iter.Lt">func</a> (\*Iter) [Lt](/src/target/iter.go?s=8151:8190#L347)
``` go
func (i *Iter) Lt(other IterTrait) bool
```
Lt determines if this iterator is less than or equal to the other iterator




### <a name="Iter.Map">func</a> (\*Iter) [Map](/src/target/iter.go?s=8593:8666#L375)
``` go
func (i *Iter) Map(f func(d collections.Data) collections.Data) IterTrait
```
Map takes a closure and creates an iterator which calls that closure on each element




### <a name="Iter.Max">func</a> (\*Iter) [Max](/src/target/iter.go?s=8811:8848#L383)
``` go
func (i *Iter) Max() collections.Data
```
Max returns the max value in the data collection




### <a name="Iter.Min">func</a> (\*Iter) [Min](/src/target/iter.go?s=9082:9119#L397)
``` go
func (i *Iter) Min() collections.Data
```
Min returns the min value in the data collection




### <a name="Iter.Ne">func</a> (\*Iter) [Ne](/src/target/iter.go?s=9369:9408#L411)
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




### <a name="Iter.Nth">func</a> (\*Iter) [Nth](/src/target/iter.go?s=9435:9477#L415)
``` go
func (i *Iter) Nth(n int) collections.Data
```



### <a name="Iter.Partition">func</a> (\*Iter) [Partition](/src/target/iter.go?s=9651:9749#L425)
``` go
func (i *Iter) Partition(f func(d collections.Data) bool) ([]collections.Data, []collections.Data)
```
Partition consumes an iterator, creating two collections from it.




### <a name="Iter.Peak">func</a> (\*Iter) [Peak](/src/target/iter.go?s=10051:10089#L439)
``` go
func (i *Iter) Peak() collections.Data
```
Peak returns the next value in the collection without consuming the iterator




### <a name="Iter.Position">func</a> (\*Iter) [Position](/src/target/iter.go?s=10267:10327#L447)
``` go
func (i *Iter) Position(f func(d collections.Data) bool) int
```
Position returns the index value of the first matching element as defined the function




### <a name="Iter.Product">func</a> (\*Iter) [Product](/src/target/iter.go?s=10490:10531#L457)
``` go
func (i *Iter) Product() collections.Data
```
Product iterates over the entire iterator, multiplying all the elements




### <a name="Iter.Reduce">func</a> (\*Iter) [Reduce](/src/target/iter.go?s=11084:11170#L480)
``` go
func (i *Iter) Reduce(f func(a, b collections.Data) collections.Data) collections.Data
```
Reduce reduces all values based on the fold function that are Operable and returns a single Data value




### <a name="Iter.Sum">func</a> (\*Iter) [Sum](/src/target/iter.go?s=11521:11558#L494)
``` go
func (i *Iter) Sum() collections.Data
```
Sum sums the elements of an iterator.




### <a name="Iter.Take">func</a> (\*Iter) [Take](/src/target/iter.go?s=11423:11449#L489)
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
