package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nodefortytwo/slice"
)

func TestInt(t *testing.T) {
	var ints = slice.Int{1, 2, 3, 4}

	assert.Equal(t, 1, ints.Index(2))
	assert.Equal(t, -1, ints.Index(5))

	assert.Equal(t, false, ints.Includes(5))
	assert.Equal(t, true, ints.Includes(4))

	assert.Equal(t, 3, ints.Next(2))
	assert.Equal(t, 0, ints.Next(3))
	assert.Equal(t, 1, ints.Next(0))

	assert.Equal(t, int64(2), ints.NextValue(0))
	assert.Equal(t, int64(1), ints.NextValue(3))

	assert.Equal(t, int64(10), ints.Sum())
	assert.Equal(t, 2.5, ints.Avg())

	ints = slice.Int{4, 1, 3, 2}
	assert.Equal(t, int64(1), ints.Sort()[0])

	assert.Equal(t, int64(2), slice.MakeInt([]string{"1", "1"}).Sum())
}
