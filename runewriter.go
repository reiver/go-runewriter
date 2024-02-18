package runewriter

// RuneWriter is the interface that wraps the WriteRune method.
//
// Why doesn't this interface exist in the Go built-in package "io"?!?!?!?
// It has io.RuneReader and io.ByteWriter. Why not io.RuneWriter?!?!?!?
// Plus the WriteRune method exists on bufio.Writer, bytes.Buffer, and strings.Builder
type RuneWriter interface {
	WriteRune(r rune) (size int, err error)
}
