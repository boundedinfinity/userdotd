package config

import (
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	ENVIRONMENT_VARIABLE_PREFIX = "userdotd"
	CONFIG_FILENAME             = "userdotd"
)

func InitConfigEnvironment() {
	viper.SetConfigFile(CONFIG_FILENAME)
	viper.AddConfigPath("$HOME/.config/userdotd")
	viper.SetEnvPrefix(ENVIRONMENT_VARIABLE_PREFIX)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}

func intPFlag(flagSet *flag.FlagSet, name string, defaultValue int, description string) {
	flagSet.Int(name, defaultValue, description)
	viper.BindPFlag(name, flagSet.Lookup(name))
}

func stringPFlag(flagSet *flag.FlagSet, name string, defaultValue string, description string) {
	flagSet.String(name, defaultValue, description)
	viper.BindPFlag(name, flagSet.Lookup(name))
}

func boolPFlag(flagSet *flag.FlagSet, name string, defaultValue bool, description string) {
	flagSet.Bool(name, defaultValue, description)
	viper.BindPFlag(name, flagSet.Lookup(name))
}
