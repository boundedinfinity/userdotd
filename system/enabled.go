package system

import (
	"github.com/boundedinfinity/userdotd/config"
)

func (this *System) Enabled() error {
	names, err := getNames(config.GetEnabledPath())

	if err != nil {
		return err
	}

	for _, n := range names {
		this.logger.Println(n)
	}

	return nil
}

func (this *System) Enable(name string) error {
	as, err := getNames(config.GetAvailablePath())

	if err != nil {
		return err
	}

	es, err := getNames(config.GetEnabledPath())

	if err != nil {
		return err
	}

	if hasName(as, name) && !hasName(es, name) {
		d, err := lookupDescriptor(name)

		if err != nil {
			return err
		}

		if err := ensureLink(d); err != nil {
			return err
		}
	}

	return nil
}
