package pathutil

import (
	"io"
	"os"

	"github.com/boundedinfinity/userdotd/embedded"
	"github.com/udhos/equalfile"
)

func EmbeddedEqual(ep, rp string) (bool, error) {
	ef, err := embedded.OpenFile(ep)

	if err != nil {
		return false, err
	}

	defer ef.Close()

	rf, err := os.Open(rp)

	if err != nil {
		return false, err
	}

	defer rf.Close()

	eq := equalfile.New(nil, equalfile.Options{})
	isEq, err := eq.CompareReader(ef, rf)

	return isEq, err
}

func EmbeddedCopy(ep, rp string) error {
	ef, err := embedded.OpenFile(ep)

	if err != nil {
		return err
	}

	defer ef.Close()

	rf, err := os.Create(rp)

	if err != nil {
		return err
	}

	defer rf.Close()

	_, err = io.Copy(rf, ef)

	return err
}
