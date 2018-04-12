package domain

import (
	"os/user"
	"path/filepath"
	"runtime"
)

type Config struct {
	TempDir         string
	InstallationDir string
	GovmHome        string
	BashrcPath      string
	ScriptPath      string
}

func (c *Config) LoadConf() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	bashFilename := ".bash_profile"
	if runtime.GOOS == "linux" {
		bashFilename = ".bashrc"
	}
	userHome := usr.HomeDir
	govmHome := filepath.Join(userHome, ".govm/")
	tempDir := filepath.Join(govmHome, "tmp/")
	installationDir := filepath.Join(govmHome, "installed/")
	bashrcPath := filepath.Join(userHome, bashFilename)
	scriptPath := filepath.Join(govmHome, "wrapper.sh")

	c.GovmHome = govmHome
	c.TempDir = tempDir
	c.InstallationDir = installationDir
	c.BashrcPath = bashrcPath
	c.ScriptPath = scriptPath
	return nil
}
