package csv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWord(t *testing.T) {
	var input = "1,2,3,4,abc,,a,"

	var golden = []string{"1", "2", "3", "4", "abc", "", "a", ""}

	var target = GetWords(input)

	assert.Equal(t, golden, target)
}

func TestGetWordCase2(t *testing.T) {
	var input = `1,2,3,4,"abc,def",,a,,`

	var golden = []string{"1", "2", "3", "4", "abc,def", "", "a", "", ""}

	var target = GetWords(input)

	assert.Equal(t, golden, target)
}

func TestGetWordCase3(t *testing.T) {
	var input = `1,2,3,4,"abc"",
def",,a`

	var golden = []string{"1", "2", "3", "4", "abc\",\ndef", "", "a"}

	var target = GetWords(input)

	assert.Equal(t, golden, target)
}

func TestGetWordCase4(t *testing.T) {
	var input = `1,2,3,4,"abc"`

	var golden = []string{"1", "2", "3", "4", "abc"}

	var target = GetWords(input)

	assert.Equal(t, golden, target)
}

func TestGetWordCase5(t *testing.T) {
	var input = `1,2,3,4,abc,,a
b,c`

	var golden = []string{"1", "2", "3", "4", "abc", "", "a\nb", "c"}

	var target = GetWords(input)

	assert.Equal(t, golden, target)
}

func TestGetWordCase6(t *testing.T) {
	var input = `,`

	var golden = []string{"", ""}

	var target = GetWords(input)

	assert.Equal(t, golden, target)
}

func TestGetWordCase7(t *testing.T) {
	var input = ``

	var golden = []string{}

	var target = GetWords(input)

	assert.Equal(t, golden, target)
}
