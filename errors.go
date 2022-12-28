package main

import (
	"fmt"
	"strings"
)

func ErrorIndexTooLarge(index int) error {
	return fmt.Errorf("error: index too large: %d", index)
}

func IsErrorIndexTooLarge(e error) bool {
	return strings.HasPrefix(e.Error(), "error: index too large:")
}

func ErrorLabelLengthInvalid(input string) error {
	return fmt.Errorf("error: label invalid, must be 1 to 3 characters long: %q", input)
}

func IsErrorLabelLengthInvalid(e error) bool {
	return strings.HasPrefix(e.Error(), "error: label invalid, must be 1 to 3 characters long:")
}

func ErrorLabelCharacterInvalid(position int, input string) error {
	return fmt.Errorf("error: label invalid, character at position: %d not in range A-Z: %q", position, input)
}

func IsErrorLabelCharacterInvalid(e error) bool {
	return strings.HasPrefix(e.Error(), "error: label invalid, character at position:")
}
