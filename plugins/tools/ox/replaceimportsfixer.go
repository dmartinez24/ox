package ox

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
)

// ReplaceImportsFixer
type ReplaceImportsFixer struct{}

func (ef ReplaceImportsFixer) Name() string {
	return "/replaceimports"
}

func (ef ReplaceImportsFixer) Fix(ctx context.Context, root string, args []string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() || filepath.Ext(info.Name()) != ".go" {
			return nil
		}

		cc, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		cc = bytes.ReplaceAll(cc, []byte("github.com/gobuffalo/pop/v5"), []byte("github.com/gobuffalo/pop/v6"))
		cc = bytes.ReplaceAll(cc, []byte("github.com/gobuffalo/suite/v3"), []byte("github.com/gobuffalo/suite/v4"))

		err = os.WriteFile(path, []byte(cc), 0644)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
