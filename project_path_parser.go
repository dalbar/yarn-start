package yarnstart

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ProjectPathParser provides a mechanism for determining the proper working
// directory for the build process.
type ProjectPathParser struct{}

// NewProjectPathParser creates an instance of a ProjectPathParser.
func NewProjectPathParser() ProjectPathParser {
	return ProjectPathParser{}
}

// Get will resolve the absolute path of the directory specified by the
// $BP_NODE_PROJECT_PATH environment variable. It assumes that
// $BP_NODE_PROJECT_PATH is relative to the provided path argument.
func (p ProjectPathParser) Get(path string) (string, error) {
	customProjPath := os.Getenv("BP_NODE_PROJECT_PATH")
	if customProjPath == "" {
		return path, nil
	}

	_, err := os.Stat(filepath.Join(path, customProjPath))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("expected value derived from BP_NODE_PROJECT_PATH [%s] to be an existing directory", customProjPath)
		}
		return "", err
	}
	return filepath.Join(path, customProjPath), nil
}
