package system

import (
	"errors"
	"io/fs"
	"os"
	"path"

	"github.com/boundedinfinity/userdotd/embedded"
	"github.com/boundedinfinity/userdotd/model"
	"github.com/boundedinfinity/userdotd/pathutil"
)

func (t *System) ShellInitialize(name string) (model.ShellInitialize, error) {
	init := model.ShellInitialize{
		Name:  name,
		Files: make([]model.ShellFile, 0),
	}

	home, err := os.UserHomeDir()

	if err != nil {
		return init, err
	}

	err = embedded.WalkShell(name, func(p string, d fs.DirEntry, err error) error {
		p2 := p
		p2 = embedded.TrimPathPrefix(p2, "shell", name)
		p2 = path.Join(home, p2)

		if d.IsDir() {
			if err := pathutil.EnsureDir(p2); err != nil {
				return err
			}
		} else {

		}

		return nil
	})

	if err != nil && !errors.Is(err, model.ErrEndWalk) {
		return init, err
	}

	return init, nil
}

func (t *System) ShellStatuses(names ...string) ([]model.ShellStatus, error) {
	statuses := make([]model.ShellStatus, 0)

	for _, name := range names {
		status, err := t.ShellStatus(name)

		if err != nil {
			return statuses, err
		}

		statuses = append(statuses, status)
	}

	return statuses, nil
}

func (t *System) ShellStatus(name string) (model.ShellStatus, error) {
	status := model.ShellStatus{
		Name:  name,
		State: model.ShellState_Unknown,
	}

	name2 := path.Join("shell", name)
	home, err := os.UserHomeDir()

	if err != nil {
		return status, err
	}

	err = embedded.WalkDir(name2, func(p string, d fs.DirEntry, err error) error {
		p2 := p
		p2 = embedded.TrimPathPrefix(p2, "shell", name)
		p2 = path.Join(home, p2)

		if !pathutil.Exists(p2) {
			status.State = model.ShellState_Uninitialized
			return model.ErrEndWalk
		}

		if !d.IsDir() {

		}

		return nil
	})

	if err != nil && !errors.Is(err, model.ErrEndWalk) {
		return status, err
	}

	return status, nil
}
