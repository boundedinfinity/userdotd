package model

import (
	"errors"
	"fmt"
)

var (
	ErrEndWalk        = errors.New("end FS Walk")
	ErrExistButNotDir = errors.New("exists but not a directory")
	ErrFileNotEqual   = errors.New("file not equal")
)

func ErrExistButNotDirNew(p string) error {
	return fmt.Errorf("%v : %w", p, ErrExistButNotDir)
}

func ErrFileNotEqualNew(p string) error {
	return fmt.Errorf("%v : %w", p, ErrFileNotEqual)
}
