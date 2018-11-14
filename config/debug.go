package config

import (
	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	debugName        = "debug"
	debugDefault     = false
	debugDescription = "Debugging mode"
)

func GetDebug() bool {
	return viper.GetBool(debugName)
}

func ConfigureDebug(flagset *flag.FlagSet) {
	boolPFlag(flagset, debugName, debugDefault, debugDescription)
}
