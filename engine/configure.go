package engine

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shivakishore14/govm/domain"
	"github.com/shivakishore14/govm/utils"
	"html/template"
	"os"
	"runtime"
	"strings"
)

var scriptBash = `
govm() {
    export GOVMOS="{{.Os}}"
    export GOVMARCH="{{.Arch}}"
    if [[ $1 == "use" ]]; then
        shift;
        d=` + "`govm path $@`" + `
        if [ ${d:0:5} == "PATH:" ]; then
            gopath=${d:5}
            PATH=$gopath/bin:$PATH
            export GOTOOLDIR="$gopath/pkg/tool/$GOVMOS_$GOVMARCH"
            export GOROOT=$gopath
            echo "Using version $1"
        else
            echo "could not find version"
        fi
    else
        command govm "$@"
    fi
}
`

func Configure() error {
	data := struct {
		Os   string
		Arch string
	}{
		Os:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
	config := &domain.Config{}
	config.LoadConf()
	if err := createNecessaryDirs(config); err != nil {
		return nil
	}

	fmt.Printf("HomeDir for Govm  [ %s ] \n", config.GovmHome)
	sourceCommand := "source " + config.ScriptPath
	scriptFile, err := os.Create(config.ScriptPath)
	if err != nil {
		return errors.Wrap(err, "error creating script file")
	}

	t, err := template.New("script").Parse(scriptBash)
	if err != nil {
		return errors.Wrap(err, "error getting parsing script template")
	}

	if err = t.Execute(scriptFile, data); err != nil {
		return errors.Wrap(err, "error executing error template")
	}

	// Update .bashrc
	if _, err := os.Stat(config.BashrcPath); err != nil {
		if os.IsNotExist(err) {
			os.Create(config.BashrcPath)
		}
	}
	if err := writeToFileIfNotPresent(config.BashrcPath, sourceCommand); err != nil {
		return err
	}
	return nil
}

func createNecessaryDirs(config *domain.Config) error {
	if err := utils.CreateDirIfNotPresent(config.GovmHome); err != nil {
		return err
	}
	if err := utils.CreateDirIfNotPresent(config.InstallationDir); err != nil {
		return err
	}
	if err := utils.CreateDirIfNotPresent(config.TempDir); err != nil {
		return err
	}
	return nil
}

func writeToFileIfNotPresent(filename string, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		return errors.Wrap(err, "error opening file :"+filename)
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	var foundFlag bool

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), content) {
			foundFlag = true
		}
	}

	if !foundFlag {
		if _, err = f.WriteString(content + "\n"); err != nil {
			return errors.Wrap(err, "error writing to file")
		}
	}
	return err
}