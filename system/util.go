package system

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/userdotd/config"
)

func ensureFile(fd fileDescriptor) error {
	if fileExists(fd.Dest) {
		return nil
	}

	if err := ensureDir(filepath.Dir(fd.Dest)); err != nil {
		return err
	}

	if config.GetDebug() {
		log.Printf("Copy:\n")
		log.Printf("  Src: %v\n", fd.Src)
		log.Printf(" Dest: %v\n", fd.Dest)
	}

	// content, err := ioutil.ReadFile(fd.Src)
	content, err := ReadFile(fd.Src)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fd.Dest, content, fd.Perm)

	if err != nil {
		return err
	}

	return nil
}

func ensureLink(fd fileDescriptor) error {
	if fileExists(fd.Link) {
		return nil
	}

	if err := ensureDir(filepath.Dir(fd.Link)); err != nil {
		return err
	}

	if config.GetDebug() {
		log.Printf("Create:\n")
		log.Printf(" Link: %v\n", fd.Link)
		log.Printf(" Dest: %v\n", fd.Dest)
	}

	if err := os.Symlink(fd.Dest, fd.Link); err != nil {
		return err
	}

	return nil
}

func removeLink(fd fileDescriptor) error {
	if !fileExists(fd.Link) {
		return nil
	}

	if config.GetDebug() {
		log.Printf("Remove:\n")
		log.Printf("   Link: %v\n", fd.Link)
	}

	if err := os.Remove(fd.Link); err != nil {
		return err
	}

	return nil
}

func dirExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func fileExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func ensureDir(elem ...string) error {
	fullPath := path.Join(elem...)

	if !dirExists(fullPath) {

		if config.GetDebug() {
			log.Printf("Creating %v\n", fullPath)
		}

		if err := os.MkdirAll(fullPath, 0700); err != nil {
			return err
		}
	}

	return nil
}

func getNames(p string) ([]string, error) {
	names := make([]string, 0)

	if !dirExists(p) {
		return names, nil
	}

	files, err := ioutil.ReadDir(p)

	if err != nil {
		return names, err
	}

	for _, f := range files {
		names = append(names, path2name(f.Name()))
	}

	return names, nil
}

func hasName(names []string, name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}

	return false
}

func path2name(p string) string {
	b := path.Base(p)
	n := strings.TrimSuffix(b, path.Ext(b))
	return n
}
