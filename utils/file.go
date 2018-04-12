package utils

import "os"

func CreateDirIfNotPresent(dirName string) error{
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		// Path does not exists
		return os.MkdirAll(dirName, os.ModePerm)
	}
	return nil
}