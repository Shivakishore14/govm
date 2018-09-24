package engine

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/mholt/archiver"
	"github.com/shivakishore14/govm/domain"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Install(v domain.Version) error {
	return Download(v)
}

func Download(v domain.Version) error {
	conf := domain.Config{}
	conf.LoadConf()
	return DownloadFile(v, conf)
}

func DownloadFile(v domain.Version, config domain.Config) error {
	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	if err := os.MkdirAll(config.TempDir, os.ModePerm); err != nil {
		return err
	}
	filePath := filepath.Join(config.TempDir, v.FileName)
	out, err := os.Create(filePath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(v.DownloadLink)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("Downloading : %s \n", v.Size)
	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Print("\rDownload complete")

	err = os.Rename(filePath+".tmp", filePath)
	if err != nil {
		return err
	}
	installPath := filepath.Join(config.InstallationDir, v.Name)
	if err := os.MkdirAll(installPath, os.ModePerm); err != nil {
		return err
	}
	fmt.Println("\rExtracting \n")
	err = archiver.TarGz.Open(filePath, installPath)

	return err
}

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}
