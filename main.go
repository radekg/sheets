package main

import "fmt"

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var lletters = len(letters)

func letterAt(i int) string {
	return letters[i : i+1]
}

type column struct {
	index      int
	n1, n2, n3 int
	label      string
}

func newColumn(index int) (*column, error) {
	return (&column{index: index - 1}).update()
}

func (c *column) update() (*column, error) {

	if c.index < lletters {
		c.label = letterAt(c.index)
		return c, nil
	}
	if c.index < lletters*lletters+lletters {
		w := int(c.index / lletters)
		r := c.index - w*lletters
		c.label = letterAt(w-1) + letterAt(r)
		return c, nil
	}

	//           // YYZ                     // ZYZ            // ZZZ
	if c.index < lletters*lletters*lletters+lletters*lletters+lletters {
		w1 := int(c.index / lletters)
		r1 := c.index - w1*lletters
		w2 := w1 / lletters
		r2 := w1 - w2*lletters
		if r2 < 1 {
			// at AZ_, we need to correct our offsets
			// without correction, we get directly into BZ_ after AZY,
			// which isn't correct
			w2 = w2 - 1
			r2 = lletters
		}
		c.label = letterAt(w2-1) + letterAt(r2-1) + letterAt(r1)
		return c, nil
	}

	return nil, fmt.Errorf("index %d too large", c.index+1)
}
