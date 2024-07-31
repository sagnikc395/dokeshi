package dokeshi

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// git as our source of truth
type GitDataSource struct{}

// create the output folder, clears it and clones the repo
func (ds *GitDataSource) Fetch(cfg *Config) ([]string, error) {
	from := cfg.Generator.Repository
	to := cfg.Generator.Temp
	branch := cfg.Generator.Branch

	if branch == "" {
		branch = "main"
	}
	fmt.Printf("Fetching the data from %s into %s 🏃‍♂️🏃‍♂️🏃‍♂️\n", from, to)

	if err := createFolderIfNotExist(to); err != nil {
		return nil, err
	}
	if err := clearFolder(to); err != nil {
		return nil, err
	}

	if err := cloneRepo(to, from, branch); err != nil {
		return nil, err
	}
	dirs, err := getContentFolders(to)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Fetching data complete.\n")

	return dirs, nil
}

func createFolderIfNotExists(path string) error {
	//create a temporary folder to copy the files
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(path, os.ModePerm); err != nil {
				return fmt.Errorf("❌ Error Creating directory %s: %v", path, err)
			}
		} else {
			return fmt.Errorf("❌ Error Accessing directory %s: %v", path, err)
		}
	}
	return nil
}

func clearFolder(path string) error {
	dir, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("❌ Error Accessing directory %s: %v", path, err)
	}
	defer dir.Close()

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return fmt.Errorf("❌ Error reading directory %s: %v", path, err)
	}

	for _, name := range names {
		if err = os.RemoveAll(filepath.Join(path, name)); err != nil {
			return fmt.Errorf("❌ Error clearing file %s: %v", name, err)
		}
	}
	return nil
}

func cloneRepo(path, repoURL, branch string) error {
	cmdName := "git"
	initargs := []string{"init", "."}
	cmd := exec.Command(cmdName, initargs...)
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("❌ Error initializing git repo at %s: %v", path, err)
	}

	remoteArgs := []string{"remote", "add", "origin", repoURL}
	cmd = exec.Command(cmdName, remoteArgs...)
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("❌ Error setting remote %s: %v", repoURL, err)
	}
	pullArgs := []string{"pull", "origin", branch}
	cmd = exec.Command(cmdName, pullArgs...)
	cmd.Dir = path
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error pulling %s at %s: %v", branch, path, err)
	}
	return nil
}
