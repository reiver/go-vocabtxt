package vocabtxt

import (
	"os"

	"codeberg.org/reiver/go-erorr"
	"codeberg.org/reiver/go-field"
	"golang.org/x/sys/unix"
)

// LoadFromFile loads (WordPiece) "vocab.txt" data from a file into a map[string]uint.
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
