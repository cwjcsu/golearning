package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func GetProjectRoot() string {
	binDir, err := executableDir()
	if err != nil {
		return ""
	}
	return path.Dir(binDir)
}

func executableDir() (string, error) {
	pathAbs, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Dir(pathAbs), nil
}

func Welcome() {
	fmt.Print("****************************")
	fmt.Print("********欢迎学习Go语言*********")
	fmt.Print("****************************")
}

func ReaderExample() {

}
