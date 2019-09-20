package integration_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack"
)

func runCNB2CF(dir string, args ...string) (string, error) {
	root, err := FindRoot()
	if err != nil {
		return "", err
	}

	command := exec.Command(filepath.Join(root, "build", "cnb2cf"), args...)
	if dir != "" {
		command.Dir = dir
	}

	buffer := bytes.NewBuffer([]byte{})
	command.Stdout = io.MultiWriter(os.Stdout, buffer)
	command.Stderr = io.MultiWriter(os.Stderr, buffer)

	err = command.Run()
	return buffer.String(), err
}

func FindRoot() (string, error) {
	dir, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}
	for {
		if dir == "/" {
			return "", fmt.Errorf("Could not find README.md in the directory hierarchy")
		}
		if exist, err := libbuildpack.FileExists(filepath.Join(dir, "README.md")); err != nil {
			return "", err
		} else if exist {
			return dir, nil
		}
		dir, err = filepath.Abs(filepath.Join(dir, ".."))
		if err != nil {
			return "", err
		}
	}
}
