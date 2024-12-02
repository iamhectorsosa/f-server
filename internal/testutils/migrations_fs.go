package testutils

import (
	"io/fs"
	"os"
	"path/filepath"
)

func CreateMigrationsFs(dir string, migrations map[string]string) (fs.FS, func(), error) {
	tempDir, err := os.MkdirTemp("", "migrations")
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		os.RemoveAll(tempDir)
	}

	migrationsDir := filepath.Join(tempDir, dir)
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		cleanup()
		return nil, nil, err
	}

	for filename, content := range migrations {
		filePath := filepath.Join(migrationsDir, filename)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			cleanup()
			return nil, nil, err
		}
	}
	return os.DirFS(tempDir), cleanup, nil
}
