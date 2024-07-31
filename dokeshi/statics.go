package dokeshi

import (
	"fmt"
	"io"
	"os"
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
