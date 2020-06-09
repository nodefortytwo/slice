package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nodefortytwo/slice"
)

func TestFloat(t *testing.T) {
	var floats = slice.Float{1, 2, 3.5, 4}

	assert.Equal(t, 1, floats.Index(2))
	assert.Equal(t, -1, floats.Index(5))

	assert.Equal(t, false, floats.Includes(5))
	assert.Equal(t, true, floats.Includes(4))

	assert.Equal(t, 3, floats.Next(2))
	assert.Equal(t, 0, floats.Next(3))
	assert.Equal(t, 1, floats.Next(0))

	assert.Equal(t, float64(2), floats.NextValue(0))
	assert.Equal(t, float64(1), floats.NextValue(3))

	assert.Equal(t, 10.5, floats.Sum())
	assert.Equal(t, 2.625, floats.Avg())

	floats = slice.Float{4, 1, 3, 2}
	assert.Equal(t, float64(1), floats.Sort()[0])

	assert.Equal(t, 2.75, slice.MustMakeFloat([]string{"1.75", "1"}).Sum())

	assert.Equal(t, float64(1), floats.Min())
	assert.Equal(t, float64(4), floats.Max())

	assert.Panics(t, func() {
		slice.MustMakeInt([]float64{0})
	})
}
