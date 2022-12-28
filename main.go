package main

import (
	"fmt"
)

// Fits A to Z inclusive
// ASCII Z-A-1 to fit A, 26 letters
var lenX = 90 - 65 + 1

// First AA to ZZ inclusive.
//       // AA-YZ    // AA-ZZ
var lenXX = lenX*lenX + lenX

// First AAA to ZZZ inclusive.
//        // AAA-YYZ       // ZYZ      // ZZZ
var lenXXX = lenX*lenX*lenX + lenX*lenX + lenX

func letterAt(i int) string {
	// i is 0-indexed so to get a letter, add i to 65 (A):
	return string([]byte{65 + byte(i)})
}

type Column interface {
	Index() int
	Label() string
	SetIndex(index int) (Column, error)
	SetLabel(label string) (Column, error)
	String() string
}

func NewColumnWintIndex(index int) (Column, error) {
	return (&defaultColumn{}).SetIndex(index)
}

func NewColumnWintLabel(label string) (Column, error) {
	return (&defaultColumn{}).SetLabel(label)
}

type defaultColumn struct {
	index int
	label string
}

func (c *defaultColumn) Index() int {
	return c.index
}

func (c *defaultColumn) Label() string {
	return c.label
}

func (c *defaultColumn) SetIndex(index int) (Column, error) {
	return c.updateWithIndex(index - 1)
}

func (c *defaultColumn) SetLabel(label string) (Column, error) {
	return c.updateWithLabel(label)
}

func (c *defaultColumn) String() string {
	return fmt.Sprintf("%s=index:%d", c.label, c.index+1)
}

func (c *defaultColumn) updateWithIndex(index int) (Column, error) {

	if index < lenX {
		c.index = index
		c.label = letterAt(index)
		return c, nil
	}
	if index < lenXX {
		w := int(index / lenX)
		r := index - w*lenX
		c.index = index
		c.label = letterAt(w-1) + letterAt(r)
		return c, nil
	}

	if index < lenXXX {
		w1 := int(index / lenX)
		r1 := index - w1*lenX
		w2 := w1 / lenX
		r2 := w1 - w2*lenX
		if r2 < 1 {
			// at AZ_, we need to correct our offsets
			// without correction, we get directly into BZ_ after AZY,
			// which isn't correct
			w2 = w2 - 1
			r2 = lenX
		}
		c.index = index
		c.label = letterAt(w2-1) + letterAt(r2-1) + letterAt(r1)
		return c, nil
	}

	return nil, ErrorIndexTooLarge(index + 1)
}

func (c *defaultColumn) updateWithLabel(label string) (Column, error) {
	bs := []byte(label)
	lbs := len(bs)
	for idx, b := range bs {
		if !isValidLetterByte(b) {
			return nil, ErrorLabelCharacterInvalid(idx, label)
		}
	}
	switch lbs {
	case 1:
		c.label = label
		c.index = int(bs[0] - 65)
	case 2:
		n1 := int(bs[0] - 65 + 1)
		n2 := int(bs[1] - 65)
		c.label = label
		c.index = n1*lenX + n2
	case 3:
		n1 := int(bs[0] - 65 + 1)
		n2 := int(bs[1] - 65 + 1)
		n3 := int(bs[2] - 65)
		c.label = label
		c.index = n1*lenX*lenX + n2*lenX + n3
	default:
		return nil, ErrorLabelLengthInvalid(label)
	}

	return c, nil
}

func isValidLetterByte(b byte) bool {
	return b >= 65 && b <= 90
}
