package pato

import (
	"bytes"
	"errors"
	"io"
	"strconv"
)

// LineCol represents a source location with file name, line, and column.
type LineCol struct {
	Source    string
	Line, Col int
}

// String returns the location formatted as "source:line:col".
func (lc LineCol) String() string {
	return string(lc.AppendString(nil))
}

// AppendString appends the formatted location to b and returns the result.
func (lc LineCol) AppendString(b []byte) []byte {
	if b == nil {
		b = make([]byte, 0, len(lc.Source)+3+3)
	}
	b = append(b, lc.Source...)
	if lc.Line == 0 {
		return b
	}
	b = append(b, ':')
	b = strconv.AppendInt(b, int64(lc.Line), 10)
	if lc.Col > 0 {
		b = append(b, ':')
		b = strconv.AppendInt(b, int64(lc.Col), 10)
	}
	return b
}

// Pos represents a byte offset in the source.
type Pos int

// ToLineCol converts a byte offset to line:column.
// Line and column are 1-indexed. Also returns the length of the line containing the offset.
// aux is a scratch buffer used for reading; its size determines read chunk size (1024B recommended).
func (pos Pos) ToLineCol(r io.ReaderAt, aux []byte) (line, col, lineLength int, err error) {
	offset := int(pos)
	if r == nil || offset < 0 {
		return 0, 0, 0, errors.New("invalid reader or offset")
	}

	line = 1
	lastNewlinePos := -1 // byte position of last newline seen (-1 means before start of file)

	// Read source up to offset to count newlines
	for readPos := 0; readPos < offset; {
		toRead := min(len(aux), offset-readPos)
		n, rerr := r.ReadAt(aux[:toRead], int64(readPos))
		if n == 0 && rerr != nil {
			return 0, 0, 0, rerr
		}

		// Count newlines and find last newline position in this chunk
		chunk := aux[:n]
		for {
			idx := bytes.IndexByte(chunk, '\n')
			if idx < 0 {
				break
			}
			line++
			lastNewlinePos = readPos + (n - len(chunk)) + idx
			chunk = chunk[idx+1:]
		}

		readPos += n
		if rerr == io.EOF {
			break
		}
	}

	col = offset - lastNewlinePos

	// Find line length by reading until next newline or EOF
	lineLength = col - 1 // at minimum, the portion before offset
	for readPos := offset; ; {
		n, rerr := r.ReadAt(aux[:], int64(readPos))
		if n == 0 && rerr != nil {
			break
		}
		idx := bytes.IndexByte(aux[:n], '\n')
		if idx >= 0 {
			lineLength = (readPos - lastNewlinePos - 1) + idx
			break
		}
		readPos += n
		if rerr == io.EOF {
			lineLength = readPos - lastNewlinePos - 1
			break
		}
	}

	return line, col, lineLength, nil
}
