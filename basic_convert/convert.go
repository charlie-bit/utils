package basic_convert

import (
	"encoding/json"
	"reflect"
	"strconv"
)

type BasicTypeConversionUtil interface {
	StrToInt(s string) int
	StrToInt64(s string) int64
	StrToInt32(s string) int32
	StrToFloat64(s string) float64
	StrToByte(s string) []byte
	IntToStr(i int) string
	IntToInt32(i int) int32
	IntToInt64(i int) int64
	Int32ToInt(i int32) int
	Int32ToInt64(i int32) int64
	Int64ToInt(i int64) int
	Int64ToInt32(i int64) int32
	Int64ToStr(i int64) string
	GetStringInterface(i interface{}) string
	GetUint32Interface(i interface{}) uint32
	AnyToInt(value interface{}) int
	AnyToFloat64(value interface{}) float64
	InterfaceListToInt(data []interface{}) []int
	StringListToIntList(data []string) []int
	Itoa(a interface{}) string
	AtoUint64(a string) uint64
	AtoUint32(a string) uint32
	Btoi(a bool) int
	ChangeInt32SliceToInt(data []int32) []int
	ChangeIntSliceToInt32(data []int) []int32
	ChangeIntSliceToUint32(data []int) []uint32
}

var NewBasicTypeConversion = basicTypeConversion{}

type basicTypeConversion struct {
}

func (basicTypeConversion) StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func (basicTypeConversion) StrToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func (basicTypeConversion) StrToInt32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

func (basicTypeConversion) StrToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func (basicTypeConversion) StrToByte(s string) []byte {
	return []byte(s)
}

func (basicTypeConversion) IntToStr(i int) string {
	return strconv.Itoa(i)
}

func (basicTypeConversion) IntToInt32(i int) int32 {
	return int32(i)
}

func (basicTypeConversion) IntToInt64(i int) int64 {
	return int64(i)
}

func (basicTypeConversion) Int32ToInt(i int32) int {
	return int(i)
}

func (basicTypeConversion) Int32ToInt64(i int32) int64 {
	return int64(i)
}

func (basicTypeConversion) Int64ToInt(i int64) int {
	return int(i)
}

func (basicTypeConversion) Int64ToInt32(i int64) int32 {
	return int32(i)
}

func (basicTypeConversion) Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func (basicTypeConversion) GetStringInterface(i interface{}) string {
	s, ok := i.(string)
	if !ok {
		return ""
	}
	return s
}

func (basicTypeConversion) GetUint32Interface(i interface{}) uint32 {
	u, ok := i.(uint32)
	if !ok {
		return 0
	}
	return u
}

func (basicTypeConversion) AnyToInt(value interface{}) int {
	if value == nil {
		return 0
	}
	switch val := value.(type) {
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case *string:
		v, err := strconv.Atoi(*val)
		if err != nil {
			return 0
		}
		return v
	case string:
		v, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return v
	case float32:
		return int(val)
	case float64:
		return int(val)
	case bool:
		if val {
			return 1
		} else {
			return 0
		}
	case json.Number:
		v, _ := val.Int64()
		return int(v)
	}

	return 0
}

func (basicTypeConversion) AnyToFloat64(value interface{}) float64 {
	if value == nil {
		return 0
	}
	switch val := value.(type) {
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)

	case *string:
		float, err := strconv.ParseFloat(*val, 64)
		if err != nil {
			return float64(0)
		}
		return float
	case string:
		float, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return float64(0)
		}
		return float
	case bool:
		if val {
			return float64(1)
		} else {
			return float64(0)
		}
	case json.Number:
		v, _ := val.Float64()
		return v
	}
	return float64(0)
}

func (b basicTypeConversion) InterfaceListToInt(data []interface{}) []int {
	result := make([]int, len(data))
	for i, info := range data {
		if info == nil {
			result[i] = 0

			continue
		}
		result[i] = b.AnyToInt(info)
	}

	return result
}

func (b basicTypeConversion) StringListToIntList(data []string) []int {
	result := make([]int, len(data))
	for i, info := range data {
		result[i] = b.StrToInt(info)
	}
	return result
}

func (basicTypeConversion) Itoa(a interface{}) string {
	switch at := a.(type) {
	case int, int8, int16, int64, int32:
		return strconv.FormatInt(reflect.ValueOf(a).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatInt(int64(reflect.ValueOf(a).Uint()), 10)
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(a).Float(), 'f', 0, 64)
	case string:
		return at
	}
	return ""
}

func (basicTypeConversion) AtoUint64(a string) uint64 {
	r, e := strconv.ParseUint(a, 10, 64)
	if e == nil {
		return r
	}
	return 0
}

func (basicTypeConversion) AtoUint32(a string) uint32 {
	r, e := strconv.ParseUint(a, 10, 32)
	if e == nil {
		return uint32(r)
	}
	return 0
}

func (basicTypeConversion) Btoi(a bool) int {
	if a {
		return 1
	}
	return 0
}

func (basicTypeConversion) ChangeInt32SliceToInt(data []int32) []int {
	result := make([]int, len(data))
	for i, info := range data {
		result[i] = int(info)
	}
	return result
}

func (basicTypeConversion) ChangeIntSliceToInt32(data []int) []int32 {
	result := make([]int32, len(data))
	for i, info := range data {
		result[i] = int32(info)
	}
	return result
}

func (basicTypeConversion) ChangeIntSliceToUint32(data []int) []uint32 {
	result := make([]uint32, len(data))
	for i, info := range data {
		result[i] = uint32(info)
	}
	return result
}
