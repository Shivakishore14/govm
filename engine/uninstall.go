package engine

import (
	"github.com/golang-vm/govm/domain"
	"os"
	"path/filepath"
)

func Uninstall(versionName string) error {
	config := domain.Config{}
	config.LoadConf()
	installedPath := filepath.Join(config.InstallationDir, versionName)
	if _, err := os.Stat(installedPath); err != nil {
		return domain.ErrVersionNotFound
	}
	return os.RemoveAll(installedPath)
}
