package system

import (
	"github.com/boundedinfinity/userdotd/config"
)

func (this *System) AvailableLogin() error {
	names, err := getNames(config.GetAvailablePath())

	if err != nil {
		return err
	}

	for _, name := range names {
		this.logger.Println(name)
	}

	return nil
}
