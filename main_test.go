package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumn(t *testing.T) {

	n := 1
	for {
		if n >= 18278 /* ZZZ */ {
			break
		}
		c, e := NewDefaultColumn(n)
		assert.Nil(t, e)
		t.Log(c.String())
		n = n + 1
	}
	/*
		col1, e := newColumn(1)
		assert.Nil(t, e)
		assert.Equal(t, col1.label, "A")

		col26, e := newColumn(26)
		assert.Nil(t, e)
		assert.Equal(t, col26.label, "Z")

		col27, e := newColumn(27)
		assert.Nil(t, e)
		assert.Equal(t, col27.label, "AA")

		col30, e := newColumn(30)
		assert.Nil(t, e)
		assert.Equal(t, col30.label, "AD")

		col38, e := newColumn(38)
		assert.Nil(t, e)
		assert.Equal(t, col38.label, "AL")

		col53, e := newColumn(53)
		assert.Nil(t, e)
		assert.Equal(t, col53.label, "BA")

		col79, e := newColumn(79)
		assert.Nil(t, e)
		assert.Equal(t, col79.label, "CA")

		col105, e := newColumn(105)
		assert.Nil(t, e)
		assert.Equal(t, col105.label, "DA")

		col676, e := newColumn(676)
		assert.Nil(t, e)
		assert.Equal(t, col676.label, "YZ")

		col677, e := newColumn(677)
		assert.Nil(t, e)
		assert.Equal(t, col677.label, "ZA")

		col702, e := newColumn(702)
		assert.Nil(t, e)
		assert.Equal(t, col702.label, "ZZ")

		col703, e := newColumn(703)
		assert.Nil(t, e)
		assert.Equal(t, col703.label, "AAA")

		col729, e := newColumn(729)
		assert.Nil(t, e)
		assert.Equal(t, col729.label, "ABA")

		col17576, e := newColumn(17576)
		assert.Nil(t, e)
		assert.Equal(t, col17576.label, "YYZ")
	*/
	/*
		col18252, e := newColumn(18252)
		assert.Nil(t, e)
		assert.Equal(t, col18252.label, "ZYZ")

		col18277, e := newColumn(18277)
		assert.Nil(t, e)
		assert.Equal(t, col18277.label, "ZZZ")
	*/
}
