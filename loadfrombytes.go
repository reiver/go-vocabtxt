package vocabtxt

import (
	"unsafe"
)

const maxKey = 1048576

// LoadFromBytes loads (WordPiece) "vocab.txt" data from a []byte into a map[string]uint.
func LoadFromBytes(destination *map[string]uint, bytes []byte) error {
	if nil == destination {
		return ErrMapNil
	}
	if nil == *destination {
		*destination = map[string]uint{}
	}
	if 0 < len(*destination) {
		return ErrMapNotEmpty
	}

	if len(bytes) <= 0 {
		return nil
	}

	// The keyBuffer length is chosen to try to be a multiple of a cache line
	var keyBuffer [128]rune
	var key []rune = keyBuffer[0:0]

	// The line number of the vocab.txt data, starting with zero (0).
	// I.e., zero-indexed (rather than one-indexed).
	var index uint

	var str string = unsafe.String(&bytes[0], len(bytes))
	for _, r := range str {
		if '\n' == r {
			if 0 < len(key) {
				if '\r' == key[len(key)-1] {
					key = key[:len(key)-1]
				}
			}
			(*destination)[string(key)] = index
			key = keyBuffer[0:0]
			index++
			continue
		}

		key = append(key, r)
		if maxKey < len(key) {
			return ErrTokenTooBig
		}
	}
	if 0 < len(key) {
		(*destination)[string(key)] = index
	}

	return nil
}
