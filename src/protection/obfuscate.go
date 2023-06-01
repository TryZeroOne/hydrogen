package protection

import (
	"os"
	"path/filepath"
	"strings"
)

var MainFile string

// Obfuscates files in the specified directory
func ObfuscateFiles(path string) {
	li, _ := os.ReadDir(path)

	for _, i := range li {
		if i.IsDir() {
			ObfuscateFiles(path + "/" + i.Name())
		} else {
			path, _ := filepath.Abs(path)

			if filepath.Ext(filepath.Join(path, i.Name())) == ".go" {

				readf, _ := os.ReadFile(filepath.Join(path, i.Name()))

				mainfile := GetMainFile(readf, filepath.Join(path, i.Name()))

				MainFile = mainfile

				if mainfile == filepath.Join(path, i.Name()) {
					continue
				}

				os.Rename(filepath.Join(path, i.Name()), filepath.Join(path, UniqueName()+".go"))
			}

		}

	}

}

// Prevents the main file from being obfuscated
func GetMainFile(content []byte, filename string) string {
	for _, i := range strings.Split(string(content), "\n") {
		if string(i) == "func main() {" {
			return filename
		}

	}
	return ""
}
