package vocabtxt

import (
	"unicode"
)

// isSpacing matches what Python's .strip() / str.isspace() consider spacing.
func isSpacing(r rune) bool {
	return unicode.IsSpace(r) || '\u001C' == r || '\u001D' == r || '\u001E' == r || '\u001F' == r
}
