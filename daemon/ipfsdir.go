package daemon

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	envIpfsPath = "IPFS_PATH"
	defIpfsDir  = ".ipfs"
)

// IpfsDir returns the path of the ipfs directory.  If dir specified, then
// returns the expanded version dir.  If dir is "", then return the directory
// set by IPFS_PATH, or if IPFS_PATH is not set, then return the default
// location in the home directory.
func IpfsDir(dir string) (string, error) {
	var err error
	if dir == "" {
		dir = os.Getenv(envIpfsPath)
	}
	if dir != "" {
		dir, err = Expand(dir)
		if err != nil {
			return "", err
		}
		return dir, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if home == "" {
		return "", errors.New("could not determine IPFS_PATH, home dir not set")
	}

	return filepath.Join(home, defIpfsDir), nil
}

// CheckIpfsDir gets the ipfs directory and checks that the directory exists.
func CheckIpfsDir(dir string) (string, error) {
	var err error
	dir, err = IpfsDir(dir)
	if err != nil {
		return "", err
	}

	_, err = os.Stat(dir)
	if err != nil {
		return "", err
	}

	return dir, nil
}

func Expand(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", errors.New("cannot expand user-specific home dir")
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, path[1:]), nil
}
