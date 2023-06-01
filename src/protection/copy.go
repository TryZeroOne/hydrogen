package protection

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var TempName string

func CopyFiles(path string) {
	TempName = "hydrogen_tmp" + "/"
	os.Mkdir(TempName, os.ModePerm)

	li, _ := os.ReadDir(path)

	for _, i := range li {
		if i.Name()+"/" == TempName {
			continue
		}
		path, _ := filepath.Abs(path)
		filepath := filepath.Join(path, i.Name())
		if i.IsDir() {
			os.Mkdir(TempName+i.Name(), os.ModePerm)
			CopyDir(filepath, TempName+i.Name())
		} else {
			copy(filepath, TempName+i.Name())
		}
	}

}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func CopyDir(source, destination string) error {
	var err error = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		var relPath string = strings.Replace(path, source, "", 1)
		if relPath == "" {
			return nil
		}
		if info.IsDir() {
			return os.Mkdir(filepath.Join(destination, relPath), 0755)
		} else {
			var data, err1 = os.ReadFile(filepath.Join(source, relPath))
			if err1 != nil {
				return err1
			}
			return os.WriteFile(filepath.Join(destination, relPath), data, 0777)
		}
	})
	return err
}
