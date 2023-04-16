package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"

	"golang.org/x/example/stringutil"
)

var ErrInvalidString = errors.New("invalid string")

// isLetterDigit wraps strconv.Atoi to be more usable in conditionals for checking is rune digit.
func isRuneDigit(letter rune) bool {
	_, err := strconv.Atoi(string(letter))
	return err == nil
}

// Unpack unpacks strings according to task setting.
func Unpack(s string) (string, error) {
	var b strings.Builder
	r := []rune(s) // Because unicode runes may occur in the input, not just bytes.
	if len(r) == 0 {
		return "", nil
	}
	if isRuneDigit(r[0]) {
		return "", ErrInvalidString
	}
	if len(r) == 1 {
		return string(r), nil
	}
	// Go from the end to the beginig.
	// If r[i] and r[i-1] are both digits - return error.
	// If r[i-1] is digit - repeat the rune r[i] as many times as necessary and jump over it.
	// Else - write r[i].
	for i := len(r) - 1; i >= 1; i-- {
		if isRuneDigit(r[i]) && isRuneDigit(r[i-1]) {
			return "", ErrInvalidString
		}
		if num, err := strconv.Atoi(string(r[i])); err == nil {
			b.WriteString(strings.Repeat(string(r[i-1]), num))
			i--
		} else {
			b.WriteRune(r[i])
		}
	}
	// During the reverse loop, we didn't write the last character if before that we did not meet a digit. Fixing it.
	if !isRuneDigit(r[1]) {
		b.WriteRune(r[0])
	}
	return stringutil.Reverse(b.String()), nil // Our b has reverse order. Reverse it again to satisfy the task statement.
}
