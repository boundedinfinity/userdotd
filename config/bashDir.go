package config

import (
	"os"
	"path"

	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	bashDirName        = "bash-dir"
	bashDirDescription = "The bash shell configuration directory"
)

var (
	bashDirDefault = path.Join(os.Getenv("HOME"), ".config", "bash")
)

func GetBashDir() string {
	return viper.GetString(bashDirName)
}

func ConfigureBashDir(flagset *flag.FlagSet) {
	stringPFlag(flagset, bashDirName, bashDirDefault, bashDirDescription)
}
