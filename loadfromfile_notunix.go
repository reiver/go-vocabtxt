//go:build !unix

package vocabtxt

import (
	"os"

	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
)

// LoadFromFile loads (WordPiece) "vocab.txt" data from a file into a map[string]uint.
func LoadFromFile(destination *map[string]uint, filename string) error {
	bytes, err := os.ReadFile(filename)
	if nil != err {
		err = erorr.Wrap(err, "failed to read file",
			field.String("file-name", filename),
		)
		return err
	}

	return LoadFromBytes(destination, bytes)
}
