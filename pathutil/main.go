package pathutil

import (
	"os"

	"github.com/boundedinfinity/userdotd/model"
)

func Exists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

func EnsureDir(p string) error {
	fi, err := os.Stat(p)

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		if err := os.MkdirAll(p, 0755); err != nil {
			return err
		}
	} else {
		if !fi.IsDir() {
			return model.ErrExistButNotDirNew(p)
		}
	}

	return nil
}
