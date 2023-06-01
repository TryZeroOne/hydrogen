package main

import (
	"os"
	"runtime"

	"github.com/TryZeroOne/hydrogen/src"
	"github.com/TryZeroOne/hydrogen/src/cmd"
	"github.com/TryZeroOne/hydrogen/src/styles"
)

func main() {

	if runtime.GOOS != "linux" {
		styles.FatalError("Unsupported OS")
		return
	}

	os.RemoveAll("./hydrogen_tmp/")
	flags, action, output := cmd.Init()
	if flags == "" {
		return
	}

	defer func() {
		os.RemoveAll("./hydrogen_tmp/")
	}()

	src.Launch(flags, action, output)

}
