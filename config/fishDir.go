package config

import (
	"os"
	"path"

	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	fishDirName        = "fish-dir"
	fishDirDescription = "The fish shell configuration directory"
)

var (
	fishDirDefault = path.Join(os.Getenv("HOME"), ".config", "fish")
)

func GetFishDir() string {
	return viper.GetString(fishDirName)
}

func ConfigureFishDir(flagset *flag.FlagSet) {
	stringPFlag(flagset, fishDirName, fishDirDefault, fishDirDescription)
}
