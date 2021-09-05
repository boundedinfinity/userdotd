package cmd

import (
	"fmt"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

var (
	forceName        = "force"
	forceDefault     = false
	forceDescription = "Force"

	debugName        = "debug"
	debugDefault     = false
	debugDescription = "Debugging mode"

	formatName        = "format"
	formatDefault     = "raw"
	validformats      = []string{"raw", "json"}
	formatDescription = fmt.Sprintf("The format type. Can be one of: %v", strings.Join(validformats, ", "))
)

func GetForce() bool {
	return viper.GetBool(forceName)
}

func ConfigureForce(flagset *flag.FlagSet) {
	boolPFlag(flagset, forceName, forceDefault, forceDescription)
}

func GetDebug() bool {
	return viper.GetBool(debugName)
}

func ConfigureDebug(flagset *flag.FlagSet) {
	boolPFlag(flagset, debugName, debugDefault, debugDescription)
}

func GetFormat() string {
	return viper.GetString(formatName)
}

func ConfigureFormat(flagset *flag.FlagSet) {
	stringPFlag(flagset, formatName, formatDefault, formatDescription)
}

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
