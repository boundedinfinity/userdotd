package config

import (
	"os"
	"path"

	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	userdotdDirName        = "userdotd-dir"
	userdotdDirDescription = "The userdotd configuration directory"
)

var (
	userdotdDirDefault = path.Join(os.Getenv("HOME"), ".config", "userdotd")
)

func GetuserdotdDir() string {
	return viper.GetString(userdotdDirName)
}

func ConfigureuserdotdDir(flagset *flag.FlagSet) {
	stringPFlag(flagset, userdotdDirName, userdotdDirDefault, userdotdDirDescription)
}
