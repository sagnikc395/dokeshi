package dokeshi

import (
	"fmt"
	"os"
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
