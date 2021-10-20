package main

import (
	"errors"
	"os"
)

type FileTask struct {
	Name string
}

type Manage interface {
	Exists() bool
	ReadToString() (string, error)
	RewriteWithString(string) error
}

func (file FileTask) Exists() bool {
	_, err := os.Stat(file.Name)

	return !errors.Is(err, os.ErrNotExist)
}

func (file FileTask) ReadToString() (string, error) {
	content, err := os.ReadFile(file.Name)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (file FileTask) RewriteWithString(newContent string) error {
	return os.WriteFile(file.Name, []byte(newContent), 0666)
}

func MakeFile(fileName string) (FileTask, bool) {
	file := FileTask{fileName}
	return file, file.Exists()
}
