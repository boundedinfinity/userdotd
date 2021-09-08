package pathutil

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/userdotd/embedded"
)

func TrimFirstComp(s string) string {
	ss := strings.Split(s, string(filepath.Separator))
	ss = ss[1:]
	s = filepath.Join(ss...)
	return s
}

func GetFilePath(s string) string {
	if s == "" {
		s = "."
	}

	return fmt.Sprintf("file://%v", s)
}

func GetEmbeddedPath(s string) string {
	if s == "" {
		s = "."
	}

	return fmt.Sprintf("file://%v", s)
}

func IsFilePath(s string) bool {
	return strings.HasPrefix(s, "file://")
}

func IsEmbeddedPath(s string) bool {
	return strings.HasPrefix(s, "embedded://")
}

func trimPrefix(s string) string {
	o := s
	o = strings.TrimPrefix(o, "file://")
	o = strings.TrimPrefix(o, "embedded://")

	return o
}

func getFs(path string) (string, fs.FS, error) {
	if IsFilePath(path) {
		root, err := os.UserHomeDir()

		if err != nil {
			return "", nil, err

		}

		return trimPrefix(path), os.DirFS(root), nil
	} else if IsEmbeddedPath(path) {
		fsys, err := fs.Sub(embedded.EmbeddedFiles, ".")

		return trimPrefix(path), fsys, err
	} else {
		return "", nil, fmt.Errorf("invalid path: %v", path)
	}
}

func open(name string) (fs.File, error) {
	path, fsys, err := getFs(name)

	if err != nil {
		return nil, err
	}

	file, err := fsys.Open(path)

	if err != nil {
		return nil, err
	}

	return file, nil
}
