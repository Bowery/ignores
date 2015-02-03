// Copyright 2015 Bowery, Inc.

package ignores

import (
	"bufio"
	"os"
	"path/filepath"
)

var (
	// VersionControlSystems is a list of version control system directories
	// that should be ignored.
	VersionControlSystems = []string{".hg", ".git", ".svn", ".bzr"}
)

// Get retrieves a list of files to ignore in a given directory.
// A set of default ignores may be given.
func Get(file string, ignores ...string) ([]string, error) {
	var matches []string

	file, err := os.Open(file)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if file != nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			ignores = append(ignores, scanner.Text())
		}

		err := scanner.Err()
		if err != nil {
			return nil, err
		}
	}

	for _, ignore := range ignores {
		ignoreMatches, err := filepath.Glob(filepath.Join(file, "..", ignore))
		if err != nil {
			return nil, err
		}

		matches = append(matches, ignoreMatches...)
	}

	return matches, nil
}
