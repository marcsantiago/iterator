package iterator

import (
	"reflect"
	"testing"

	"github.com/marcsantiago/collections"
)

func TestIter_Next(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Data
		i      int
	}{
		{
			name: "should get the next item in the collection",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(4),
					collections.IntValue(1),
					collections.IntValue(2),
				},
			},
			want: []collections.Data{
				collections.IntValue(4),
				collections.IntValue(1),
				collections.IntValue(2),
				nil,
			},
			i: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}

			for ii := 0; ii < tt.i; ii++ {
				if _, got := i.Next(); !reflect.DeepEqual(got, tt.want[ii]) {
					t.Errorf("Next() = %v, want %v", got, tt.want[ii])
				}
			}
		})
	}
}

func TestIter_NextWithCycle(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Data
		i      int
	}{
		{
			name: "should get the next item in the collection",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(4),
					collections.IntValue(1),
					collections.IntValue(2),
				},
			},
			want: []collections.Data{
				collections.IntValue(4),
				collections.IntValue(1),
				collections.IntValue(2),
				collections.IntValue(4),
				collections.IntValue(1),
				collections.IntValue(2),
				collections.IntValue(4),
				collections.IntValue(1),
			},
			i: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			i.Cycle()

			for ii := 0; ii < tt.i; ii++ {
				if _, got := i.Next(); !reflect.DeepEqual(got, tt.want[ii]) {
					t.Errorf("Next() with cycle = got %v, want %v", got, tt.want[ii])
				}
			}
		})
	}
}

func TestIter_All(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
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
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(4),
					collections.IntValue(6),
					collections.IntValue(2),
				},
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
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(4),
					collections.IntValue(7),
					collections.IntValue(2),
				},
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
			i := Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.All(tt.args.f); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Any(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
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
			name: "at least 1 value must be even",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(4),
					collections.IntValue(7),
					collections.IntValue(2),
				},
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
			name: "at least 1 value must be old",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(4),
					collections.IntValue(7),
					collections.IntValue(2),
				},
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 1 {
						return true
					}
					return false
				},
			},
			want: true,
		},
		{
			name: "no values are odd",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					collections.IntValue(220),
				},
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 1 {
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
			i := Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Any(tt.args.f); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Chain(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		other IterTrait
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IterTrait
	}{
		{
			name: "should chain two iterators",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					nil,
					collections.IntValue(220),
				},
			},
			args: args{other: IntoIter([]collections.Data{
				collections.IntValue(2),
				collections.IntValue(3),
				collections.IntValue(10),
			})},
			want: IntoIter([]collections.Data{
				collections.IntValue(2),
				collections.IntValue(10),
				nil,
				collections.IntValue(220),
				collections.IntValue(2),
				collections.IntValue(3),
				collections.IntValue(10),
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}

			if got := i.Chain(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chain() = got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Collect(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}

	type want struct {
		i   []int
		i32 []int32
		i64 []int64
		f32 []float32
		f64 []float64
		s   []string
		b   []bool
		dd  []collections.Data
	}

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "should convert all values to collections",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(0),
					collections.IntValue(0),
					collections.IntValue(1),
					collections.IntValue(1),
					collections.IntValue(0),
				},
			},
			want: want{
				i:   []int{0, 0, 1, 1, 0},
				i32: []int32{0, 0, 1, 1, 0},
				i64: []int64{0, 0, 1, 1, 0},
				f32: []float32{0.0, 0.0, 1.0, 1.0, 0.0},
				f64: []float64{0.0, 0.0, 1.0, 1.0, 0.0},
				s:   []string{"0", "0", "1", "1", "0"},
				b:   []bool{false, false, true, true, false},
				dd:  []collections.Data{collections.IntValue(0), collections.IntValue(0), collections.IntValue(1), collections.IntValue(1), collections.IntValue(0)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}

			for i := 0; i < 8; i++ {
				switch i {
				case 0:
					if got := iter.CollectInts(); !reflect.DeepEqual(got, tt.want.i) {
						t.Errorf("CollectInts() = got %v, want %v", got, tt.want.i)
					}
				case 1:
					if got := iter.CollectInt32s(); !reflect.DeepEqual(got, tt.want.i32) {
						t.Errorf("CollectInt32s() = got %v, want %v", got, tt.want.i32)
					}
				case 2:
					if got := iter.CollectInt64s(); !reflect.DeepEqual(got, tt.want.i64) {
						t.Errorf("CollectInt64s() = got %v, want %v", got, tt.want.i64)
					}
				case 3:
					if got := iter.CollectFloat32s(); !reflect.DeepEqual(got, tt.want.f32) {
						t.Errorf("CollectFloat32s() = got %v, want %v", got, tt.want.f32)
					}
				case 4:
					if got := iter.CollectFloat64s(); !reflect.DeepEqual(got, tt.want.f64) {
						t.Errorf("CollectFloat64s() = got %v, want %v", got, tt.want.f64)
					}
				case 5:
					if got := iter.CollectStrings(); !reflect.DeepEqual(got, tt.want.s) {
						t.Errorf("CollectStrings() = got %v, want %v", got, tt.want.s)
					}
				case 6:
					if got := iter.CollectBools(); !reflect.DeepEqual(got, tt.want.b) {
						t.Errorf("CollectBools() = got %v, want %v", got, tt.want.b)
					}
				case 7:
					if got := iter.Collect(); !reflect.DeepEqual(got, tt.want.dd) {
						t.Errorf("Collect() = got %v, want %v", got, tt.want.dd)
					}
				}
			}

		})
	}
}

