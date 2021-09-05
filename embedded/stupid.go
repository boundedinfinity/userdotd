package embedded

import (
	"fmt"
	"io/fs"
	"path"
	"strings"

	"github.com/boundedinfinity/userdotd/model"
)

type stupidGoEmbedDirEntry struct {
	d fs.DirEntry
}

func stupidGoEmbedTrim(name string) string {
	comps1 := strings.Split(name, "/")
	comps2 := make([]string, 0)

	for _, comp := range comps1 {
		comps2 = append(comps2, strings.TrimPrefix(comp, model.StupidGoEmbed_HiddenPrefix))
	}

	name2 := path.Join(comps2...)
	return name2
}

func stupidGoEmbedAdd(name string) string {
	comps1 := strings.Split(name, "/")
	comps2 := make([]string, 0)

	for _, comp := range comps1 {
		if len(comp) > 1 && strings.HasPrefix(comp, ".") {
			comps2 = append(comps2, fmt.Sprintf("%v%v", model.StupidGoEmbed_HiddenPrefix, comp))
		} else {
			comps2 = append(comps2, comp)
		}
	}

	name2 := path.Join(comps2...)
	return name2
}

func (d *stupidGoEmbedDirEntry) Name() string               { return stupidGoEmbedTrim(d.d.Name()) }
func (d *stupidGoEmbedDirEntry) IsDir() bool                { return d.d.IsDir() }
func (d *stupidGoEmbedDirEntry) Type() fs.FileMode          { return d.d.Type() }
func (d *stupidGoEmbedDirEntry) Info() (fs.FileInfo, error) { return d.d.Info() }
