package protection

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Compile(output string, modifyentry bool) error {

	if modifyentry {
		os.WriteFile("./hydrogen_tmp/entry.asm", []byte(`.globl _start

_start:
 call main
 mov $0, %ebx
 mov $1, %eax
 int $0x80	`), os.ModePerm)
		exec.Command("as", "-o", "./hydrogen_tmp/entry.o", "./hydrogen_tmp/entry.asm").Run()

	}

	e := exec.Command("garble", "-h").Run()
	if e != nil && e.Error() != "exit status 2" {
		return errors.New("Please install garble ( https://github.com/burrowers/garble )")
	}

	var buildtype int
	if strings.HasSuffix(output, ".exe") {
		buildtype = 1
	}

	if MainFile == "" {
		return errors.New("I can't find the main file")
	}

	if runtime.GOOS == "linux" {

		if modifyentry {
			var goos string

			if buildtype == 1 {
				goos = "GOOS=windows"
			} else {
				goos = "GOOS=linux"
			}

			if runtime.GOOS != strings.Split(goos, "GOOS=")[1] {
				return errors.New("Can't compile a " + strings.Split(goos, "GOOS=")[1] + " executable file on " + runtime.GOOS + " ( don't use custom entry point )")
			}

			os.WriteFile("./hydrogen_tmp/compile.sh", []byte(goos+` garble -seed=random -literals -tiny build -trimpath -ldflags="-linkmode=external -extldflags=-static -s -w -X main.main=entry.o" -o ../`+output), os.ModePerm)

			cmdd := exec.Command("bash", "compile.sh")
			cmdd.Dir = "./hydrogen_tmp"
			cmdd.Run()

			cleanHeader(output)

		} else {

			cmd := exec.Command("garble", "-seed=random", "-literals", "-tiny", "build", "-trimpath", "-o", "../"+output)
			cmd.Dir = "./hydrogen_tmp"
			var out strings.Builder
			env := os.Environ()

			if buildtype == 1 {
				env = append(env, "GOOS=windows")
			} else {
				env = append(env, "GOOS=linux")
			}
			cmd.Env = env
			cmd.Stdout = &out

			err := cmd.Run()

			if err != nil {
				return errors.New("Compile error: " + err.Error())
			}

			if buildtype != 1 {
				cleanHeader(output)
			}
		}
	}

	return nil

}
