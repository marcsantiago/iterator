package iterator

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/marcsantiago/collections/set"

	"github.com/marcsantiago/collections"
)

func TestMapIter_All(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "all the values should be even",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(4)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(2)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 0 {
						return true
					}
					return false
				},
			},
			want: true,
		},
		{
			name: "not all the values are even",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(4)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(7)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(2)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 0 {
						return true
					}
					return false
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.All(tt.args.f); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Any(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "at least 1 value is even",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 0 {
						return true
					}
					return false
				},
			},
			want: true,
		},
		{
			name: "no values are even",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(7)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 0 {
						return true
					}
					return false
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Any(tt.args.f); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Chain(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		other IterTraitMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IterTraitMap
	}{
		{
			name: "should chain two iterators",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("d"), Value: collections.IntValue(7)},
					{Key: collections.StringValue("f"), Value: collections.IntValue(4)},
					{Key: collections.StringValue("g"), Value: collections.IntValue(9)},
				}),
			},
			want: NewMapIterFromElements([]collections.Element{
				{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
				{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
				{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				{Key: collections.StringValue("d"), Value: collections.IntValue(7)},
				{Key: collections.StringValue("f"), Value: collections.IntValue(4)},
				{Key: collections.StringValue("g"), Value: collections.IntValue(9)},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}

			got := IntoIter(m.Chain(tt.args.other).CollectValues()).CollectInts()
			wantC := IntoIter(tt.want.CollectValues()).CollectInts()
			sort.Sort(sort.IntSlice(got))
			sort.Sort(sort.IntSlice(wantC))
			if !reflect.DeepEqual(got, wantC) {
				t.Errorf("Chain() = got %v, want %v", got, wantC)
			}
		})
	}
}

func TestMapIter_Collect(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Element
	}{
		{
			name: "should convert the internal map into a slice of []collections.Element",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: []collections.Element{
				{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
				{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
				{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Collect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_CollectKeys(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Data
	}{
		{
			name: "should convert return the maps keys as a slice of []collections.Data",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: []collections.Data{
				collections.StringValue("a"),
				collections.StringValue("b"),
				collections.StringValue("c"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			got := IntoIter(m.CollectKeys()).CollectStrings()
			sort.Sort(sort.StringSlice(got))
			wantC := IntoIter(tt.want).CollectStrings()
			if !reflect.DeepEqual(got, wantC) {
				t.Errorf("CollectKeys() = %v, want %v", got, wantC)
			}
		})
	}
}

func TestMapIter_CollectValues(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Data
	}{
		{
			name: "should convert return the maps values as a slice of []collections.Data",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: []collections.Data{
				collections.IntValue(1),
				collections.IntValue(3),
				collections.IntValue(6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}

			got := IntoIter(m.CollectValues()).CollectInts()
			sort.Sort(sort.IntSlice(got))
			wantC := IntoIter(tt.want).CollectInts()
			if !reflect.DeepEqual(got, wantC) {
				t.Errorf("CollectValues() = %v, want %v", got, wantC)
			}
		})
	}
}

func TestMapIter_Count(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "should count 3 non nil items",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("k"), Value: nil},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("x"), Value: nil},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Eq(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		other IterTraitMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should be equal",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should not be equal values are different",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(7)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: false,
		},
		{
			name: "should not be equal sizes are different",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					}),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Eq(tt.args.other); got != tt.want {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Filter(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IterTraitMap
	}{
		{
			name: "should filter out all positive numbers",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 0 {
						return false
					}
					return true
				},
			},
			want: &MapIter{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
		},
		{
			name: "should filter out all positive numbers, but there are none",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(5)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 0 {
						return false
					}
					return true
				},
			},
			want: &MapIter{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(5)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Filter(tt.args.f); !got.Eq(tt.want) {
				t.Errorf("Filter() = %v, want %v", got.Collect(), tt.want.Collect())
			}
		})
	}
}

