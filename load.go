package vocabtxt

import (
	"os"
	"unsafe"

	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"golang.org/x/sys/unix"
)

const maxKey = 1073741824

// LoadFromBytes loads (WordPiece) "vocab.txt" data from a []byte into a map[string]uint.
func LoadFromBytes(destination *map[string]uint, bytes []byte) error {
	if nil == destination {
		return ErrMapNil
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

// LoadFromBytes loads (WordPiece) "vocab.txt" data from a file into a map[string]uint.
func LoadFromFile(destination *map[string]uint, filename string) error {
	file, err := os.Open(filename)
	if nil != err {
		err = erorr.Wrap(err, "failed to open file",
			field.String("file-name", filename),
		)
		return err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if nil != err {
		err = erorr.Wrap(err, "failed to stat file",
			field.String("file-name", filename),
		)
		return err
	}

	size := fileinfo.Size()
	if size <= 0 {
		return nil
	}

	bytes, err := unix.Mmap(int(file.Fd()), 0, int(size), unix.PROT_READ, unix.MAP_SHARED)
	if nil != err {
		err = erorr.Wrap(err, "failed to mmap file",
			field.String("file-name", filename),
		)
		return err
	}
	defer unix.Munmap(bytes)

	return LoadFromBytes(destination, bytes)
}
