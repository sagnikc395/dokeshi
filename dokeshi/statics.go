package dokeshi

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type StaticsGenerator struct {
	Config *StaticsConfig
}

// static config holds the data for the static sites
type StaticsConfig struct {
	FileToDestination map[string]string
	TemplateToFile    map[string]string
	Template          *template.Template
	Writer            *IndexWriter
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Error reading file %s: %v", src, err)
	}

	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Error creating file %s: %v", dst, err)
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("Error Writing file %s: %v", dst, err)
	}

	if err := out.Sync(); err != nil {
		return fmt.Errorf("Error writing file %s: %v", dst, err)
	}

	return nil
}

func getFolder(path string) string {
	return filepath.Dir(path)
}

func getTitle(path string) string {
	ext := filepath.Ext(path)
	name := filepath.Base(path)
	fileName := name[:len(name)-len(ext)]
	return fmt.Sprintf("%s%s", strings.ToUpper(string(fileName[0])), fileName[1:])
}

func createFolderIfNotExist(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(path, os.ModePerm); err != nil {
				return fmt.Errorf("error creating directory %s: %v", path, err)
			}
		} else {
			return fmt.Errorf("error accessing directory %s: %v", path, err)
		}
	}
	return nil
}
