package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString (t *testing.T){
	var strs = []string{"peach", "apple", "pear", "plum"}


	assert.Equal(t, StrSlice(strs).Index("apple"), 1)
	assert.Equal(t, StrSlice(strs).Index("bogus"), -1)

	assert.Equal(t, StrSlice(strs).Contains("bogus"),false)
	assert.Equal(t, StrSlice(strs).Contains("apple"),true)


}