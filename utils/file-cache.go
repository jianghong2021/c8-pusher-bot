package utils

import (
	"fmt"
	"os"
)

var (
	CachePath = "data/"
)

func init() {
	checkDir()
}

func checkDir() {
	if _, err := os.Stat(CachePath); os.IsNotExist(err) {
		os.MkdirAll(CachePath, 0755)
	}
}

func SetCache(key string, val string) error {
	checkDir()
	fname := fmt.Sprintf("data-%s.txt", key)

	f, err := os.OpenFile(CachePath+fname, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = f.WriteString(val)
	if err != nil {
		return err
	}

	return nil
}

func GetCache(key string) (string, error) {
	fname := fmt.Sprintf("data-%s.txt", key)

	if _, err := os.Stat(CachePath + fname); os.IsNotExist(err) {
		return "", err
	}

	data, err := os.ReadFile(CachePath + fname)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
