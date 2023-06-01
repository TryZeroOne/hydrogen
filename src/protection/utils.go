package protection

import (
	"crypto/rand"
	"math/big"
	random "math/rand"
	"os"
	"path/filepath"
	"time"
)

var Names = make([]string, 0)
var FuncsNames = make([]string, 0)
var Vars = make([]string, 0)

func UniqueName() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, Random(4, 10))
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
		for _, i := range Names {
			if string(s) == i {
				UniqueName()
			} else {
				Names = append(Names, string(s))
			}
		}
	}
	return string(s)

}

func UniqueFunc() string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, Random(4, 10))
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
		for _, i := range FuncsNames {
			if string(s) == i {
				UniqueFunc()
			} else {
				FuncsNames = append(FuncsNames, string(s))
			}
		}
	}
	return string(s)

}

func UniqueVar() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, Random(4, 10))
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
		for _, i := range Vars {
			if string(s) == i {
				UniqueVar()
			} else {
				Vars = append(Vars, string(s))
			}
		}
	}
	return string(s)

}
func RandomArray(arr []string) string {
	cStrs := len(arr)

	s1 := random.NewSource(time.Now().UnixNano())
	r1 := random.New(s1)

	randstr := arr[r1.Intn(cStrs)]
	return randstr
}

func RandomString(min int64, max int64) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, Random(min, max))
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
	}

	return string(s)

}

func RandomComment(min int64, max int64) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	s := make([]rune, Random(min, max))
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
	}

	return string(s)

}

func Random(min, max int64) int64 {
	bg := big.NewInt(max - min)

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	return n.Int64() + min
}

func GetMain(path string) {
	li, _ := os.ReadDir(path)

	for _, i := range li {
		if i.IsDir() {
			GetMain(path + "/" + i.Name())
		} else {
			path, _ := filepath.Abs(path)

			if filepath.Ext(filepath.Join(path, i.Name())) == ".go" {

				readf, _ := os.ReadFile(filepath.Join(path, i.Name()))

				mainfile := GetMainFile(readf, filepath.Join(path, i.Name()))

				MainFile = mainfile

				if mainfile == filepath.Join(path, i.Name()) {
					return
				}

			}

		}

	}

}

func GenString(length int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
	}
	return string(s)
}
