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

func (t *System) ShellEmbeddedList(request model.ShellEmbeddedListRequest) (model.ShellEmbeddedListResponse, error) {
	response := model.ShellEmbeddedListResponse{}

	embedded.WalkDirRaw(".", func(p string, d fs.DirEntry, err error) error {
		if p == "." || p == model.StupidGoEmbed_KeepFile {
			return nil
		}

		response.Files = append(response.Files, p)
		return nil
	})

	return response, nil
}

func (t *System) ShellInitialize(request model.ShellInitializeRequest) (model.ShellInitializeResponse, error) {
	response := model.ShellInitializeResponse{
		Name:  request.Name,
		Files: make([]model.ShellFile, 0),
	}

	home, err := os.UserHomeDir()

	if err != nil {
		return response, err
	}

	err = embedded.WalkShell(request.Name, func(ep string, d fs.DirEntry, err error) error {
		if d.Name() == model.StupidGoEmbed_KeepFile {
			return nil
		}

		rp := ep
		rp = embedded.TrimPathPrefix(rp, "shell", request.Name)
		rp = path.Join(home, rp)

		if d.IsDir() {
			if err := pathutil.EnsureDir(rp); err != nil {
				return err
			}
		} else {
			if pathutil.Exists(rp) {
				eq, err := pathutil.EmbeddedEqual(ep, rp)

				if err != nil {
					return err
				}

				if eq {
					return nil
				} else {
					if request.Force {
						if err := pathutil.BackupFile(rp); err != nil {
							return err
						}

						if err := pathutil.EmbeddedCopy(ep, rp); err != nil {
							return err
						}
					} else {
						return model.ErrFileNotEqualNew(rp)
					}
				}
			} else {
				if err := pathutil.EmbeddedCopy(ep, rp); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil && !errors.Is(err, model.ErrEndWalk) {
		return response, err
	}

	return response, nil
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

	home, err := os.UserHomeDir()

	if err != nil {
		return status, err
	}

	err = embedded.WalkShell(name, func(p string, d fs.DirEntry, err error) error {
		p2 := p
		p2 = embedded.TrimPathPrefix(p2, "shell", name)
		p2 = path.Join(home, p2)

		if !pathutil.Exists(p2) {
			status.State = model.ShellState_Uninitialized
			return model.ErrEndWalk
		}

		return nil
	})

	if err != nil && !errors.Is(err, model.ErrEndWalk) {
		return status, err
	}

	return status, nil
}
