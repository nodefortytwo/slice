package slice

import (
	"fmt"
	"sort"
	"strconv"
)

type Float []float64

func MustMakeFloat(i interface{}) Float {
	is, err := MakeFloat(i)
	if err != nil {
		panic(err)
	}

	return is
}

func MakeFloat(i interface{}) (Float, error) {
	switch v := i.(type) {
	case []int:
		return MakeFloatFromFloat(v), nil
	case []int8:
		return MakeFloatFromFloat8(v), nil
	case []int16:
		return MakeFloatFromFloat16(v), nil
	case []int32:
		return MakeFloatFromFloat32(v), nil
	case []float64:
		return v, nil
	case []string:
		return MakeFloatFromString(v), nil
	default:
		return nil, &SliceError{fmt.Sprintf(ErrorInvalidType, v, "slice.Float")}
	}
}

func (is Float) Index(t float64) int {
	for i, v := range is {
		if v == t {
			return i
		}
	}

	return -1
}

func (is Float) Includes(t float64) bool {
	return is.Index(t) >= 0
}

func (is Float) Next(i int) int {
	i++
	if i >= is.Len() {
		return 0
	}

	return i
}

func (is Float) NextValue(i int) float64 {
	return is[is.Next(i)]
}

func (is Float) Sum() float64 {
	var sum float64
	for _, v := range is {
		sum += v
	}

	return sum
}

func (is Float) Avg() float64 {
	return is.Sum() / float64(is.Len())
}

func (is Float) Sort() Float {
	sort.Sort(is)
	return is
}

func (is Float) Min() float64 {
	min := is[0]
	for _, v := range is {
		if v < min {
			min = v
		}
	}

	return min
}

func (is Float) Max() float64 {
	var max float64
	for _, v := range is {
		if v > max {
			max = v
		}
	}

	return max
}

func (is Float) Len() int           { return len(is) }
func (is Float) Less(i, j int) bool { return is[i] < is[j] }
func (is Float) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }

func (is Float) Any(f func(float64) bool) bool {
	for _, v := range is {
		if f(v) {
			return true
		}
	}

	return false
}

func (is Float) All(f func(float64) bool) bool {
	for _, v := range is {
		if !f(v) {
			return false
		}
	}

	return true
}

func (is Float) Filter(f func(float64) bool) Float {
	vsf := make(Float, 0)

	for _, v := range is {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

func (is Float) Map(f func(float64) float64) Float {
	vsm := make([]float64, len(is))
	for i, v := range is {
		vsm[i] = f(v)
	}

	return vsm
}

func (is Float) Slice() []float64 {
	return is
}

func MakeFloatFromFloat(i []int) Float {
	is := make(Float, len(i))

	for k, v := range i {
		is[k] = float64(v)
	}

	return is
}

func MakeFloatFromFloat8(i []int8) Float {
	is := make(Float, len(i))

	for k, v := range i {
		is[k] = float64(v)
	}

	return is
}

func MakeFloatFromFloat16(i []int16) Float {
	is := make(Float, len(i))

	for k, v := range i {
		is[k] = float64(v)
	}

	return is
}

func MakeFloatFromFloat32(i []int32) Float {
	is := make(Float, len(i))

	for k, v := range i {
		is[k] = float64(v)
	}

	return is
}

func MakeFloatFromString(i []string) Float {
	is := make(Float, len(i))

	for _, val := range i {
		number, _ := strconv.ParseFloat(val, 64)
		is = append(is, number)
	}

	return is
}
