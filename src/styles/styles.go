package styles

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func FatalError(e string) {
	fmt.Printf("\n%s%s\n", color.HiBlackString("[Hydrogen Fatal] "), color.HiWhiteString(e))
	os.Exit(0)
}
