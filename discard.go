package runewriter

import (
	"unicode/utf8"
)

var Discard DiscardRuneWriter

type DiscardRuneWriter struct {}

var _ RuneWriter = DiscardRuneWriter{}

func (DiscardRuneWriter) WriteRune(r rune) (size int, err error) {
	size = utf8.RuneLen(r)

	return size, nil
}
