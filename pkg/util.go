package pkg

import (
	"errors"
	"os"
)

// FileExists checks if the filepath exists and is not a directory.
// Returns false in case it's not possible to describe the named file.
func FileExists(path string) bool {
	md, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !md.IsDir()
}

func CheckOrCreateConfigFile() error {
	if _, err := os.Stat(CONFIGDIR); err != nil {
		var pError *os.PathError
		if errors.As(err, &pError) {
			return os.MkdirAll(CONFIGDIR, 0750)
		}
		return err
	}

	if !FileExists(CONFIGFILE) {
		f, err := os.Create(CONFIGFILE)
		if err != nil {
			return err
		}
		_ = f.Close()
	}

	return nil
}
