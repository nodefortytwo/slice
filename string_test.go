package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString (t *testing.T){
	var strs = []string{"peach", "apple", "pear", "plum"}


	assert.Equal(t, String(strs).Index("apple"), 1)
	assert.Equal(t, String(strs).Index("bogus"), -1)

	assert.Equal(t, String(strs).Includes("bogus"),false)
	assert.Equal(t, String(strs).Includes("apple"),true)

	assert.Equal(t, String(strs).Next(2),3)
	assert.Equal(t, String(strs).Next(3),0)
	assert.Equal(t, String(strs).Next(0),1)


	assert.Equal(t, String(strs).NextValue(0),"apple")
	assert.Equal(t, String(strs).NextValue(3),"peach")

	assert.Equal(t, String(strs).Any(isPeach),true)
	assert.Equal(t, String(strs).Filter(isPeach), String{"peach"})
	assert.Equal(t, String(strs).Map(makePeach)[0], "peach")


	strings := String{"awesome"}
	assert.Equal(t,strings.Slice(), []string{"awesome"})

}

func isPeach(s string) bool{
	if s == "peach" {
		return true
	}

	return false
}

func makePeach(s string) string {
	return "peach"
}