package config

import (
	"os"
	"path"

	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	userDirName        = "user-dir"
	userDirDescription = "The user directory"
)

var (
	userDirDefault = path.Join(os.Getenv("HOME"), ".config")
)

func GetUserDir() string {
	return viper.GetString(debugName)
}

func ConfigureUserDir(flagset *flag.FlagSet) {
	stringPFlag(flagset, userDirName, userDirDefault, userDirDescription)
}
