package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const zzzLabelIndex = 18278

func TestColumn(t *testing.T) {

	n := 1
	for {
		if n > zzzLabelIndex {
			break
		}
		c1, e := NewColumnWintIndex(n)
		assert.Nil(t, e)

		c2, e := NewColumnWintLabel(c1.Label())
		assert.Nil(t, e)
		assert.Equal(t, c1.Index(), c2.Index())

		n = n + 1
	}
}

func TestLabelErrors(t *testing.T) {
	c, e := NewColumnWintLabel("")
	assert.Nil(t, c)
	assert.True(t, IsErrorLabelLengthInvalid(e))

	c, e = NewColumnWintLabel("AAAA")
	assert.Nil(t, c)
	assert.True(t, IsErrorLabelLengthInvalid(e))

	c, e = NewColumnWintLabel("AB{")
	assert.Nil(t, c)
	assert.True(t, IsErrorLabelCharacterInvalid(e))

	c, e = NewColumnWintIndex(zzzLabelIndex + 1)
	assert.Nil(t, c)
	assert.True(t, IsErrorIndexTooLarge(e))
}
