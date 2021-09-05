package embedded

import (
	"embed"
	"io/fs"
	"path"
	"strings"

	"github.com/boundedinfinity/userdotd/model"
)

//go:embed content/*
var files embed.FS

func contentRootTrim(n string) string {
	return TrimPathPrefix(n, model.Embedded_Root)
}

func contentRootAdd(n string) string {
	n2 := n
	n2 = path.Join(model.Embedded_Root, n2)
	return n2
}

func TrimPathPrefix(p string, elem ...string) string {
	elems := path.Join(elem...)
	p2 := p
	p2 = strings.TrimPrefix(p2, elems)
	p2 = strings.TrimPrefix(p2, "/")
	return p2
}

func ReadDir(name string) ([]fs.DirEntry, error) {
	name2 := stupidGoEmbedAdd(name)
	name2 = contentRootAdd(name2)
	entries, err := files.ReadDir(name2)

	if err != nil {
		return entries, nil
	}

	entires2 := make([]fs.DirEntry, 0)

	for _, entry := range entries {
		entires2 = append(entires2, &stupidGoEmbedDirEntry{entry})
	}

	return entires2, nil
}

func ReadFile(name string) ([]byte, error) {
	name2 := stupidGoEmbedAdd(name)
	name2 = contentRootAdd(name2)
	return files.ReadFile(name2)
}

func OpenFile(name string) (fs.File, error) {
	name2 := stupidGoEmbedAdd(name)
	name2 = contentRootAdd(name2)
	return files.Open(name2)
}

func WalkDir(root string, fn fs.WalkDirFunc) error {
	root2 := stupidGoEmbedAdd(root)
	root2 = path.Join(model.Embedded_Root, root2)

	return fs.WalkDir(files, root2, func(p string, d fs.DirEntry, err error) error {
		p2 := stupidGoEmbedTrim(p)
		p2 = contentRootTrim(p2)

		return fn(p2, &stupidGoEmbedDirEntry{d}, err)
	})
}

func WalkShell(name string, fn fs.WalkDirFunc) error {
	name2 := path.Join("shell", name)
	return WalkDir(name2, fn)
}
