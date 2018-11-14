package system

import (
	"errors"

	"github.com/boundedinfinity/userdotd/config"
)

var (
	ErrNotInitialized = errors.New("system not initialized")
)

func (this *System) InitializeSystem() error {
	if err := this.initializeFish(); err != nil {
		return err
	}

	if err := this.initializeBash(); err != nil {
		return err
	}

	return nil
}

func (this *System) IsInitialized() bool {
	return dirExists(config.GetuserdotdDir())
}

func (this *System) EnsureInitialized() error {
	if !this.IsInitialized() {
		return ErrNotInitialized
	}

	return nil
}

func (this *System) initializeFish() error {
	for _, cd := range getFishDescriptors() {
		if err := ensureFile(cd); err != nil {
			return err
		}
	}

	return nil
}

func (this *System) initializeBash() error {
	return nil
}
