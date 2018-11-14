package system

import (
	"errors"
	"os"

	"github.com/boundedinfinity/userdotd/config"
)

type fileDescriptor struct {
	Name string
	Src  string
	Dest string
	Link string
	Perm os.FileMode
}

var fishDescriptors []fileDescriptor

func getFishDescriptors() []fileDescriptor {
	if fishDescriptors == nil {
		fishDescriptors = []fileDescriptor{
			fileDescriptor{
				Src:  config.GetContentPath("fish/config.fish"),
				Dest: config.GetFishPath("config.fish"),
				Perm: os.FileMode(int(0644)),
			},

			fileDescriptor{
				Name: "home-bin",
				Src:  config.GetContentPath("userdotd/fish/login.d/available/home-bin.fish"),
				Dest: config.GetFishLoginAvailablePath("home-bin.fish"),
				Link: config.GetBashLoginEnabledPath("home-bin.fish"),
				Perm: os.FileMode(int(0600)),
			},

			fileDescriptor{
				Name: "go-path",
				Src:  config.GetContentPath("userdotd/fish/login.d/available/go-path.fish"),
				Dest: config.GetFishLoginAvailablePath("go-path.fish"),
				Link: config.GetBashLoginEnabledPath("go-path.fish"),
				Perm: os.FileMode(int(0600)),
			},
		}
	}

	return fishDescriptors
}

var (
	ErrorDescriptorNotFound = errors.New("descriptor not found")
)

func lookupDescriptor(name string) (fileDescriptor, error) {
	var ds []fileDescriptor

	if config.IsFishShell() {
		ds = getFishDescriptors()
	}

	if ds == nil {
		return fileDescriptor{}, ErrorDescriptorNotFound
	}

	for _, d := range ds {
		if d.Name == name {
			return d, nil
		}
	}

	return fileDescriptor{}, ErrorDescriptorNotFound
}
