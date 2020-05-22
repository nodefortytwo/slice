package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString (t *testing.T){
	var strs = []string{"peach", "apple", "pear", "plum"}


	assert.Equal(t, String(strs).Index("apple"), 1)
	assert.Equal(t, String(strs).Index("bogus"), -1)

	assert.Equal(t, String(strs).Contains("bogus"),false)
	assert.Equal(t, String(strs).Contains("apple"),true)


}