func TestIter_Count(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "should count 3 non nil items",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(4),
					nil,
					collections.IntValue(7),
					nil,
					collections.IntValue(2),
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Eq(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		other IterTrait
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
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					collections.IntValue(220),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(2),
						collections.IntValue(10),
						collections.IntValue(220),
					},
				},
			},
			want: true,
		},
		{
			name: "should not be equal values are different",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					collections.IntValue(220),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(2),
						collections.IntValue(11),
						collections.IntValue(220),
					},
				},
			},
			want: false,
		},
		{
			name: "should not be equal sizes are different",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					collections.IntValue(220),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(2),
						collections.IntValue(220),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Eq(tt.args.other); got != tt.want {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Iterate(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Data
	}{
		{
			name: "should iterate over the values using a range",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					collections.IntValue(220),
				},
			},
			want: []collections.Data{
				collections.IntValue(2),
				collections.IntValue(10),
				collections.IntValue(220),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}

			var got []collections.Data
			for data := range i.Iterate() {
				got = append(got, data)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Filter(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IterTrait
	}{
		{
			name: "should filter out all positive numbers",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(11),
					collections.IntValue(220),
				},
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 0 {
						return true
					}
					return false
				},
			},
			want: &Iter{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(11),
				},
			},
		},
		{
			name: "should filter out all positive numbers, but there are none",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.Int()%2 == 0 {
						return true
					}
					return false
				},
			},
			want: &Iter{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Filter(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Find(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   collections.Data
	}{
		{
			name: "should find the target value",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				f: func(d collections.Data) bool {
					if d == collections.IntValue(11) {
						return true
					}
					return false
				},
			},
			want: collections.IntValue(11),
		},
		{
			name: "should not find the target value",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				f: func(d collections.Data) bool {
					if d == collections.FloatValue64(11) {
						return true
					}
					return false
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Find(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Flatten(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   IterTrait
	}{
		{
			name: "should return itself since the data is already flat",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			want: &Iter{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Flatten(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Fold(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		init collections.OperableData
		f    func(result collections.OperableData, next collections.OperableData) collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   collections.Data
	}{
		{
			name: "should flatten the data to a single value",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				init: collections.IntValue(0),
				f: func(result, next collections.OperableData) collections.Data {
					return result.Add(next)
				},
			},
			want: collections.IntValue(253),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Fold(tt.args.init, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Ge(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		other IterTrait
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
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: true,
		},
		{
			name: "should be greater than there are more values",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
					},
				},
			},
			want: true,
		},
		{
			name: "should be greater than at least 1 value is larger",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(15),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: true,
		},
		{
			name: "should not be greater or equal than it has less values",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(15),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Ge(tt.args.other); got != tt.want {
				t.Errorf("Ge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Gt(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		other IterTrait
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
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: false,
		},
		{
			name: "should be greater than there are more values",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
					},
				},
			},
			want: true,
		},
		{
			name: "should be greater than at least 1 value is larger",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(15),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: true,
		},
		{
			name: "should not be greater or equal than it has less values",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(15),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Gt(tt.args.other); got != tt.want {
				t.Errorf("Gt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Last(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Data
		i      int
	}{
		{
			name: "should get the last item in the collection",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(4),
					collections.IntValue(1),
					collections.IntValue(2),
				},
			},
			want: []collections.Data{
				collections.IntValue(2),
				collections.IntValue(1),
				collections.IntValue(4),
				nil,
			},
			i: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}

			for ii := 0; ii < tt.i; ii++ {
				if _, got := i.Last(); !reflect.DeepEqual(got, tt.want[ii]) {
					t.Errorf("Last() = %v, want %v", got, tt.want[ii])
				}
			}
		})
	}
}

func TestIter_LastWithCycle(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   []collections.Data
		i      int
	}{
		{
			name: "should get the next item in the collection",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(1),
					collections.IntValue(2),
					collections.IntValue(3),
				},
			},
			want: []collections.Data{
				collections.IntValue(3),
				collections.IntValue(2),
				collections.IntValue(1),
				collections.IntValue(3),
				collections.IntValue(2),
				collections.IntValue(1),
				collections.IntValue(3),
				collections.IntValue(2),
			},
			i: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			i.Cycle()

			for ii := 0; ii < tt.i; ii++ {
				if _, got := i.Last(); !reflect.DeepEqual(got, tt.want[ii]) {
					t.Errorf("Last() with cycle = got %v, want %v", got, tt.want[ii])
				}
			}
		})
	}
}

