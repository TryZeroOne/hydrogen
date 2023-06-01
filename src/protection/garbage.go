package protection

import (
	"crypto/rand"
	"errors"
	"os"
)

func Garbage(file string) error {
	target, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("Can't find the file: " + file)
	}
	defer target.Close()
	targetSt, _ := target.Stat()
	targetSize := targetSt.Size()

	if targetSize < 5 {
		return errors.New("Can't find the file: " + file)
	}
	offs := targetSize + Random(5000, 10000)

	for i := 0; i <= 200; i++ {
		_, err = target.WriteString(GenGarbage(targetSize - offs/2)[:Random(1000, 3000)])
	}
	if err != nil {
		return errors.New("Garbage error: " + err.Error())
	}

	return nil
}

func GenGarbage(size int64) string {
	garb := make([]byte, size)

	_, err := rand.Read(garb)

	if err != nil {
		panic(err)
	}

	return string(garb)
}
