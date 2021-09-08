package pathutil

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"

	"github.com/boundedinfinity/userdotd/model"
	"github.com/udhos/equalfile"
)

func WalkDir(root string, fn fs.WalkDirFunc, pathFilters []PathFilterFunc, dirFilters []DirFilterFunc) error {
	path, fsys, err := getFs(root)

	if err != nil {
		return err
	}

	wrapper := func(p string, d fs.DirEntry, err error) error {
		if p == "." {
			return nil
		}

		if d.Name() == "." {
			return nil
		}

		if DirShouldFilter(d, dirFilters...) {
			return nil
		}

		if PathShouldFilter(p, pathFilters...) {
			return nil
		}

		return fn(p, d, err)
	}

	return fs.WalkDir(fsys, path, wrapper)
}

func Exists(p string) (bool, error) {
	path, fsys, err := getFs(p)

	if err != nil {
		return false, err
	}

	_, err = fs.Stat(fsys, path)

	if err != nil && err != fs.ErrNotExist {
		return false, err
	}

	return !os.IsNotExist(err), nil
}

func EnsureDir(p string) error {
	if IsEmbeddedPath(p) {
		return fmt.Errorf("can't ensure embedded directory: %v", p)
	}

	path, fsys, err := getFs(p)

	if err != nil {
		return err
	}

	fi, err := fs.Stat(fsys, path)

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
	if IsEmbeddedPath(p) {
		return fmt.Errorf("can't backup embedded directory: %v", p)
	}

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

func Equal(path1, path2 string) (bool, error) {
	file1, err := open(path1)

	if err != nil {
		return false, err
	}

	defer file1.Close()

	file2, err := open(path2)

	if err != nil {
		return false, err
	}

	defer file2.Close()

	eq := equalfile.New(nil, equalfile.Options{})
	isEq, err := eq.CompareReader(file1, file2)

	return isEq, err
}

func Copy(from, to string) error {
	if IsEmbeddedPath(to) {
		return fmt.Errorf("can't copy to embedded path: %v", to)
	}

	fromFile, err := open(from)

	if err != nil {
		return err
	}

	defer fromFile.Close()

	tofile, err := os.Create(to)

	if err != nil {
		return err
	}

	defer tofile.Close()

	_, err = io.Copy(tofile, fromFile)

	return err
}