func TestMapIter_FilterKeys(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IterTraitMap
	}{
		{
			name: "should filter out keys that start with a",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if strings.HasPrefix(d.String(), "a") {
						return false
					}
					return true
				},
			},
			want: &MapIter{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
		},
		{
			name: "should all the keys since they all start with a",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("aa"), Value: collections.IntValue(5)},
					{Key: collections.StringValue("aaa"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if strings.HasPrefix(d.String(), "a") {
						return false
					}
					return true
				},
			},
			want: &MapIter{
				internalMap: NewMapIterFromElements([]collections.Element{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.FilterKeys(tt.args.f); !got.Eq(tt.want) {
				t.Errorf("FilterKeys() = %v, want %v", got.Collect(), tt.want.Collect())
			}
		})
	}
}

func TestMapIter_Find(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantItem collections.Element
		wantOk   bool
	}{
		{
			name: "should find the target value",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("aa"), Value: collections.IntValue(5)},
					{Key: collections.StringValue("aaa"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d == collections.StringValue("aaa") {
						return true
					}
					return false
				},
			},
			wantItem: collections.Element{Key: collections.StringValue("aaa"), Value: collections.IntValue(3)},
			wantOk:   true,
		},
		{
			name: "should not find the target value",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("aa"), Value: collections.IntValue(5)},
					{Key: collections.StringValue("aaa"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d == collections.FloatValue64(11) {
						return true
					}
					return false
				},
			},
			wantItem: collections.Element{},
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			gotItem, gotOk := m.Find(tt.args.f)
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("Find() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Find() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMapIter_FindByValue(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantItem collections.Element
		wantOk   bool
	}{
		{
			name: "should find the target value",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("aa"), Value: collections.IntValue(5)},
					{Key: collections.StringValue("aaa"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d == collections.IntValue(5) {
						return true
					}
					return false
				},
			},
			wantItem: collections.Element{Key: collections.StringValue("aa"), Value: collections.IntValue(5)},
			wantOk:   true,
		},
		{
			name: "should not find the target value",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("aa"), Value: collections.IntValue(5)},
					{Key: collections.StringValue("aaa"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(d collections.Data) bool {
					if d == collections.FloatValue64(11) {
						return true
					}
					return false
				},
			},
			wantItem: collections.Element{},
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			gotItem, gotOk := m.FindByValue(tt.args.f)
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("FindByValue() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FindByValue() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMapIter_Ge(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		other IterTraitMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should be equal",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should be greater than there are more values",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should be greater than at least 1 value is larger",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(11)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should not be greater or equal than it has less values",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Ge(tt.args.other); got != tt.want {
				t.Errorf("Ge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Gt(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		other IterTraitMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should not be greater they are equal",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: false,
		},
		{
			name: "should be greater than there are more values",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should be greater than at least 1 value is larger",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(11)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should not be greater or equal than it has less values",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Gt(tt.args.other); got != tt.want {
				t.Errorf("Gt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_KeysToIterSlice(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   IterTraitSlice
	}{
		{
			name: "should convert keys to an IterSlice",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: IntoIter([]collections.Data{
				collections.StringValue("a"),
				collections.StringValue("b"),
				collections.StringValue("c"),
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			got := m.KeysToIterSlice().CollectStrings()
			sort.Sort(sort.StringSlice(got))

			if !reflect.DeepEqual(got, tt.want.CollectStrings()) {
				t.Errorf("KeysToIterSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Le(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		other IterTraitMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should be equal",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should be less than there are less values",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should be less than at least 1 value is smaller",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(2)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Le(tt.args.other); got != tt.want {
				t.Errorf("Le() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Lt(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		other IterTraitMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should be equal and therefore not less",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: false,
		},
		{
			name: "should be less than there are less values",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should be less than at least 1 value is smaller",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(2)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Lt(tt.args.other); got != tt.want {
				t.Errorf("Lt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Map(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		f func(d collections.Data) collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IterTraitMap
	}{
		{
			name: "should double all the values",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(2)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(-2)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(4)},
				}),
			},
			args: args{f: func(d collections.Data) collections.Data {
				return d.(collections.OperableData).Mul(collections.IntValue(2))
			}},
			want: &MapIter{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(4)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(-4)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(8)},
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			got := m.Map(tt.args.f)
			if !got.Eq(tt.want) {
				t.Errorf("Map() = %v, want %v", got.Collect(), tt.want.Collect())
			}
		})
	}
}

func TestMapIter_Max(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   collections.Data
	}{
		{
			name: "should find the largest value in the collection",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(-1)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: collections.IntValue(3),
		},
		{
			name: "should return nil the collection is empty",
			fields: fields{
				internalMap: nil,
			},
			want: nil,
		},
		{
			name: "should the only value in the collection",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
				}),
			},
			want: collections.IntValue(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Max(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Min(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   collections.Data
	}{
		{
			name: "should find the smallest value in the collection",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(-1)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: collections.IntValue(-1),
		},
		{
			name: "should return nil the collection is empty",
			fields: fields{
				internalMap: nil,
			},
			want: nil,
		},
		{
			name: "should the only value in the collection",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
				}),
			},
			want: collections.IntValue(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Min(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Ne(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		other IterTraitMap
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should be not equal",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(61)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: true,
		},
		{
			name: "should be equal and therefore false",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(61)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				other: &MapIter{
					internalMap: NewMapIterFromElements([]collections.Element{
						{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
						{Key: collections.StringValue("b"), Value: collections.IntValue(61)},
						{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
					}),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Ne(tt.args.other); got != tt.want {
				t.Errorf("Ne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_Next(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Element
		i      int
	}{
		{
			name: "should be equal and therefore false",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(61)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: []collections.Element{
				{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
				{Key: collections.StringValue("b"), Value: collections.IntValue(61)},
				{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
			},
			i: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}

			keys := set.New()
			values := set.New()
			for _, e := range tt.want {
				keys.Add(e.Key)
				values.Add(e.Value)
			}

			for i := 0; i < tt.i; i++ {
				element, _ := m.Next()
				keys.Remove(element.Key)
				values.Remove(element.Value)
			}

			if len(keys) != 0 || len(values) != 0 {
				t.Errorf("Next() wanted an empty set got keys %v, values %v", keys, values)
			}
		})
	}
}

func TestMapIter_Reduce(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	type args struct {
		f func(a, b collections.Data) collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   collections.Data
	}{
		{
			name: "should reduce the collection to a single value",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(61)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			args: args{
				f: func(a, b collections.Data) collections.Data {
					if a.Equal(b) || a.Greater(b) {
						return a
					}
					return b
				},
			},
			want: collections.IntValue(61),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			if got := m.Reduce(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter_ValuesToIterSlice(t *testing.T) {
	type fields struct {
		internalMap internalMap
	}
	tests := []struct {
		name   string
		fields fields
		want   IterTraitSlice
	}{
		{
			name: "should convert keys to an IterSlice",
			fields: fields{
				internalMap: NewMapIterFromElements([]collections.Element{
					{Key: collections.StringValue("a"), Value: collections.IntValue(1)},
					{Key: collections.StringValue("b"), Value: collections.IntValue(6)},
					{Key: collections.StringValue("c"), Value: collections.IntValue(3)},
				}),
			},
			want: IntoIter([]collections.Data{
				collections.IntValue(1),
				collections.IntValue(3),
				collections.IntValue(6),
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MapIter{
				nil,
				tt.fields.internalMap,
			}
			got := m.ValuesToIterSlice().CollectInts()
			sort.Sort(sort.IntSlice(got))

			if !reflect.DeepEqual(got, tt.want.CollectInts()) {
				t.Errorf("ValuesToIterSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
