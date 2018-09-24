package engine

import (
	"github.com/shivakishore14/govm/domain"
	"os"
	"path/filepath"
)

func Path(name string) (string, error) {
	config := domain.Config{}
	config.LoadConf()
	versionPath := filepath.Join(config.InstallationDir, name, "go")

	if _, err := os.Stat(versionPath); os.IsNotExist(err) {
		return "", domain.ErrVersionNotFound
	}

	return versionPath, nil
}
