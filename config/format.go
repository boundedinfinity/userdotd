package config

import (
	"fmt"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/spf13/viper"
)

const (
	formatName = "format"
)

var (
	formatDefault     = "raw"
	validformats      = []string{"raw", "json"}
	formatDescription = fmt.Sprintf("The format type. Can be one of: %v", strings.Join(validformats, ", "))
)

func GetFormat() string {
	return viper.GetString(formatName)
}

func ConfigureFormat(flagset *flag.FlagSet) {
	stringPFlag(flagset, formatName, formatDefault, formatDescription)
}
