package system

import (
	"github.com/boundedinfinity/userdotd/config"
)

func (this *System) Disable(name string) error {
	es, err := getNames(config.GetEnabledPath())

	if err != nil {
		return err
	}

	if hasName(es, name) {
		d, err := lookupDescriptor(name)

		if err != nil {
			return err
		}

		if err := removeLink(d); err != nil {
			return err
		}
	}

	return nil
}
