package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nodefortytwo/slice"
)

func TestString(t *testing.T) {
	var strs = slice.String{"peach", "apple", "pear", "plum"}

	assert.Equal(t, strs.Index("apple"), 1)
	assert.Equal(t, strs.Index("bogus"), -1)

	assert.Equal(t, strs.Includes("bogus"), false)
	assert.Equal(t, strs.Includes("apple"), true)

	assert.Equal(t, strs.Next(2), 3)
	assert.Equal(t, strs.Next(3), 0)
	assert.Equal(t, strs.Next(0), 1)

	assert.Equal(t, strs.NextValue(0), "apple")
	assert.Equal(t, strs.NextValue(3), "peach")

	assert.Equal(t, strs.Any(isPeach), true)
	assert.Equal(t, strs.Filter(isPeach), slice.String{"peach"})
	assert.Equal(t, strs.Map(makePeach)[0], "peach")

	strings := slice.String{"awesome"}
	assert.Equal(t, strings.Slice(), []string{"awesome"})

	strings = slice.String{"test", "test", "test2"}
	assert.Equal(t, strings.Unique(), slice.String{"test", "test2"})
}

func isPeach(s string) bool {
	return s == "peach"
}

func makePeach(s string) string {
	return "peach"
}
