package model

import (
	"os"
)

var (
	Shells = []string{
		"bash",
		"fish",
		"zsh",
	}
)

type ShellStatus struct {
	Name  string     `json:"name"`
	State ShellState `json:"state"`
}

type ShellInitialize struct {
	Name  string      `json:"name"`
	Files []ShellFile `json:"files"`
}

type ShellFile struct {
	Source      string      `json:"source"`
	Destination string      `json:"destination"`
	Perm        os.FileMode `json:"perm"`
}
