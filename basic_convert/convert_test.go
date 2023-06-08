package basic_convert

import (
	"reflect"
	"testing"
)

func Test_basicTypeConversion_StrToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `StrToInt "1" > 1`,
			args: args{
				"1",
			},
			want: 1,
		},
		{
			name: `StrToInt "a" > 0`,
			args: args{
				"a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBasicTypeConversion.StrToInt(tt.args.s); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_AnyToFloat64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: `AnyToFloat64 "1" > 1`,
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: `AnyToFloat64 "a" > 0`,
			args: args{
				value: "a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.AnyToFloat64(tt.args.value); got != tt.want {
				t.Errorf("AnyToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_AnyToInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `AnyToInt "1" > 1`,
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: `AnyToInt "a" > 0`,
			args: args{
				value: "a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.AnyToInt(tt.args.value); got != tt.want {
				t.Errorf("AnyToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_AtoUint32(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: `AtoUint32 "1" > 1`,
			args: args{
				a: "1",
			},
			want: 1,
		},
		{
			name: `AtoUint32 "a" > 0`,
			args: args{
				a: "a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.AtoUint32(tt.args.a); got != tt.want {
				t.Errorf("AtoUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_AtoUint64(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: `AtoUint64 "1" > 1`,
			args: args{
				a: "1",
			},
			want: 1,
		},
		{
			name: `AtoUint64 "a" > 0`,
			args: args{
				a: "a",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.AtoUint64(tt.args.a); got != tt.want {
				t.Errorf("AtoUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_Btoi(t *testing.T) {
	type args struct {
		a bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `Btoi true > 1`,
			args: args{
				a: true,
			},
			want: 1,
		},
		{
			name: `Btoi false > 0`,
			args: args{
				a: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.Btoi(tt.args.a); got != tt.want {
				t.Errorf("Btoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_ChangeInt32SliceToInt(t *testing.T) {
	type args struct {
		data []int32
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: `ChangeInt32SliceToInt`,
			args: args{
				[]int32{1, 2},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.ChangeInt32SliceToInt(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeInt32SliceToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_ChangeIntSliceToInt32(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: `ChangeIntSliceToInt32`,
			args: args{
				data: []int{1, 2, 3},
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.ChangeIntSliceToInt32(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeIntSliceToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_ChangeIntSliceToUint32(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: `ChangeIntSliceToUint32`,
			args: args{
				data: []int{1, 2, 3},
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.ChangeIntSliceToUint32(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangeIntSliceToUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_GetStringInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `ChangeIntSliceToUint32`,
			args: args{
				i: []int{1, 2, 3},
			},
			want: "",
		},
		{
			name: `ChangeIntSliceToUint32`,
			args: args{
				i: 1,
			},
			want: "",
		},
		{
			name: `ChangeIntSliceToUint32`,
			args: args{
				i: "a",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.GetStringInterface(tt.args.i); got != tt.want {
				t.Errorf("GetStringInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_GetUint32Interface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: `GetUint32Interface`,
			args: args{
				i: []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: `GetUint32Interface`,
			args: args{
				i: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.GetUint32Interface(tt.args.i); got != tt.want {
				t.Errorf("GetUint32Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_Int32ToInt(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `Int32ToInt`,
			args: args{
				i: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.Int32ToInt(tt.args.i); got != tt.want {
				t.Errorf("Int32ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_Int32ToInt64(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: `Int32ToInt64`,
			args: args{
				i: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.Int32ToInt64(tt.args.i); got != tt.want {
				t.Errorf("Int32ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_Int64ToInt(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				i: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.Int64ToInt(tt.args.i); got != tt.want {
				t.Errorf("Int64ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_Int64ToInt32(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				i: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.Int64ToInt32(tt.args.i); got != tt.want {
				t.Errorf("Int64ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_Int64ToStr(t *testing.T) {
	type args struct {
		i int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				i: 1,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.Int64ToStr(tt.args.i); got != tt.want {
				t.Errorf("Int64ToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_IntToInt32(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				i: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.IntToInt32(tt.args.i); got != tt.want {
				t.Errorf("IntToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_IntToInt64(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				i: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.IntToInt64(tt.args.i); got != tt.want {
				t.Errorf("IntToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_IntToStr(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				i: 1,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.IntToStr(tt.args.i); got != tt.want {
				t.Errorf("IntToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_InterfaceListToInt(t *testing.T) {
	type args struct {
		data []interface{}
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				data: []interface{}{1, "2"},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := basicTypeConversion{}
			if got := b.InterfaceListToInt(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceListToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_Itoa(t *testing.T) {
	type args struct {
		a interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				a: []interface{}{1, "2"},
			},
			want: "",
		},
		{
			args: args{
				a: 1,
			},
			want: "1",
		},
		{
			args: args{
				a: 1.02,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.Itoa(tt.args.a); got != tt.want {
				t.Errorf("Itoa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_StrToByte(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				s: "1",
			},
			want: []byte("1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.StrToByte(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrToByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_StrToFloat64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				s: "1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.StrToFloat64(tt.args.s); got != tt.want {
				t.Errorf("StrToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_StrToInt1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.StrToInt(tt.args.s); got != tt.want {
				t.Errorf("StrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_StrToInt32(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				s: "1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.StrToInt32(tt.args.s); got != tt.want {
				t.Errorf("StrToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_StrToInt64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				s: "1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := basicTypeConversion{}
			if got := ba.StrToInt64(tt.args.s); got != tt.want {
				t.Errorf("StrToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicTypeConversion_StringListToIntList(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				data: []string{"1", "2"},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := basicTypeConversion{}
			if got := b.StringListToIntList(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringListToIntList() = %v, want %v", got, tt.want)
			}
		})
	}
}
