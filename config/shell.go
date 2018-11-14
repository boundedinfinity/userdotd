package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	shellName = "shell"
)

var (
	shellDefault     = filepath.Base(os.Getenv("SHELL"))
	validShells      = []string{"fish", "bash"}
	shellDescription = fmt.Sprintf("The shell type. Can be one of: %v", strings.Join(validShells, ", "))
)

func GetShell() string {
	return viper.GetString(shellName)
}

func ConfigureShell(flagset *flag.FlagSet) {
	stringPFlag(flagset, shellName, shellDefault, shellDescription)
}

func IsFishShell() bool {
	return GetShell() == "fish"
}

func IsBashShell() bool {
	return GetShell() == "bash"
}
