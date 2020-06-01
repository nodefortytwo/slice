package slice

import (
	"fmt"
	"sort"
	"strconv"
)

type Int []int64

func MakeInt(i interface{}) Int {
	switch v := i.(type) {
	case []int:
		return MakeIntFromInt(v)
	case []int8:
		return MakeIntFromInt8(v)
	case []int16:
		return MakeIntFromInt16(v)
	case []int32:
		return MakeIntFromInt32(v)
	case []int64:
		return v
	case []string:
		return MakeIntFromString(v)
	default:
		panic(fmt.Sprintf("Unspported type %T", v))
	}
}

func (is Int) Index(t int64) int {
	for i, v := range is {
		if v == t {
			return i
		}
	}

	return -1
}

func (is Int) Includes(t int64) bool {
	return is.Index(t) >= 0
}

func (is Int) Next(i int) int {
	i++
	if i >= is.Len() {
		return 0
	}

	return i
}

func (is Int) NextValue(i int) int64 {
	return is[is.Next(i)]
}

func (is Int) Sum() int64 {
	var sum int64
	for _, v := range is {
		sum += v
	}

	return sum
}

func (is Int) Avg() float64 {
	return float64(is.Sum()) / float64(is.Len())
}

func (is Int) Sort() Int {
	sort.Sort(is)
	return is
}

func (is Int) Len() int           { return len(is) }
func (is Int) Less(i, j int) bool { return is[i] < is[j] }
func (is Int) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }

func (is Int) Any(f func(int64) bool) bool {
	for _, v := range is {
		if f(v) {
			return true
		}
	}

	return false
}

func (is Int) All(f func(int64) bool) bool {
	for _, v := range is {
		if !f(v) {
			return false
		}
	}

	return true
}

func (is Int) Filter(f func(int64) bool) Int {
	vsf := make(Int, 0)

	for _, v := range is {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

func (is Int) Map(f func(int64) int64) Int {
	vsm := make([]int64, len(is))
	for i, v := range is {
		vsm[i] = f(v)
	}

	return vsm
}

func (is Int) Slice() []int64 {
	return is
}

func MakeIntFromInt(i []int) Int {
	is := make(Int, len(i))

	for k, v := range i {
		is[k] = int64(v)
	}

	return is
}

func MakeIntFromInt8(i []int8) Int {
	is := make(Int, len(i))

	for k, v := range i {
		is[k] = int64(v)
	}

	return is
}

func MakeIntFromInt16(i []int16) Int {
	is := make(Int, len(i))

	for k, v := range i {
		is[k] = int64(v)
	}

	return is
}

func MakeIntFromInt32(i []int32) Int {
	is := make(Int, len(i))

	for k, v := range i {
		is[k] = int64(v)
	}

	return is
}

func MakeIntFromString(i []string) Int {
	is := make(Int, len(i))

	for _, val := range i {
		number, _ := strconv.Atoi(val)
		is = append(is, int64(number))
	}

	return is
}
