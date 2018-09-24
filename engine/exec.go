package engine

import (
	"bytes"
	"fmt"
	"github.com/shivakishore14/govm/utils"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Exec(version string, args []string) error {
	path, err := Path(version)
	if err != nil {
		return err
	}
	binPath := filepath.Join(path, "bin/")
	goToolDir := filepath.Join(path, "/pkg/tools/",runtime.GOOS+"_"+runtime.GOARCH)
	os.Setenv("GOTOOLDIR", goToolDir)
	os.Setenv("GOROOT", path)

	if err = os.Setenv("GOROOT", path); err != nil {
		fmt.Println(err)
	}
	utils.AddToPath(binPath)

	cmd := exec.Command(args[0], args[1:]...)
	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput

	err = cmd.Run() // will wait for command to return
	if err != nil {
		return err
	}
	// Only output the commands stdout
	printOutput(cmdOutput.Bytes())

	return nil
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf(string(outs))
	}
}
