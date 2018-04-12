package domain

import "errors"

type Version struct {
	Name         string
	FileName     string
	Kind         string
	Os           string
	Arch         string
	Size         string
	SHA1         string
	DownloadLink string
}
type Versions []Version
func (v Version) IsEmpty() bool {
	if v.Name != "" && v.Size != "" && v.DownloadLink != "" {
		return false
	}
	return true
}

var ErrVersionNotFound = errors.New("version not found")