func TestIter_Le(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
		direction  _direction
	}
	type args struct {
		other IterTrait
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
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: true,
		},
		{
			name: "should be less than there are less values",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: true,
		},
		{
			name: "should be less than at least 1 value is smaller",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(5),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
				direction:  tt.fields.direction,
			}
			if got := i.Le(tt.args.other); got != tt.want {
				t.Errorf("Le() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Len(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
		direction  _direction
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "should return the number of elements in the internal slice",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(5),
					collections.IntValue(221),
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
				direction:  tt.fields.direction,
			}
			if got := i.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Lt(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
		direction  _direction
	}
	type args struct {
		other IterTrait
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
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: false,
		},
		{
			name: "should be less than there are less values",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: true,
		},
		{
			name: "should be less than at least 1 value is smaller",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(5),
					collections.IntValue(221),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(21),
						collections.IntValue(11),
						collections.IntValue(221),
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
				direction:  tt.fields.direction,
			}
			if got := i.Lt(tt.args.other); got != tt.want {
				t.Errorf("Lt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Map(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		f func(d collections.Data) collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IterTrait
	}{
		{
			name: "should double all the values",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(1),
					collections.IntValue(2),
					collections.IntValue(3),
				},
			},
			args: args{f: func(d collections.Data) collections.Data {
				return d.(collections.OperableData).Mul(collections.IntValue(2))
			}},
			want: &Iter{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(4),
					collections.IntValue(6),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Map(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Max(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
		direction  _direction
	}
	tests := []struct {
		name   string
		fields fields
		want   collections.Data
	}{
		{
			name: "should find the largest value in the collection",
			fields: fields{
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(5),
					collections.IntValue(221),
					collections.IntValue(10),
					collections.IntValue(-1),
					collections.IntValue(40),
				},
			},
			want: collections.IntValue(221),
		},
		{
			name: "should return nil the collection is empty",
			fields: fields{
				values: []collections.Data{},
			},
			want: nil,
		},
		{
			name: "should the only value in the collection",
			fields: fields{
				values: []collections.Data{
					collections.IntValue(221),
				},
			},
			want: collections.IntValue(221),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
				direction:  tt.fields.direction,
			}
			if got := i.Max(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Min(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
		direction  _direction
	}
	tests := []struct {
		name   string
		fields fields
		want   collections.Data
	}{
		{
			name: "should find the smallest value in the collection",
			fields: fields{
				values: []collections.Data{
					collections.IntValue(21),
					collections.IntValue(11),
					collections.IntValue(5),
					collections.IntValue(221),
					collections.IntValue(10),
					collections.IntValue(-1),
					collections.IntValue(40),
				},
			},
			want: collections.IntValue(-1),
		},
		{
			name: "should return nil the collection is empty",
			fields: fields{
				values: []collections.Data{},
			},
			want: nil,
		},
		{
			name: "should the only value in the collection",
			fields: fields{
				values: []collections.Data{
					collections.IntValue(221),
				},
			},
			want: collections.IntValue(221),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
				direction:  tt.fields.direction,
			}
			if got := i.Min(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Ne(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
		direction  _direction
	}
	type args struct {
		other IterTrait
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should not be not equal",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					collections.IntValue(220),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(2),
						collections.IntValue(10),
						collections.IntValue(220),
					},
				},
			},
			want: false,
		},
		{
			name: "should be not equal values are different",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					collections.IntValue(220),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(2),
						collections.IntValue(11),
						collections.IntValue(220),
					},
				},
			},
			want: true,
		},
		{
			name: "should be not equal sizes are different",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(10),
					collections.IntValue(220),
				},
			},
			args: args{
				other: &Iter{
					currentIdx: 0,
					values: []collections.Data{
						collections.IntValue(2),
						collections.IntValue(220),
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
				direction:  tt.fields.direction,
			}
			if got := i.Ne(tt.args.other); got != tt.want {
				t.Errorf("Ne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Peak(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
		direction  _direction
	}
	tests := []struct {
		name   string
		fields fields
		want   collections.Data
	}{
		{
			name: "should return the next element without consuming the iterator",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(220),
				},
			},
			want: collections.IntValue(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
				direction:  tt.fields.direction,
			}

			for j := 0; j < 5; j++ {
				if got := i.Peak(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Peak() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestIter_Position(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
		direction  _direction
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "should find the position of the first instance that matches",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(11),
					collections.IntValue(220),
					collections.IntValue(20),
					collections.IntValue(220),
				},
			},
			args: args{
				f: func(d collections.Data) bool {
					if d == collections.IntValue(220) {
						return true
					}
					return false
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
				direction:  tt.fields.direction,
			}
			if got := i.Position(tt.args.f); got != tt.want {
				t.Errorf("Position() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Product(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   collections.Data
	}{
		{
			name: "should multiply all values in the collection",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(1),
					collections.IntValue(2),
					collections.IntValue(3),
					collections.IntValue(4),
					collections.IntValue(5),
				},
			},
			want: collections.IntValue(120),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Product(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Product() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Reduce(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
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
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(11),
					collections.IntValue(220),
					collections.IntValue(20),
					collections.IntValue(220),
				},
			},
			args: args{
				f: func(a, b collections.Data) collections.Data {
					if a.Equal(b) || a.Greater(b) {
						return a
					}
					return b
				},
			},
			want: collections.IntValue(220),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Reduce(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Take(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []collections.Data
		n      int
	}{
		{
			name: "should take the first 2 values and return nil for other next calls",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(2),
					collections.IntValue(11),
					collections.IntValue(220),
					collections.IntValue(20),
					collections.IntValue(220),
				},
			},
			args: args{n: 2},
			want: []collections.Data{
				collections.IntValue(2),
				collections.IntValue(11),
				nil,
				nil,
			},
			n: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			iter.Take(tt.args.n)

			for i := 0; i < tt.n; i++ {
				_, got := iter.Next()
				want := tt.want[i]

				if !reflect.DeepEqual(got, want) {
					t.Errorf("Take() = %v, want %v", got, want)
				}
			}
		})
	}
}

func TestIter_Sum(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	tests := []struct {
		name   string
		fields fields
		want   collections.Data
	}{
		{
			name: "should sum all the values",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(1),
					collections.IntValue(2),
					collections.IntValue(3),
					collections.IntValue(4),
					collections.IntValue(5),
					collections.IntValue(6),
					collections.IntValue(7),
					collections.IntValue(8),
					collections.IntValue(9),
					collections.IntValue(10),
				},
			},
			want: collections.IntValue(55),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_Nth(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		n int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       collections.Data
		wantLength int
	}{
		{
			name: "should return the second element",
			fields: fields{
				currentIdx: 0,
				values: []collections.Data{
					collections.IntValue(1),
					collections.IntValue(2),
					collections.IntValue(3),
					collections.IntValue(4),
					collections.IntValue(5),
					collections.IntValue(6),
					collections.IntValue(7),
					collections.IntValue(8),
					collections.IntValue(9),
					collections.IntValue(10),
				},
			},
			args: args{
				n: 2,
			},
			want:       collections.IntValue(3),
			wantLength: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			if got := i.Nth(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Nth() = %v, want %v", got, tt.want)
			}

			if i.Len() != tt.wantLength {
				t.Errorf("Nth Len() = %v, want %v", i.Len(), tt.wantLength)
			}

		})
	}
}

func TestIter_Partition(t *testing.T) {
	type fields struct {
		currentIdx int
		values     []collections.Data
	}
	type args struct {
		f func(d collections.Data) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []collections.Data
		want1  []collections.Data
	}{
		{
			name: "should create two collections, 1 even and 1 odd",
			fields: fields{
				values: []collections.Data{
					collections.IntValue(1),
					collections.IntValue(2),
					collections.IntValue(3),
					collections.IntValue(4),
					collections.IntValue(5),
					collections.IntValue(6),
					collections.IntValue(7),
					collections.IntValue(8),
					collections.IntValue(9),
					collections.IntValue(10),
				},
			},
			args: args{
				f: func(d collections.Data) bool {
					if d.(collections.OperableData).Mod(collections.IntValue(2)).Int() == 0 {
						return true
					}
					return false
				},
			},
			want: []collections.Data{
				collections.IntValue(2),
				collections.IntValue(4),
				collections.IntValue(6),
				collections.IntValue(8),
				collections.IntValue(10),
			},
			want1: []collections.Data{
				collections.IntValue(1),
				collections.IntValue(3),
				collections.IntValue(5),
				collections.IntValue(7),
				collections.IntValue(9),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Iter{
				currentIdx: tt.fields.currentIdx,
				values:     tt.fields.values,
			}
			got, got1 := i.Partition(tt.args.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Partition() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Partition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
