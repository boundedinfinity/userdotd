// Code generated by fileb0x at "2018-11-13 20:28:44.491642615 -0800 PST m=+0.003034749" from config file "fileb0x.yml" DO NOT EDIT.
// modification hash(e9863bad6b098215ead1c13b4090ecf1.3cee90bfbdc8800a832f28b0cd97aedc)

package system

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FileContentBashBashProfile is "content/bash/bash_profile"
var FileContentBashBashProfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

// FileContentBashBashrc is "content/bash/bashrc"
var FileContentBashBashrc = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

// FileContentFishConfigFish is "content/fish/config.fish"
var FileContentFishConfigFish = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\x4b\x0a\xc2\x30\x10\xc6\xf1\x7d\x4e\xf1\x41\xea\x46\x48\x73\x0a\xc1\x8d\x78\x04\xd1\x64\x52\x07\xca\x8c\xe4\x01\x1e\x5f\x52\x28\xf5\xb1\x31\x90\x10\xfe\x8b\x99\x5f\xd2\x0c\x06\x0b\x86\xe3\xf9\x74\xf0\x63\x50\x49\x3c\xf9\x56\x28\x47\xad\xd1\x27\x2e\x77\x3f\xeb\xc4\x32\x46\x4f\x72\xbd\xcd\x14\xfd\x7e\xec\xd9\x00\x00\x27\x54\x2a\x15\x2e\x63\xe0\xa5\xf4\x53\xb4\xe5\x40\x6b\x21\x89\xa6\x5f\x63\x91\x9a\x84\xca\x2a\x58\x17\x5c\xe8\xc9\x15\xce\xa9\xb8\x47\xd6\x40\xa5\xb8\xa5\xec\x0a\xcd\xc9\xd8\x65\xda\xbf\x46\x6d\xf5\x17\x69\x57\xd3\x17\x75\xeb\x9f\xe0\xad\x77\xb2\x7d\xfb\xf5\xf7\x15\x00\x00\xff\xff\x2c\xb0\x0b\x94\x2f\x01\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	err = FS.Mkdir(CTX, "content/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "content/bash/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "content/fish/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileContentBashBashProfile)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "content/bash/bash_profile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileContentBashBashrc)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "content/bash/bashrc", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileContentFishConfigFish)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "content/fish/config.fish", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}