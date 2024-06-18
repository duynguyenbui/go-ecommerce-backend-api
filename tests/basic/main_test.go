package basic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddOne(t *testing.T) {
	assert.Equal(t, AddOne(2), 3, "AddOne(2) should equal 3")
}

func TestAddOne2(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne2(input)

	if actual != output {
		t.Errorf("AddOne2(%d) should equal %d", input, output)
	}
}
