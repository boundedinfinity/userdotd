package embedded

import (
	"embed"
)

//go:embed bash/*
//go:embed fish/*
//go:embed zsh/*
//go:embed user.d/*
var EmbeddedFiles embed.FS
