package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Copy(src, dest string) error {
	if _, err := os.Stat(src); !os.IsExist(err) {
		dir := filepath.Dir(dest)

		err := MakeDirs(dir)
		if err != nil {
			return err
		}

		if b, err := ioutil.ReadFile(src); err == nil {
			err = ioutil.WriteFile(dest, b, os.ModePerm)
			return err
		} else {
			return err
		}

	} else {
		return err
	}
}

func MakeDirs(path string) error {
	_, err := os.Stat(path)

	if !os.IsNotExist(err) {
		return err
	} else {
		err = os.MkdirAll(path, os.ModePerm)
		return err
	}
}
