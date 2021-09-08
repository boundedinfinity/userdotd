package pathutil

import (
	"io/fs"
	"strings"

	"github.com/boundedinfinity/userdotd/model"
)

type DirFilterFunc func(fs.DirEntry) bool

var (
	DirFilterDot = func(entry fs.DirEntry) bool {
		return entry.Name() == "."
	}

	DirFilterDir = func(entry fs.DirEntry) bool {
		return entry.IsDir()
	}

	DirFilterFile = func(entry fs.DirEntry) bool {
		return !DirFilterDir(entry)
	}

	DirFilterKeepMarker = func(entry fs.DirEntry) bool {
		return entry.Name() == model.StupidGoEmbed_KeepFile
	}
)

func DirShouldFilter(entry fs.DirEntry, fns ...DirFilterFunc) bool {
	for _, fn := range fns {
		if fn(entry) {
			return true
		}
	}

	return false
}

type PathFilterFunc func(string) bool

var (
	PathFilterEmpty = func(p string) bool {
		return p == ""
	}

	PathFilterDot = func(p string) bool {
		return p == "."
	}

	PathFilterKeepMarker = func(p string) bool {
		return strings.HasSuffix(p, model.StupidGoEmbed_KeepFile)
	}
)

func PathShouldFilter(p string, fns ...PathFilterFunc) bool {
	for _, fn := range fns {
		if fn(p) {
			return true
		}
	}

	return false
}
