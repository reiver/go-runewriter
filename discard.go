package runewriter

import (
	"unicode/utf8"
)

type Discard struct {}

var _ RuneWriter = Discard{}

func (Discard) WriteRune(r rune) (size int, err error) {
	size = utf8.RuneLen(r)

	return size, nil
}
