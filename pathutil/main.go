package pathutil

import (
	"fmt"
	"io"
	"os"
	"time"

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

func BackupFile(p string) error {
	cf, err := os.Open(p)

	if err != nil {
		return err
	}

	defer cf.Close()

	now := time.Now()
	bp := p
	bp = fmt.Sprintf("%v.%v", bp, model.BackupFile_Prefix)
	bp = fmt.Sprintf("%v-%v", bp, now.Format(time.RFC3339))

	bf, err := os.Create(bp)

	if err != nil {
		return err
	}

	defer bf.Close()

	_, err = io.Copy(bf, cf)

	if err != nil {
		return err
	}

	return nil
}
