package styles

import (
	"fmt"
	"time"
)

var Stages int

func ProgressMenu(compile bool, obf bool, output string, junk bool, garbage, compress, encrypt bool) {

	var stages int

	if compile {
		stages++
	}
	if encrypt {
		stages++
	}
	if obf {
		stages++
	}
	if junk {
		stages++
	}
	if garbage {
		stages++
	}
	if compress {
		stages++
	}

	go func() {
		for per := 0; per <= 100; per++ {
			if per == 99 {
				per = 0
				continue
			}
			ProgressBar(per, stages)
		}
	}()

}

func ProgressBar(per int, allstages int) {
	chars := per * 50 / 100

	fmt.Printf("\rHydrogen %d/%d [", Stages, allstages)

	for char := 0; char < chars; char++ {
		fmt.Printf("*")
	}
	for charx := 0; charx < 50-chars; charx++ {
		fmt.Printf(" ")
	}
	fmt.Printf("]")
	time.Sleep(3 * time.Millisecond)
}
