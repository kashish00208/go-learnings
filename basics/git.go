package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/ini.v1"
)

type GitRepository struct {
	Worktree string
	GitDir   string
	Conf     *ini.File
}

func NewGitRepository(path string, force bool) (*GitRepository, error) {
	repo := &GitRepository{
		Worktree: path,
		GitDir:   filepath.Join(path, ".git"),
	}

	if !(force || isDir(repo.GitDir)) {
		return nil, fmt.Errorf("not a Git repository: %s", path)
	}

	// Read configuration file
	cf := filepath.Join(repo.GitDir, "config")
	if _, err := os.Stat(cf); err == nil {
		cfg, err := ini.Load(cf)
		if err != nil {
			return nil, fmt.Errorf("failed to read config: %v", err)
		}
		repo.Conf = cfg
	} else if !force {
		return nil, fmt.Errorf("configuration file missing")
	}

	if !force {
		versionStr := repo.Conf.Section("core").Key("repositoryformatversion").String()
		version, err := strconv.Atoi(versionStr)
		if err != nil || version != 0 {
			return nil, fmt.Errorf("unsupported repositoryformatversion: %s", versionStr)
		}
	}

	return repo, nil
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
